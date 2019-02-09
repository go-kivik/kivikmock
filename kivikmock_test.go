package kivikmock

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/flimzy/diff"
	"github.com/flimzy/testy"
	"github.com/go-kivik/couchdb"
	"github.com/go-kivik/kivik"
)

type mockTest struct {
	setup func(Mock)
	test  func(*testing.T, *kivik.Client)
	err   string
}

func testMock(t *testing.T, test mockTest) {
	client, mock, err := New()
	if err != nil {
		t.Fatal("error creating mock database")
	}
	defer client.Close(context.TODO()) // nolint: errcheck
	if test.setup != nil {
		test.setup(mock)
	}
	if test.test != nil {
		test.test(t, client)
	}
	err = mock.ExpectationsWereMet()
	testy.ErrorRE(t, test.err, err)
}

func TestCloseClient(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("err", mockTest{
		setup: func(m Mock) {
			m.ExpectClose().WillReturnError(errors.New("close failed"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.Close(context.TODO())
			testy.Error(t, "close failed", err)
		},
		err: "",
	})
	tests.Add("unexpected", mockTest{
		test: func(t *testing.T, c *kivik.Client) {
			err := c.Close(context.TODO())
			testy.Error(t, "call to Close() was not expected, all expectations already fulfilled", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m Mock) {
			m.ExpectClose().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.Close(newCanceledContext())
			testy.Error(t, "context canceled", err)
		},
	})

	tests.Run(t, testMock)
}

func TestAllDBs(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m Mock) {
			m.ExpectAllDBs().WillReturnError(fmt.Errorf("AllDBs failed"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.AllDBs(context.TODO())
			testy.Error(t, "AllDBs failed", err)
		},
	})
	tests.Add("unexpected", mockTest{
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.AllDBs(context.TODO())
			testy.Error(t, "call to AllDBs() was not expected, all expectations already fulfilled", err)
		},
	})
	tests.Add("success", func() interface{} {
		expected := []string{"a", "b", "c"}
		return mockTest{
			setup: func(m Mock) {
				m.ExpectAllDBs().WillReturn(expected)
			},
			test: func(t *testing.T, c *kivik.Client) {
				result, err := c.AllDBs(context.TODO())
				testy.Error(t, "", err)
				if d := diff.Interface(expected, result); d != nil {
					t.Error(d)
				}
			},
		}
	})
	tests.Add("delay", mockTest{
		setup: func(m Mock) {
			m.ExpectAllDBs().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.AllDBs(newCanceledContext())
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("options", mockTest{
		setup: func(m Mock) {
			m.ExpectAllDBs().WithOptions(map[string]interface{}{"foo": 123})
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.AllDBs(context.TODO(), map[string]interface{}{"foo": 123})
			testy.Error(t, "", err)
		},
	})
	tests.Run(t, testMock)
}

func TestAuthenticate(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m Mock) {
			m.ExpectAuthenticate().WillReturnError(errors.New("auth error"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.Authenticate(context.TODO(), couchdb.BasicAuth("foo", "bar"))
			testy.Error(t, "auth error", err)
		},
	})
	tests.Add("wrong authenticator", mockTest{
		setup: func(m Mock) {
			m.ExpectAuthenticate().WithAuthenticator(int(3))
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.Authenticate(context.TODO(), couchdb.CookieAuth("foo", "bar"))
			expected := `Expectation not met:
Expected: call to Authenticate() which:
	- has an authenticator of type: int
  Actual: call to Authenticate() which:
	- has an authenticator of type: authFunc`
			testy.Error(t, expected, err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m Mock) {
			m.ExpectAuthenticate().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.Authenticate(newCanceledContext(), int(1))
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Run(t, testMock)
}

func TestClusterSetup(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m Mock) {
			m.ExpectClusterSetup().WillReturnError(errors.New("setup error"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.ClusterSetup(context.TODO(), 123)
			testy.Error(t, "setup error", err)
		},
	})
	tests.Add("action", mockTest{
		setup: func(m Mock) {
			m.ExpectClusterSetup().WithAction(123)
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.ClusterSetup(context.TODO(), 123)
			testy.Error(t, "", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m Mock) {
			m.ExpectClusterSetup().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.ClusterSetup(newCanceledContext(), 123)
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("unexpected", mockTest{
		test: func(t *testing.T, c *kivik.Client) {
			err := c.ClusterSetup(context.TODO(), 123)
			testy.Error(t, "call to ClusterSetup() was not expected, all expectations already fulfilled", err)
		},
	})
	tests.Run(t, testMock)
}

func TestClusterStatus(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m Mock) {
			m.ExpectClusterStatus().WillReturnError(errors.New("status error"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.ClusterStatus(context.TODO())
			testy.Error(t, "status error", err)
		},
	})
	tests.Add("options", mockTest{
		setup: func(m Mock) {
			m.ExpectClusterStatus().WithOptions(map[string]interface{}{"foo": 123})
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.ClusterStatus(context.TODO())
			testy.ErrorRE(t, `map\[foo:123]`, err)
		},
	})
	tests.Add("success", func() interface{} {
		const expected = "oink"
		return mockTest{
			setup: func(m Mock) {
				m.ExpectClusterStatus().WillReturn(expected)
			},
			test: func(t *testing.T, c *kivik.Client) {
				result, err := c.ClusterStatus(context.TODO())
				testy.Error(t, "", err)
				if result != expected {
					t.Errorf("Unexpected result: %s", result)
				}
			},
		}
	})
	tests.Add("delay", mockTest{
		setup: func(m Mock) {
			m.ExpectClusterStatus().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.ClusterStatus(newCanceledContext())
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("unordered", mockTest{
		setup: func(m Mock) {
			m.ExpectClose()
			m.ExpectClusterStatus()
			m.MatchExpectationsInOrder(false)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.ClusterStatus(context.TODO())
			testy.Error(t, "", err)
		},
		err: "there is a remaining unmet expectation: call to Close()",
	})
	tests.Add("unexpected", mockTest{
		setup: func(m Mock) {
			m.ExpectClose()
			m.MatchExpectationsInOrder(false)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.ClusterStatus(context.TODO())
			testy.Error(t, "call to ClusterStatus(ctx, ?) was not expected", err)
		},
	})
	tests.Run(t, testMock)
}

func TestDBExists(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m Mock) {
			m.ExpectDBExists().WillReturnError(errors.New("existence error"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DBExists(context.TODO(), "foo")
			testy.Error(t, "existence error", err)
		},
	})
	tests.Add("name", mockTest{
		setup: func(m Mock) {
			m.ExpectDBExists().WithName("foo")
		},
		test: func(t *testing.T, c *kivik.Client) {
			exists, err := c.DBExists(context.TODO(), "foo")
			testy.Error(t, "", err)
			if exists {
				t.Errorf("DB shouldn't exist")
			}
		},
	})
	tests.Add("options", mockTest{
		setup: func(m Mock) {
			m.ExpectDBExists().WithOptions(map[string]interface{}{"foo": 123})
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DBExists(context.TODO(), "foo")
			testy.ErrorRE(t, `map\[foo:123]`, err)
		},
	})
	tests.Add("exists", mockTest{
		setup: func(m Mock) {
			m.ExpectDBExists().WillReturn(true)
		},
		test: func(t *testing.T, c *kivik.Client) {
			exists, err := c.DBExists(context.TODO(), "foo")
			testy.ErrorRE(t, "", err)
			if !exists {
				t.Errorf("DB should exist")
			}
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m Mock) {
			m.ExpectDBExists().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DBExists(newCanceledContext(), "foo")
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Run(t, testMock)
}

func TestDestroyDB(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m Mock) {
			m.ExpectDestroyDB().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DestroyDB(newCanceledContext(), "foo")
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("name", mockTest{
		setup: func(m Mock) {
			m.ExpectDestroyDB().WithName("foo")
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DestroyDB(newCanceledContext(), "foo")
			testy.Error(t, "", err)
		},
	})
	tests.Add("options", mockTest{
		setup: func(m Mock) {
			m.ExpectDestroyDB().WithOptions(kivik.Options{"foo": 123})
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DestroyDB(newCanceledContext(), "foo")
			testy.ErrorRE(t, `map\[foo:123]`, err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m Mock) {
			m.ExpectDestroyDB().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DestroyDB(newCanceledContext(), "foo")
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Run(t, testMock)
}
