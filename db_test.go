package kivikmock

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/flimzy/diff"
	"github.com/flimzy/testy"
	"github.com/go-kivik/kivik"
	"github.com/go-kivik/kivik/driver"
)

func TestCloseDB(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectClose().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").Close(context.TODO())
			testy.Error(t, "foo err", err)
		},
		err: "",
	})
	tests.Add("unexpected", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").Close(context.TODO())
			testy.Error(t, "call to DB.Close() was not expected, all expectations already fulfilled", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectClose().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").Close(newCanceledContext())
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectClose()
			foo.ExpectClose()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			err := foo.Close(context.TODO())
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Add("callback", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectClose().WillExecute(func(_ context.Context) error {
				return errors.New("custom error")
			})
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").Close(context.TODO())
			testy.Error(t, "custom error", err)
		},
	})
	tests.Run(t, testMock)
}

func TestAllDocs(t *testing.T) { // nolint: gocyclo
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectAllDocs().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.AllDocs(context.TODO())
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("unexpected", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.AllDocs(context.TODO())
			testy.Error(t, "call to DB.AllDocs() was not expected, all expectations already fulfilled", err)
		},
	})
	tests.Add("rows close error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectAllDocs().WillReturn(NewRows().CloseError(errors.New("bar err")))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			rows, err := db.AllDocs(context.TODO())
			testy.Error(t, "", err)
			testy.Error(t, "bar err", rows.Close())
		},
	})
	tests.Add("rows offset", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectAllDocs().WillReturn(NewRows().Offset(123))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			rows, err := db.AllDocs(context.TODO())
			testy.Error(t, "", err)
			if o := rows.Offset(); o != 123 {
				t.Errorf("Unexpected offset: %d", o)
			}
		},
	})
	tests.Add("rows totalrows", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectAllDocs().WillReturn(NewRows().TotalRows(123))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			rows, err := db.AllDocs(context.TODO())
			testy.Error(t, "", err)
			if o := rows.TotalRows(); o != 123 {
				t.Errorf("Unexpected total rows: %d", o)
			}
		},
	})
	tests.Add("rows update seq", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectAllDocs().WillReturn(NewRows().UpdateSeq("1-xxx"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			rows, err := db.AllDocs(context.TODO())
			testy.Error(t, "", err)
			if o := rows.UpdateSeq(); o != "1-xxx" {
				t.Errorf("Unexpected update seq: %s", o)
			}
		},
	})
	tests.Add("rows warning", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectAllDocs().WillReturn(NewRows().Warning("Caution!"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			rows, err := db.AllDocs(context.TODO())
			testy.Error(t, "", err)
			if o := rows.Warning(); o != "Caution!" {
				t.Errorf("Unexpected warning seq: %s", o)
			}
		},
	})
	tests.Add("rows", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectAllDocs().WillReturn(NewRows().
				AddRow(&driver.Row{ID: "foo"}).
				AddRow(&driver.Row{ID: "bar"}).
				AddRow(&driver.Row{ID: "baz"}))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			rows, err := db.AllDocs(context.TODO())
			testy.Error(t, "", err)
			ids := []string{}
			for rows.Next() {
				ids = append(ids, rows.ID())
			}
			expected := []string{"foo", "bar", "baz"}
			if d := diff.Interface(expected, ids); d != nil {
				t.Error(d)
			}
		},
	})
	tests.Add("row error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectAllDocs().WillReturn(NewRows().
				AddRow(&driver.Row{ID: "foo"}).
				AddRowError(errors.New("foo err")))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			rows, err := db.AllDocs(context.TODO())
			testy.Error(t, "", err)
			ids := []string{}
			for rows.Next() {
				ids = append(ids, rows.ID())
			}
			expected := []string{"foo"}
			if d := diff.Interface(expected, ids); d != nil {
				t.Error(d)
			}
			testy.Error(t, "foo err", rows.Err())
		},
	})
	tests.Add("options", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectAllDocs().WithOptions(map[string]interface{}{"foo": 123})
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.AllDocs(context.TODO())
			testy.ErrorRE(t, `map\[foo:123]`, err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectAllDocs().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.AllDocs(newCanceledContext())
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("row delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectAllDocs().WillReturn(NewRows().
				AddDelay(time.Millisecond).
				AddRow(&driver.Row{ID: "foo"}).
				AddDelay(time.Second).
				AddRow(&driver.Row{ID: "bar"}))
		},
		test: func(t *testing.T, c *kivik.Client) {
			ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
			defer cancel()
			rows, err := c.DB(ctx, "foo").AllDocs(ctx)
			testy.Error(t, "", err)
			ids := []string{}
			for rows.Next() {
				ids = append(ids, rows.ID())
			}
			expected := []string{"foo"}
			if d := diff.Interface(expected, ids); d != nil {
				t.Error(d)
			}
			testy.Error(t, "context deadline exceeded", rows.Err())
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectAllDocs()
			foo.ExpectAllDocs()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			_, err := foo.AllDocs(context.TODO())
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Run(t, testMock)
}

func TestBulkGet(t *testing.T) { // nolint: gocyclo
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectBulkGet().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.BulkGet(context.TODO(), []kivik.BulkGetReference{})
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("unexpected", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.BulkGet(context.TODO(), []kivik.BulkGetReference{})
			testy.Error(t, "call to DB.BulkGet() was not expected, all expectations already fulfilled", err)
		},
	})
	tests.Add("rows", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectBulkGet().WillReturn(NewRows().
				AddRow(&driver.Row{ID: "foo"}).
				AddRow(&driver.Row{ID: "bar"}).
				AddRow(&driver.Row{ID: "baz"}))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			rows, err := db.BulkGet(context.TODO(), []kivik.BulkGetReference{})
			testy.Error(t, "", err)
			ids := []string{}
			for rows.Next() {
				ids = append(ids, rows.ID())
			}
			expected := []string{"foo", "bar", "baz"}
			if d := diff.Interface(expected, ids); d != nil {
				t.Error(d)
			}
		},
	})
	tests.Add("options", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectBulkGet().WithOptions(map[string]interface{}{"foo": 123})
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.BulkGet(context.TODO(), []kivik.BulkGetReference{})
			testy.ErrorRE(t, `map\[foo:123]`, err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectBulkGet().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.BulkGet(newCanceledContext(), []kivik.BulkGetReference{})
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectBulkGet()
			foo.ExpectBulkGet()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			_, err := foo.BulkGet(context.TODO(), []kivik.BulkGetReference{})
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Run(t, testMock)
}

func TestFind(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectFind().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.Find(context.TODO(), nil)
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("unmatched query", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectFind().WithQuery(123)
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.Find(context.TODO(), 321)
			testy.ErrorRE(t, "has query: 123", err)
		},
	})
	tests.Add("rows", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectFind().WillReturn(NewRows().
				AddRow(&driver.Row{ID: "foo"}).
				AddRow(&driver.Row{ID: "bar"}).
				AddRow(&driver.Row{ID: "baz"}))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			rows, err := db.Find(context.TODO(), 7)
			testy.Error(t, "", err)
			ids := []string{}
			for rows.Next() {
				ids = append(ids, rows.ID())
			}
			expected := []string{"foo", "bar", "baz"}
			if d := diff.Interface(expected, ids); d != nil {
				t.Error(d)
			}
		},
	})
	tests.Add("query", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectFind().WithQuery(map[string]interface{}{"foo": "123"})
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.Find(context.TODO(), map[string]string{"foo": "123"})
			testy.ErrorRE(t, "", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectFind().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.Find(newCanceledContext(), 0)
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectFind()
			foo.ExpectFind()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			_, err := foo.Find(context.TODO(), 1)
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Run(t, testMock)
}

func TestCreateIndex(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectCreateIndex().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").CreateIndex(context.TODO(), "foo", "bar", 123)
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("unmatched index", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectCreateIndex().WithIndex(321)
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").CreateIndex(context.TODO(), "foo", "bar", 123)
			testy.ErrorRE(t, "has index: 321", err)
		},
	})
	tests.Add("ddoc", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectCreateIndex().WithDDocID("moo")
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").CreateIndex(context.TODO(), "foo", "bar", 123)
			testy.ErrorRE(t, "has ddoc: moo", err)
		},
	})
	tests.Add("name", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectCreateIndex().WithName("moo")
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").CreateIndex(context.TODO(), "foo", "bar", 123)
			testy.ErrorRE(t, "has name: moo", err)
		},
	})
	tests.Add("index", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectCreateIndex().WithIndex("moo")
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").CreateIndex(context.TODO(), "foo", "bar", "moo")
			testy.Error(t, "", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectCreateIndex().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").CreateIndex(newCanceledContext(), "foo", "bar", "moo")
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectCreateIndex()
			foo.ExpectCreateIndex()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			err := foo.CreateIndex(context.TODO(), "foo", "bar", 123)
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Run(t, testMock)
}

func TestGetIndexes(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectGetIndexes().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").GetIndexes(context.TODO())
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("indexes", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectGetIndexes().WillReturn([]driver.Index{
				{Name: "foo"},
				{Name: "bar"},
			})
		},
		test: func(t *testing.T, c *kivik.Client) {
			indexes, err := c.DB(context.TODO(), "foo").GetIndexes(context.TODO())
			testy.Error(t, "", err)
			expected := []kivik.Index{
				{Name: "foo"},
				{Name: "bar"},
			}
			if d := diff.Interface(expected, indexes); d != nil {
				t.Error(d)
			}
		},
	})
	tests.Add("unexpected", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").GetIndexes(context.TODO())
			testy.Error(t, "call to DB.GetIndexes() was not expected, all expectations already fulfilled", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectGetIndexes().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").GetIndexes(newCanceledContext())
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectGetIndexes()
			foo.ExpectGetIndexes()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			_, err := foo.GetIndexes(context.TODO())
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Run(t, testMock)
}

func TestDeleteIndex(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectDeleteIndex().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").DeleteIndex(context.TODO(), "foo", "bar")
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("ddoc", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectDeleteIndex().WithDDoc("oink")
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").DeleteIndex(context.TODO(), "foo", "bar")
			testy.ErrorRE(t, "has ddoc: oink", err)
		},
	})
	tests.Add("name", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectDeleteIndex().WithName("oink")
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").DeleteIndex(context.TODO(), "foo", "bar")
			testy.ErrorRE(t, "has name: oink", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectDeleteIndex().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").DeleteIndex(newCanceledContext(), "foo", "bar")
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectDeleteIndex()
			foo.ExpectDeleteIndex()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			err := foo.DeleteIndex(context.TODO(), "foo", "bar")
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Run(t, testMock)
}

func TestExplain(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectExplain().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").Explain(context.TODO(), "foo")
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("unexpected", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").Explain(context.TODO(), "foo")
			testy.Error(t, "call to DB.Explain() was not expected, all expectations already fulfilled", err)
		},
	})
	tests.Add("query", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectExplain().WithQuery(map[string]string{"foo": "bar"})
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").Explain(context.TODO(), map[string]interface{}{"foo": "bar"})
			testy.Error(t, "", err)
		},
	})
	tests.Add("plan", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectExplain().WillReturn(&driver.QueryPlan{DBName: "foo"})
		},
		test: func(t *testing.T, c *kivik.Client) {
			plan, err := c.DB(context.TODO(), "foo").Explain(context.TODO(), map[string]interface{}{"foo": "bar"})
			testy.Error(t, "", err)
			expected := &kivik.QueryPlan{DBName: "foo"}
			if d := diff.Interface(expected, plan); d != nil {
				t.Error(d)
			}
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectExplain().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").Explain(newCanceledContext(), 123)
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectExplain()
			foo.ExpectExplain()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			_, err := foo.Explain(context.TODO(), 123)
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Run(t, testMock)
}

func TestCreateDoc(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectCreateDoc().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, _, err := c.DB(context.TODO(), "foo").CreateDoc(context.TODO(), "foo")
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("return", func() interface{} {
		docID, rev := "foo", "1-xxx"
		return mockTest{
			setup: func(m *Client) {
				db := m.NewDB()
				m.ExpectDB().WillReturn(db)
				db.ExpectCreateDoc().WillReturn(docID, rev)
			},
			test: func(t *testing.T, c *kivik.Client) {
				i, r, err := c.DB(context.TODO(), "foo").CreateDoc(context.TODO(), "foo")
				testy.Error(t, "", err)
				if i != docID || r != rev {
					t.Errorf("Unexpected docID/Rev: %s/%s", i, r)
				}
			},
		}
	})
	tests.Add("mismatched doc", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectCreateDoc().WithDoc("foo")
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, _, err := c.DB(context.TODO(), "foo").CreateDoc(context.TODO(), "bar")
			testy.ErrorRE(t, `has doc: "foo"`, err)
		},
	})
	tests.Add("options", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectCreateDoc().WithOptions(map[string]interface{}{"foo": "bar"})
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, _, err := c.DB(context.TODO(), "foo").CreateDoc(context.TODO(), "bar", map[string]interface{}{})
			testy.ErrorRE(t, `has options: map\[foo:bar]`, err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectCreateDoc().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, _, err := c.DB(context.TODO(), "foo").CreateDoc(newCanceledContext(), 123)
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectCreateDoc()
			foo.ExpectCreateDoc()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			_, _, err := foo.CreateDoc(context.TODO(), 123)
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Run(t, testMock)
}

func TestCompact(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectCompact().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").Compact(context.TODO())
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectCompact().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").Compact(newCanceledContext())
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("unexpected", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").Compact(context.TODO())
			testy.Error(t, "call to DB.Compact() was not expected, all expectations already fulfilled", err)
		},
	})
	tests.Run(t, testMock)
}

func TestCompactView(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectCompactView().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").CompactView(context.TODO(), "foo")
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("ddocID", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectCompactView().WithDDoc("foo")
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").CompactView(context.TODO(), "foo")
			testy.Error(t, "", err)
		},
	})
	tests.Add("unexpected ddoc", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectCompactView().WithDDoc("foo")
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").CompactView(context.TODO(), "bar")
			testy.ErrorRE(t, "has ddocID: foo", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectCompactView().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").CompactView(newCanceledContext(), "foo")
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectCompactView()
			foo.ExpectCompactView()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			err := foo.CompactView(context.TODO(), "foo")
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Run(t, testMock)
}

func TestViewCleanup(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectViewCleanup().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").ViewCleanup(context.TODO())
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectViewCleanup().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").ViewCleanup(newCanceledContext())
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("unexpected", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").ViewCleanup(context.TODO())
			testy.Error(t, "call to DB.ViewCleanup() was not expected, all expectations already fulfilled", err)
		},
	})
	tests.Run(t, testMock)
}

func TestPut(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectPut().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").Put(context.TODO(), "foo", 123)
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectPut().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").Put(newCanceledContext(), "foo", 123)
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("unexpected", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").Put(context.TODO(), "foo", 123)
			testy.Error(t, "call to DB.Put() was not expected, all expectations already fulfilled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectPut()
			foo.ExpectPut()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			_, err := foo.Put(context.TODO(), "foo", 123)
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Add("wrong id", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectPut().WithDocID("foo")
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").Put(context.TODO(), "bar", 123)
			testy.ErrorRE(t, "has docID: foo", err)
		},
	})
	tests.Add("wrong doc", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectPut().WithDoc(map[string]string{"foo": "bar"})
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").Put(context.TODO(), "foo", 123)
			testy.ErrorRE(t, "has docID: foo", err)
		},
	})
	tests.Add("wrong options", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectPut().WithOptions(map[string]interface{}{"foo": "bar"})
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").Put(context.TODO(), "foo", 123, map[string]interface{}{"foo": 123})
			testy.ErrorRE(t, "has docID: foo", err)
		},
	})
	tests.Add("success", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectPut().WillReturn("oink")
		},
		test: func(t *testing.T, c *kivik.Client) {
			result, err := c.DB(context.TODO(), "foo").Put(context.TODO(), "foo", 123)
			testy.ErrorRE(t, "", err)
			if result != "oink" {
				t.Errorf("Unexpected result: %s", result)
			}
		},
	})
	tests.Run(t, testMock)
}

func TestGetMeta(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectGetMeta().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, _, err := c.DB(context.TODO(), "foo").GetMeta(context.TODO(), "foo")
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectGetMeta().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, _, err := c.DB(context.TODO(), "foo").GetMeta(newCanceledContext(), "foo")
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("unexpected", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, _, err := c.DB(context.TODO(), "foo").GetMeta(context.TODO(), "foo")
			testy.Error(t, "call to DB.GetMeta() was not expected, all expectations already fulfilled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectGetMeta()
			foo.ExpectGetMeta()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			_, _, err := foo.GetMeta(context.TODO(), "foo")
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Add("wrong id", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectGetMeta().WithDocID("foo")
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, _, err := c.DB(context.TODO(), "foo").GetMeta(context.TODO(), "bar")
			testy.ErrorRE(t, "has docID: foo", err)
		},
	})
	tests.Add("wrong options", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectGetMeta().WithOptions(map[string]interface{}{"foo": "bar"})
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, _, err := c.DB(context.TODO(), "foo").GetMeta(context.TODO(), "foo", map[string]interface{}{"foo": 123})
			testy.ErrorRE(t, "has docID: foo", err)
		},
	})
	tests.Add("success", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectGetMeta().WillReturn(123, "1-oink")
		},
		test: func(t *testing.T, c *kivik.Client) {
			size, rev, err := c.DB(context.TODO(), "foo").GetMeta(context.TODO(), "foo")
			testy.ErrorRE(t, "", err)
			if size != 123 {
				t.Errorf("Unexpected size: %d", size)
			}
			if rev != "1-oink" {
				t.Errorf("Unexpected rev: %s", rev)
			}
		},
	})
	tests.Run(t, testMock)
}

func TestFlush(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectFlush().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").Flush(context.TODO())
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectFlush().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").Flush(newCanceledContext())
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectFlush()
			foo.ExpectFlush()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			err := foo.Flush(context.TODO())
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Run(t, testMock)
}

func TestDeleteAttachment(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectDeleteAttachment().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").DeleteAttachment(context.TODO(), "foo", "1-foo", "foo.txt")
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectDeleteAttachment().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").DeleteAttachment(newCanceledContext(), "foo", "1-foo", "foo.txt")
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectDeleteAttachment()
			foo.ExpectDeleteAttachment()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			_, err := foo.DeleteAttachment(context.TODO(), "foo", "1-foo", "foo.txt")
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Add("wrong docID", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectDeleteAttachment().WithDocID("bar")
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").DeleteAttachment(context.TODO(), "foo", "1-foo", "foo.txt")
			testy.ErrorRE(t, "has docID: bar", err)
		},
	})
	tests.Add("wrong rev", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectDeleteAttachment().WithRev("2-asd")
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").DeleteAttachment(context.TODO(), "foo", "1-foo", "foo.txt")
			testy.ErrorRE(t, "has rev: 2-asd", err)
		},
	})
	tests.Add("wrong filename", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectDeleteAttachment().WithFilename("bar.txt")
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").DeleteAttachment(context.TODO(), "foo", "1-foo", "foo.txt")
			testy.ErrorRE(t, "has filename: bar.txt", err)
		},
	})
	tests.Add("wrong options", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectDeleteAttachment().WithOptions(map[string]interface{}{"foo": "baz"})
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").DeleteAttachment(context.TODO(), "foo", "1-foo", "foo.txt")
			testy.ErrorRE(t, `has options: map\[foo:baz]`, err)
		},
	})
	tests.Add("success", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectDeleteAttachment().WillReturn("2-fds")
		},
		test: func(t *testing.T, c *kivik.Client) {
			rev, err := c.DB(context.TODO(), "foo").DeleteAttachment(context.TODO(), "foo", "1-foo", "foo.txt")
			testy.Error(t, "", err)
			if rev != "2-fds" {
				t.Errorf("Unexpected rev: %s", rev)
			}
		},
	})
	tests.Run(t, testMock)
}

func TestDelete(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectDelete().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").Delete(context.TODO(), "foo", "1-foo")
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectDelete().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").Delete(newCanceledContext(), "foo", "1-foo")
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectDelete()
			foo.ExpectDelete()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			_, err := foo.Delete(context.TODO(), "foo", "1-foo")
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Add("wrong docID", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectDelete().WithDocID("bar")
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").Delete(context.TODO(), "foo", "1-foo")
			testy.ErrorRE(t, "has docID: bar", err)
		},
	})
	tests.Add("wrong rev", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectDelete().WithRev("2-lkj")
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").Delete(context.TODO(), "foo", "1-foo")
			testy.ErrorRE(t, "has rev: 2-lkj", err)
		},
	})
	tests.Add("wrong options", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectDelete().WithOptions(map[string]interface{}{"foo": "baz"})
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").Delete(context.TODO(), "foo", "1-foo")
			testy.ErrorRE(t, `has options: map\[foo:baz]`, err)
		},
	})
	tests.Add("success", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectDelete().WillReturn("2-uio")
		},
		test: func(t *testing.T, c *kivik.Client) {
			rev, err := c.DB(context.TODO(), "foo").Delete(context.TODO(), "foo", "1-foo")
			testy.Error(t, "", err)
			if rev != "2-uio" {
				t.Errorf("Unexpected rev: %s", rev)
			}
		},
	})
	tests.Run(t, testMock)
}

func TestCopy(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectCopy().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").Copy(context.TODO(), "foo", "bar")
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectCopy().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").Copy(newCanceledContext(), "foo", "bar")
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectCopy()
			foo.ExpectCopy()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			_, err := foo.Copy(context.TODO(), "foo", "1-foo")
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Add("wrong targetID", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectCopy().WithTargetID("bar")
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").Copy(context.TODO(), "foo", "bar")
			testy.ErrorRE(t, "has targetID: bar", err)
		},
	})
	tests.Add("wrong sourceID", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectCopy().WithSourceID("baz")
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").Copy(context.TODO(), "foo", "bar")
			testy.ErrorRE(t, "has sourceID: baz", err)
		},
	})
	tests.Add("wrong options", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectCopy().WithOptions(map[string]interface{}{"foo": "baz"})
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").Copy(context.TODO(), "foo", "bar")
			testy.ErrorRE(t, `has options: map\[foo:baz]`, err)
		},
	})
	tests.Add("success", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectCopy().WillReturn("2-oiu")
		},
		test: func(t *testing.T, c *kivik.Client) {
			rev, err := c.DB(context.TODO(), "foo").Copy(context.TODO(), "foo", "bar")
			testy.Error(t, "", err)
			if rev != "2-oiu" {
				t.Errorf("Unexpected rev: %s", rev)
			}
		},
	})
	tests.Run(t, testMock)
}

func TestGet(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectGet().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").Get(context.TODO(), "foo").Err
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectGet().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").Get(newCanceledContext(), "foo").Err
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectGet()
			foo.ExpectGet()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			err := foo.Get(context.TODO(), "foo").Err
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Add("wrong docID", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectGet().WithDocID("bar")
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").Get(context.TODO(), "foo").Err
			testy.ErrorRE(t, "has docID: bar", err)
		},
	})
	tests.Add("wrong options", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectGet().WithOptions(map[string]interface{}{"foo": "baz"})
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").Get(context.TODO(), "foo").Err
			testy.ErrorRE(t, `has options: map\[foo:baz]`, err)
		},
	})
	tests.Add("success", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectGet().WillReturn(&driver.Document{Rev: "2-bar"})
		},
		test: func(t *testing.T, c *kivik.Client) {
			row := c.DB(context.TODO(), "foo").Get(context.TODO(), "foo")
			testy.Error(t, "", row.Err)
			if rev := row.Rev; rev != "2-bar" {
				t.Errorf("Unexpected rev: %s", rev)
			}
		},
	})
	tests.Run(t, testMock)
}

func TestGetAttachmentMeta(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectGetAttachmentMeta().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").GetAttachmentMeta(context.TODO(), "foo", "foo.txt")
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectGetAttachmentMeta().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").GetAttachmentMeta(newCanceledContext(), "foo", "foo.txt")
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectGetAttachmentMeta()
			foo.ExpectGetAttachmentMeta()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			_, err := foo.GetAttachmentMeta(context.TODO(), "foo", "foo.txt")
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Add("wrong docID", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectGetAttachmentMeta().WithDocID("bar")
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").GetAttachmentMeta(context.TODO(), "foo", "foo.txt")
			testy.ErrorRE(t, "has docID: bar", err)
		},
	})
	tests.Add("wrong filename", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectGetAttachmentMeta().WithFilename("bar.jpg")
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").GetAttachmentMeta(context.TODO(), "foo", "foo.txt")
			testy.ErrorRE(t, "has filename: bar.jpg", err)
		},
	})
	tests.Add("wrong options", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectGetAttachmentMeta().WithOptions(map[string]interface{}{"foo": "baz"})
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").GetAttachmentMeta(context.TODO(), "foo", "foo.txt")
			testy.ErrorRE(t, `has options: map\[foo:baz]`, err)
		},
	})
	tests.Add("success", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectGetAttachmentMeta().WillReturn(&driver.Attachment{Filename: "foo.txt"})
		},
		test: func(t *testing.T, c *kivik.Client) {
			att, err := c.DB(context.TODO(), "foo").GetAttachmentMeta(context.TODO(), "foo", "foo.txt")
			testy.Error(t, "", err)
			if filename := att.Filename; filename != "foo.txt" {
				t.Errorf("Unexpected filename: %s", filename)
			}
		},
	})
	tests.Run(t, testMock)
}

func TestLocalDocs(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectLocalDocs().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.LocalDocs(context.TODO(), nil)
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("success", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectLocalDocs().WillReturn(NewRows().
				AddRow(&driver.Row{ID: "foo"}).
				AddRow(&driver.Row{ID: "bar"}).
				AddRow(&driver.Row{ID: "baz"}))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			rows, err := db.LocalDocs(context.TODO())
			testy.Error(t, "", err)
			ids := []string{}
			for rows.Next() {
				ids = append(ids, rows.ID())
			}
			expected := []string{"foo", "bar", "baz"}
			if d := diff.Interface(expected, ids); d != nil {
				t.Error(d)
			}
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectLocalDocs().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.LocalDocs(newCanceledContext())
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectLocalDocs()
			foo.ExpectLocalDocs()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			_, err := foo.LocalDocs(context.TODO())
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Run(t, testMock)
}

func TestPurge(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectPurge().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.Purge(context.TODO(), nil)
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("success", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectPurge().WillReturn(&driver.PurgeResult{Seq: 123})
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			result, err := db.Purge(context.TODO(), nil)
			testy.Error(t, "", err)
			if seq := result.Seq; seq != 123 {
				t.Errorf("Unexpected seq: %v", seq)
			}
		},
	})
	tests.Add("wrong map", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectPurge().WithDocRevMap(map[string][]string{"foo": {"a", "b"}})
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.Purge(context.TODO(), nil)
			testy.ErrorRE(t, "has docRevMap: map", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectPurge().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.Purge(newCanceledContext(), nil)
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectPurge()
			foo.ExpectPurge()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			_, err := foo.Purge(context.TODO(), nil)
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Run(t, testMock)
}

func TestPutAttachment(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectPutAttachment().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").PutAttachment(context.TODO(), "foo", "1-foo", &kivik.Attachment{Filename: "foo.txt"})
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectPutAttachment().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").PutAttachment(newCanceledContext(), "foo", "1-foo", &kivik.Attachment{Filename: "foo.txt"})
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectPutAttachment()
			foo.ExpectPutAttachment()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			_, err := foo.PutAttachment(context.TODO(), "foo", "1-foo", &kivik.Attachment{Filename: "foo.txt"})
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Add("wrong id", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectPutAttachment().WithDocID("bar")
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").PutAttachment(context.TODO(), "foo", "1-foo", &kivik.Attachment{Filename: "foo.txt"})
			testy.ErrorRE(t, "has docID: bar", err)
		},
	})
	tests.Add("wrong rev", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectPutAttachment().WithRev("2-bar")
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").PutAttachment(context.TODO(), "foo", "1-foo", &kivik.Attachment{Filename: "foo.txt"})
			testy.ErrorRE(t, "has rev: 2-bar", err)
		},
	})
	tests.Add("wrong attachment", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectPutAttachment().WithAttachment(&driver.Attachment{Filename: "bar.jpg"})
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").PutAttachment(context.TODO(), "foo", "1-foo", &kivik.Attachment{Filename: "foo.txt"})
			testy.ErrorRE(t, "has attachment: bar.jpg", err)
		},
	})
	tests.Add("wrong options", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectPutAttachment().WithOptions(map[string]interface{}{"foo": "bar"})
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").PutAttachment(context.TODO(), "foo", "1-foo", &kivik.Attachment{Filename: "foo.txt"}, map[string]interface{}{"foo": 123})
			testy.ErrorRE(t, "has docID: foo", err)
		},
	})
	tests.Add("success", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectPutAttachment().WillReturn("2-boo")
		},
		test: func(t *testing.T, c *kivik.Client) {
			result, err := c.DB(context.TODO(), "foo").PutAttachment(context.TODO(), "foo", "1-foo", &kivik.Attachment{Filename: "foo.txt"})
			testy.ErrorRE(t, "", err)
			if result != "2-boo" {
				t.Errorf("Unexpected result: %s", result)
			}
		},
	})
	tests.Run(t, testMock)
}

func TestQuery(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectQuery().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.Query(context.TODO(), "foo", "bar")
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("success", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectQuery().WillReturn(NewRows().
				AddRow(&driver.Row{ID: "foo"}).
				AddRow(&driver.Row{ID: "bar"}).
				AddRow(&driver.Row{ID: "baz"}))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			rows, err := db.Query(context.TODO(), "foo", "bar")
			testy.Error(t, "", err)
			ids := []string{}
			for rows.Next() {
				ids = append(ids, rows.ID())
			}
			expected := []string{"foo", "bar", "baz"}
			if d := diff.Interface(expected, ids); d != nil {
				t.Error(d)
			}
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectQuery().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.Query(newCanceledContext(), "foo", "bar")
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectQuery()
			foo.ExpectQuery()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			_, err := foo.Query(context.TODO(), "foo", "bar")
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Add("wrong ddocID", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectQuery().WithDDocID("bar")
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.Query(context.TODO(), "foo", "bar")
			testy.ErrorRE(t, "has ddocID: bar", err)
		},
	})
	tests.Add("wrong view", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectQuery().WithView("baz")
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.Query(context.TODO(), "foo", "bar")
			testy.ErrorRE(t, "has view: baz", err)
		},
	})
	tests.Run(t, testMock)
}

var (
	driverSec = &driver.Security{Admins: driver.Members{Names: []string{"bob"}}}
	clientSec = &kivik.Security{Admins: kivik.Members{Names: []string{"bob"}}}
)

func TestSecurity(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectSecurity().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.Security(context.TODO())
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("success", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectSecurity().WillReturn(driverSec)
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			result, err := db.Security(context.TODO())
			testy.Error(t, "", err)
			if d := diff.Interface(clientSec, result); d != nil {
				t.Error(d)
			}
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectSecurity().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.Security(newCanceledContext())
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectSecurity()
			foo.ExpectSecurity()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			_, err := foo.Security(context.TODO())
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Run(t, testMock)
}

func TestSetSecurity(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectSetSecurity().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			err := db.SetSecurity(context.TODO(), clientSec)
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("success", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectSetSecurity().WithSecurity(driverSec)
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			err := db.SetSecurity(context.TODO(), clientSec)
			testy.Error(t, "", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectSetSecurity().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			err := db.SetSecurity(newCanceledContext(), clientSec)
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectSetSecurity()
			foo.ExpectSetSecurity()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			err := foo.SetSecurity(context.TODO(), clientSec)
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Run(t, testMock)
}

func TestStats(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectStats().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.Stats(context.TODO())
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("success", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectStats().WillReturn(&driver.DBStats{Name: "foo"})
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			result, err := db.Stats(context.TODO())
			testy.Error(t, "", err)
			expected := &kivik.DBStats{Name: "foo"}
			if d := diff.Interface(expected, result); d != nil {
				t.Error(d)
			}
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectStats().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.Stats(newCanceledContext())
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectStats()
			foo.ExpectStats()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			_, err := foo.Stats(context.TODO())
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Run(t, testMock)
}

func TestBulkDocs(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectBulkDocs().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.BulkDocs(context.TODO(), []interface{}{1})
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("unexpected", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.BulkDocs(context.TODO(), []interface{}{1})
			testy.Error(t, "call to DB.BulkDocs() was not expected, all expectations already fulfilled", err)
		},
	})
	tests.Add("rows close error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectBulkDocs().WillReturn(NewBulkResults().CloseError(errors.New("bar err")))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			rows, err := db.BulkDocs(context.TODO(), []interface{}{1})
			testy.Error(t, "", err)
			testy.Error(t, "bar err", rows.Close())
		},
	})
	tests.Add("results", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectBulkDocs().WillReturn(NewBulkResults().
				AddResult(&driver.BulkResult{ID: "foo"}).
				AddResult(&driver.BulkResult{ID: "bar"}).
				AddResult(&driver.BulkResult{ID: "baz"}))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			rows, err := db.BulkDocs(context.TODO(), []interface{}{1})
			testy.Error(t, "", err)
			ids := []string{}
			for rows.Next() {
				ids = append(ids, rows.ID())
			}
			expected := []string{"foo", "bar", "baz"}
			if d := diff.Interface(expected, ids); d != nil {
				t.Error(d)
			}
		},
	})
	tests.Add("result error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectBulkDocs().WillReturn(NewBulkResults().
				AddResult(&driver.BulkResult{ID: "foo"}).
				AddResultError(errors.New("foo err")))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			rows, err := db.BulkDocs(context.TODO(), []interface{}{1})
			testy.Error(t, "", err)
			ids := []string{}
			for rows.Next() {
				ids = append(ids, rows.ID())
			}
			expected := []string{"foo"}
			if d := diff.Interface(expected, ids); d != nil {
				t.Error(d)
			}
			testy.Error(t, "foo err", rows.Err())
		},
	})
	tests.Add("options", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectBulkDocs().WithOptions(map[string]interface{}{"foo": 123})
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.BulkDocs(context.TODO(), []interface{}{1})
			testy.ErrorRE(t, `map\[foo:123]`, err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectBulkDocs().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.BulkDocs(newCanceledContext(), []interface{}{1})
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("result delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectBulkDocs().WillReturn(NewBulkResults().
				AddDelay(time.Millisecond).
				AddResult(&driver.BulkResult{ID: "foo"}).
				AddDelay(time.Second).
				AddResult(&driver.BulkResult{ID: "bar"}))
		},
		test: func(t *testing.T, c *kivik.Client) {
			ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
			defer cancel()
			rows, err := c.DB(ctx, "foo").BulkDocs(ctx, []interface{}{1})
			testy.Error(t, "", err)
			ids := []string{}
			for rows.Next() {
				ids = append(ids, rows.ID())
			}
			expected := []string{"foo"}
			if d := diff.Interface(expected, ids); d != nil {
				t.Error(d)
			}
			testy.Error(t, "context deadline exceeded", rows.Err())
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectBulkDocs()
			foo.ExpectBulkDocs()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			_, err := foo.BulkDocs(context.TODO(), []interface{}{1})
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Run(t, testMock)
}

func TestGetAttachment(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectGetAttachment().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").GetAttachment(context.TODO(), "foo", "bar")
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectGetAttachment().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").GetAttachment(newCanceledContext(), "foo", "bar")
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectGetAttachment()
			foo.ExpectGetAttachment()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			_, err := foo.GetAttachment(context.TODO(), "foo", "bar")
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Add("wrong docID", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectGetAttachment().WithDocID("bar")
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").GetAttachment(context.TODO(), "foo", "bar")
			testy.ErrorRE(t, "has docID: bar", err)
		},
	})
	tests.Add("success", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectGetAttachment().WillReturn(&driver.Attachment{Filename: "foo.txt"})
		},
		test: func(t *testing.T, c *kivik.Client) {
			att, err := c.DB(context.TODO(), "foo").GetAttachment(context.TODO(), "foo", "bar")
			testy.Error(t, "", err)
			if name := att.Filename; name != "foo.txt" {
				t.Errorf("Unexpected filename: %s", name)
			}
		},
	})
	tests.Run(t, testMock)
}

func TestDesignDocs(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectDesignDocs().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.DesignDocs(context.TODO(), nil)
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("success", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectDesignDocs().WillReturn(NewRows().
				AddRow(&driver.Row{ID: "foo"}).
				AddRow(&driver.Row{ID: "bar"}).
				AddRow(&driver.Row{ID: "baz"}))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			rows, err := db.DesignDocs(context.TODO())
			testy.Error(t, "", err)
			ids := []string{}
			for rows.Next() {
				ids = append(ids, rows.ID())
			}
			expected := []string{"foo", "bar", "baz"}
			if d := diff.Interface(expected, ids); d != nil {
				t.Error(d)
			}
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectDesignDocs().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.DesignDocs(newCanceledContext())
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectDesignDocs()
			foo.ExpectDesignDocs()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			_, err := foo.DesignDocs(context.TODO())
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Run(t, testMock)
}

func TestChanges(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectChanges().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.Changes(context.TODO())
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("unexpected", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.Changes(context.TODO())
			testy.Error(t, "call to DB.Changes() was not expected, all expectations already fulfilled", err)
		},
	})
	tests.Add("close error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectChanges().WillReturn(NewChanges().CloseError(errors.New("bar err")))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			rows, err := db.Changes(context.TODO())
			testy.Error(t, "", err)
			testy.Error(t, "bar err", rows.Close())
		},
	})
	tests.Add("changes", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectChanges().WillReturn(NewChanges().
				AddChange(&driver.Change{ID: "foo"}).
				AddChange(&driver.Change{ID: "bar"}).
				AddChange(&driver.Change{ID: "baz"}))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			rows, err := db.Changes(context.TODO())
			testy.Error(t, "", err)
			ids := []string{}
			for rows.Next() {
				ids = append(ids, rows.ID())
			}
			expected := []string{"foo", "bar", "baz"}
			if d := diff.Interface(expected, ids); d != nil {
				t.Error(d)
			}
		},
	})
	tests.Add("row error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectChanges().WillReturn(NewChanges().
				AddChange(&driver.Change{ID: "foo"}).
				AddChangeError(errors.New("foo err")))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			rows, err := db.Changes(context.TODO())
			testy.Error(t, "", err)
			ids := []string{}
			for rows.Next() {
				ids = append(ids, rows.ID())
			}
			expected := []string{"foo"}
			if d := diff.Interface(expected, ids); d != nil {
				t.Error(d)
			}
			testy.Error(t, "foo err", rows.Err())
		},
	})
	tests.Add("options", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectChanges().WithOptions(map[string]interface{}{"foo": 123})
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.Changes(context.TODO())
			testy.ErrorRE(t, `map\[foo:123]`, err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectChanges().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.Changes(newCanceledContext())
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("change delay", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectChanges().WillReturn(NewChanges().
				AddDelay(time.Millisecond).
				AddChange(&driver.Change{ID: "foo"}).
				AddDelay(time.Second).
				AddChange(&driver.Change{ID: "bar"}))
		},
		test: func(t *testing.T, c *kivik.Client) {
			ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
			defer cancel()
			rows, err := c.DB(ctx, "foo").Changes(ctx)
			testy.Error(t, "", err)
			ids := []string{}
			for rows.Next() {
				ids = append(ids, rows.ID())
			}
			expected := []string{"foo"}
			if d := diff.Interface(expected, ids); d != nil {
				t.Error(d)
			}
			testy.Error(t, "context deadline exceeded", rows.Err())
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *Client) {
			foo := m.NewDB()
			bar := m.NewDB()
			m.ExpectDB().WithName("foo").WillReturn(foo)
			m.ExpectDB().WithName("bar").WillReturn(bar)
			bar.ExpectChanges()
			foo.ExpectChanges()
		},
		test: func(t *testing.T, c *kivik.Client) {
			foo := c.DB(context.TODO(), "foo")
			_ = c.DB(context.TODO(), "bar")
			_, err := foo.Changes(context.TODO())
			testy.ErrorRE(t, `Expected: call to DB\(bar`, err)
		},
		err: "there is a remaining unmet expectation: call to DB().Close()",
	})
	tests.Add("changes last_seq", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectChanges().WillReturn(NewChanges().LastSeq("1-asdf"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			ch, err := db.Changes(context.TODO())
			testy.Error(t, "", err)
			if o := ch.LastSeq(); o != "1-asdf" {
				t.Errorf("Unexpected last_seq: %s", o)
			}
		},
	})
	tests.Add("changes pending", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectChanges().WillReturn(NewChanges().Pending(123))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			ch, err := db.Changes(context.TODO())
			testy.Error(t, "", err)
			if o := ch.Pending(); o != 123 {
				t.Errorf("Unexpected pending: %d", o)
			}
		},
	})
	tests.Add("changes etag", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectChanges().WillReturn(NewChanges().ETag("etag-foo"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			ch, err := db.Changes(context.TODO())
			testy.Error(t, "", err)
			if o := ch.ETag(); o != "etag-foo" {
				t.Errorf("Unexpected pending: %s", o)
			}
		},
	})
	tests.Run(t, testMock)
}

func TestRevsDiff(t *testing.T) {
	revMap := map[string]interface{}{"foo": []string{"1", "2"}}
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectRevsDiff().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.RevsDiff(context.TODO(), revMap)
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("success", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectRevsDiff().
				WithRevLookup(revMap).
				WillReturn(map[string]driver.RevDiff{
					"foo": {
						Missing:           []string{"1"},
						PossibleAncestors: []string{"2"},
					},
				})
		},
		test: func(t *testing.T, c *kivik.Client) {
			expected := kivik.Diffs{
				"foo": kivik.RevDiff{
					Missing:           []string{"1"},
					PossibleAncestors: []string{"2"},
				},
			}
			db := c.DB(context.TODO(), "foo")
			result, err := db.RevsDiff(context.TODO(), revMap)
			testy.Error(t, "", err)
			if d := diff.Interface(expected, result); d != nil {
				t.Error(d)
			}
		},
	})

	tests.Run(t, testMock)
}
