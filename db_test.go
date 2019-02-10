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
		setup: func(m *MockClient) {
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
		setup: func(m *MockClient) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").Close(context.TODO())
			testy.Error(t, "call to DB.Close() was not expected, all expectations already fulfilled", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *MockClient) {
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
		setup: func(m *MockClient) {
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
	tests.Run(t, testMock)
}

func TestAllDocs(t *testing.T) { // nolint: gocyclo
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *MockClient) {
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
		setup: func(m *MockClient) {
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
		setup: func(m *MockClient) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectAllDocs().WillReturnRows(db.NewRows().CloseError(errors.New("bar err")))
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			rows, err := db.AllDocs(context.TODO())
			testy.Error(t, "", err)
			testy.Error(t, "bar err", rows.Close())
		},
	})
	tests.Add("rows offset", mockTest{
		setup: func(m *MockClient) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectAllDocs().WillReturnRows(db.NewRows().Offset(123))
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
		setup: func(m *MockClient) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectAllDocs().WillReturnRows(db.NewRows().TotalRows(123))
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
		setup: func(m *MockClient) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectAllDocs().WillReturnRows(db.NewRows().UpdateSeq("1-xxx"))
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
		setup: func(m *MockClient) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectAllDocs().WillReturnRows(db.NewRows().Warning("Caution!"))
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
		setup: func(m *MockClient) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectAllDocs().WillReturnRows(db.NewRows().
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
		setup: func(m *MockClient) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectAllDocs().WillReturnRows(db.NewRows().
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
		setup: func(m *MockClient) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectAllDocs().WithOptions(map[string]interface{}{"foo": 123}).
				WillReturnRows(db.NewRows())
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.AllDocs(context.TODO())
			testy.ErrorRE(t, `map\[foo:123]`, err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *MockClient) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectAllDocs().WillDelay(time.Second).
				WillReturnRows(db.NewRows())
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.AllDocs(newCanceledContext())
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("row delay", mockTest{
		setup: func(m *MockClient) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectAllDocs().WillReturnRows(db.NewRows().
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
		setup: func(m *MockClient) {
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
		setup: func(m *MockClient) {
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
		setup: func(m *MockClient) {
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
		setup: func(m *MockClient) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectBulkGet().WillReturnRows(db.NewRows().
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
		setup: func(m *MockClient) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectBulkGet().WithOptions(map[string]interface{}{"foo": 123}).
				WillReturnRows(db.NewRows())
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.BulkGet(context.TODO(), []kivik.BulkGetReference{})
			testy.ErrorRE(t, `map\[foo:123]`, err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *MockClient) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectBulkGet().WillDelay(time.Second).
				WillReturnRows(db.NewRows())
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.BulkGet(newCanceledContext(), []kivik.BulkGetReference{})
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *MockClient) {
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

func TestFind(t *testing.T) { // nolint: gocyclo
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *MockClient) {
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
		setup: func(m *MockClient) {
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
		setup: func(m *MockClient) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectFind().WillReturnRows(db.NewRows().
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
		setup: func(m *MockClient) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectFind().WithQuery(map[string]interface{}{"foo": "123"}).
				WillReturnRows(db.NewRows())
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.Find(context.TODO(), map[string]string{"foo": "123"})
			testy.ErrorRE(t, "", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *MockClient) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectFind().WillDelay(time.Second).
				WillReturnRows(db.NewRows())
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			_, err := db.Find(newCanceledContext(), 0)
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("wrong db", mockTest{
		setup: func(m *MockClient) {
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
		setup: func(m *MockClient) {
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
		setup: func(m *MockClient) {
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
		setup: func(m *MockClient) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectCreateIndex().WithDDoc("moo")
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").CreateIndex(context.TODO(), "foo", "bar", 123)
			testy.ErrorRE(t, "has ddoc: moo", err)
		},
	})
	tests.Add("name", mockTest{
		setup: func(m *MockClient) {
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
		setup: func(m *MockClient) {
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
		setup: func(m *MockClient) {
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
		setup: func(m *MockClient) {
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
		setup: func(m *MockClient) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectGetIndexes().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").GetIndexes(context.TODO())
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("indexes", func() interface{} {
		expected := []kivik.Index{
			{Name: "foo"},
			{Name: "bar"},
		}
		return mockTest{
			setup: func(m *MockClient) {
				db := m.NewDB()
				m.ExpectDB().WillReturn(db)
				db.ExpectGetIndexes().WillReturn(expected)
			},
			test: func(t *testing.T, c *kivik.Client) {
				indexes, err := c.DB(context.TODO(), "foo").GetIndexes(context.TODO())
				testy.Error(t, "", err)
				if d := diff.Interface(expected, indexes); d != nil {
					t.Error(d)
				}
			},
		}
	})
	tests.Add("unexpected", mockTest{
		setup: func(m *MockClient) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").GetIndexes(context.TODO())
			testy.Error(t, "call to DB.GetIndexes() was not expected, all expectations already fulfilled", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *MockClient) {
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
		setup: func(m *MockClient) {
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
		setup: func(m *MockClient) {
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
		setup: func(m *MockClient) {
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
		setup: func(m *MockClient) {
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
		setup: func(m *MockClient) {
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
		setup: func(m *MockClient) {
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
		setup: func(m *MockClient) {
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
		setup: func(m *MockClient) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").Explain(context.TODO(), "foo")
			testy.Error(t, "call to DB.Explain() was not expected, all expectations already fulfilled", err)
		},
	})
	tests.Add("query", mockTest{
		setup: func(m *MockClient) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectExplain().WithQuery(map[string]string{"foo": "bar"})
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DB(context.TODO(), "foo").Explain(context.TODO(), map[string]interface{}{"foo": "bar"})
			testy.Error(t, "", err)
		},
	})
	tests.Add("plan", func() interface{} {
		expected := &kivik.QueryPlan{DBName: "foo"}
		return mockTest{
			setup: func(m *MockClient) {
				db := m.NewDB()
				m.ExpectDB().WillReturn(db)
				db.ExpectExplain().WillReturn(expected)
			},
			test: func(t *testing.T, c *kivik.Client) {
				plan, err := c.DB(context.TODO(), "foo").Explain(context.TODO(), map[string]interface{}{"foo": "bar"})
				testy.Error(t, "", err)
				if d := diff.Interface(expected, plan); d != nil {
					t.Error(d)
				}
			},
		}
	})
	tests.Add("delay", mockTest{
		setup: func(m *MockClient) {
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
		setup: func(m *MockClient) {
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
