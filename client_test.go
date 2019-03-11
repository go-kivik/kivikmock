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
	"github.com/go-kivik/kivik/driver"
)

type mockTest struct {
	setup func(*Client)
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
		setup: func(m *Client) {
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
		setup: func(m *Client) {
			m.ExpectClose().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.Close(newCanceledContext())
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("callback", mockTest{
		setup: func(m *Client) {
			m.ExpectClose().WillExecute(func(_ context.Context) error {
				return errors.New("custom error")
			})
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.Close(context.TODO())
			testy.Error(t, "custom error", err)
		},
	})
	tests.Run(t, testMock)
}

func TestAllDBs(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
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
			setup: func(m *Client) {
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
		setup: func(m *Client) {
			m.ExpectAllDBs().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.AllDBs(newCanceledContext())
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("options", mockTest{
		setup: func(m *Client) {
			m.ExpectAllDBs().WithOptions(map[string]interface{}{"foo": 123})
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.AllDBs(context.TODO(), map[string]interface{}{"foo": 123})
			testy.Error(t, "", err)
		},
	})
	tests.Add("callback", mockTest{
		setup: func(m *Client) {
			m.ExpectAllDBs().WillExecute(func(_ context.Context, _ map[string]interface{}) ([]string, error) {
				return nil, errors.New("custom error")
			})
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.AllDBs(context.TODO())
			testy.Error(t, "custom error", err)
		},
	})
	tests.Run(t, testMock)
}

func TestAuthenticate(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			m.ExpectAuthenticate().WillReturnError(errors.New("auth error"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.Authenticate(context.TODO(), couchdb.BasicAuth("foo", "bar"))
			testy.Error(t, "auth error", err)
		},
	})
	tests.Add("wrong authenticator", mockTest{
		setup: func(m *Client) {
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
		setup: func(m *Client) {
			m.ExpectAuthenticate().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.Authenticate(newCanceledContext(), int(1))
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("callback", mockTest{
		setup: func(m *Client) {
			m.ExpectAuthenticate().WillExecute(func(_ context.Context, _ interface{}) error {
				return errors.New("custom error")
			})
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.Authenticate(context.TODO(), int(1))
			testy.Error(t, "custom error", err)
		},
	})
	tests.Run(t, testMock)
}

func TestClusterSetup(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			m.ExpectClusterSetup().WillReturnError(errors.New("setup error"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.ClusterSetup(context.TODO(), 123)
			testy.Error(t, "setup error", err)
		},
	})
	tests.Add("action", mockTest{
		setup: func(m *Client) {
			m.ExpectClusterSetup().WithAction(123)
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.ClusterSetup(context.TODO(), 123)
			testy.Error(t, "", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
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
	tests.Add("callback", mockTest{
		setup: func(m *Client) {
			m.ExpectClusterSetup().WillExecute(func(_ context.Context, _ interface{}) error {
				return errors.New("custom error")
			})
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.ClusterSetup(context.TODO(), 123)
			testy.Error(t, "custom error", err)
		},
	})
	tests.Run(t, testMock)
}

func TestClusterStatus(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			m.ExpectClusterStatus().WillReturnError(errors.New("status error"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.ClusterStatus(context.TODO())
			testy.Error(t, "status error", err)
		},
	})
	tests.Add("options", mockTest{
		setup: func(m *Client) {
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
			setup: func(m *Client) {
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
		setup: func(m *Client) {
			m.ExpectClusterStatus().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.ClusterStatus(newCanceledContext())
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("unordered", mockTest{
		setup: func(m *Client) {
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
		setup: func(m *Client) {
			m.ExpectClose()
			m.MatchExpectationsInOrder(false)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.ClusterStatus(context.TODO())
			testy.Error(t, "call to ClusterStatus(ctx, [?]) was not expected", err)
		},
	})
	tests.Add("callback", mockTest{
		setup: func(m *Client) {
			m.ExpectClusterStatus().WillExecute(func(_ context.Context, _ map[string]interface{}) (string, error) {
				return "", errors.New("custom error")
			})
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.ClusterStatus(newCanceledContext())
			testy.Error(t, "custom error", err)
		},
	})
	tests.Run(t, testMock)
}

func TestDBExists(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			m.ExpectDBExists().WillReturnError(errors.New("existence error"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DBExists(context.TODO(), "foo")
			testy.Error(t, "existence error", err)
		},
	})
	tests.Add("name", mockTest{
		setup: func(m *Client) {
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
		setup: func(m *Client) {
			m.ExpectDBExists().WithOptions(map[string]interface{}{"foo": 123})
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DBExists(context.TODO(), "foo")
			testy.ErrorRE(t, `map\[foo:123]`, err)
		},
	})
	tests.Add("exists", mockTest{
		setup: func(m *Client) {
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
		setup: func(m *Client) {
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
		setup: func(m *Client) {
			m.ExpectDestroyDB().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DestroyDB(newCanceledContext(), "foo")
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("name", mockTest{
		setup: func(m *Client) {
			m.ExpectDestroyDB().WithName("foo")
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DestroyDB(newCanceledContext(), "foo")
			testy.Error(t, "", err)
		},
	})
	tests.Add("options", mockTest{
		setup: func(m *Client) {
			m.ExpectDestroyDB().WithOptions(kivik.Options{"foo": 123})
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DestroyDB(newCanceledContext(), "foo")
			testy.ErrorRE(t, `map\[foo:123]`, err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			m.ExpectDestroyDB().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DestroyDB(newCanceledContext(), "foo")
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Run(t, testMock)
}

func TestDBsStats(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			m.ExpectDBsStats().WillReturnError(errors.New("stats error"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DBsStats(context.TODO(), []string{"foo"})
			testy.Error(t, "stats error", err)
		},
	})
	tests.Add("names", mockTest{
		setup: func(m *Client) {
			m.ExpectDBsStats().WithNames([]string{"a", "b"})
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DBsStats(context.TODO(), []string{"foo"})
			testy.ErrorRE(t, "[a b]", err)
		},
	})
	tests.Add("success", func() interface{} {
		return mockTest{
			setup: func(m *Client) {
				m.ExpectDBsStats().WillReturn([]*driver.DBStats{
					{Name: "foo", Cluster: &driver.ClusterStats{Replicas: 5}},
					{Name: "bar"},
				})
			},
			test: func(t *testing.T, c *kivik.Client) {
				result, err := c.DBsStats(context.TODO(), []string{"foo", "bar"})
				testy.ErrorRE(t, "", err)
				expected := []*kivik.DBStats{
					{Name: "foo", Cluster: &kivik.ClusterConfig{Replicas: 5}},
					{Name: "bar"},
				}
				if d := diff.Interface(expected, result); d != nil {
					t.Error(d)
				}
			},
		}
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			m.ExpectDBsStats().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DBsStats(newCanceledContext(), []string{"foo"})
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Run(t, testMock)
}

func TestPing(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("unreachable", mockTest{
		setup: func(m *Client) {
			m.ExpectPing()
		},
		test: func(t *testing.T, c *kivik.Client) {
			reachable, err := c.Ping(context.TODO())
			testy.Error(t, "", err)
			if reachable {
				t.Errorf("Expected db to be unreachable")
			}
		},
	})
	tests.Add("reachable", mockTest{
		setup: func(m *Client) {
			m.ExpectPing().WillReturn(true)
		},
		test: func(t *testing.T, c *kivik.Client) {
			reachable, err := c.Ping(context.TODO())
			testy.Error(t, "", err)
			if !reachable {
				t.Errorf("Expected db to be reachable")
			}
		},
	})
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			m.ExpectPing().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.Ping(context.TODO())
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("unexpected", mockTest{
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.Ping(context.TODO())
			testy.Error(t, "call to Ping() was not expected, all expectations already fulfilled", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			m.ExpectPing().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.Ping(newCanceledContext())
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Run(t, testMock)
}

func TestSession(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("session", func() interface{} {
		return mockTest{
			setup: func(m *Client) {
				m.ExpectSession().WillReturn(&driver.Session{
					Name: "bob",
				})
			},
			test: func(t *testing.T, c *kivik.Client) {
				session, err := c.Session(context.TODO())
				testy.Error(t, "", err)
				expected := &kivik.Session{
					Name: "bob",
				}
				if d := diff.Interface(expected, session); d != nil {
					t.Error(d)
				}
			},
		}
	})
	tests.Add("unexpected", mockTest{
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.Session(context.TODO())
			testy.Error(t, "call to Session() was not expected, all expectations already fulfilled", err)
		},
	})
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			m.ExpectSession().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.Session(context.TODO())
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			m.ExpectSession().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.Session(newCanceledContext())
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Run(t, testMock)
}

func TestVersion(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("version", func() interface{} {
		return mockTest{
			setup: func(m *Client) {
				m.ExpectVersion().WillReturn(&driver.Version{Version: "1.2"})
			},
			test: func(t *testing.T, c *kivik.Client) {
				session, err := c.Version(context.TODO())
				testy.Error(t, "", err)
				expected := &kivik.Version{Version: "1.2"}
				if d := diff.Interface(expected, session); d != nil {
					t.Error(d)
				}
			},
		}
	})
	tests.Add("unexpected", mockTest{
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.Version(context.TODO())
			testy.Error(t, "call to Version() was not expected, all expectations already fulfilled", err)
		},
	})
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			m.ExpectVersion().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.Version(context.TODO())
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			m.ExpectVersion().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.Version(newCanceledContext())
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Run(t, testMock)
}

func TestDB(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("name", mockTest{
		setup: func(m *Client) {
			m.ExpectDB().WithName("foo")
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").Err()
			testy.Error(t, "", err)
		},
	})
	tests.Add("unexpected", mockTest{
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").Err()
			testy.Error(t, "call to DB() was not expected, all expectations already fulfilled", err)
		},
	})
	tests.Add("options", mockTest{
		setup: func(m *Client) {
			m.ExpectDB().WithOptions(map[string]interface{}{"foo": 123})
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo", kivik.Options{"foo": 123}).Err()
			testy.Error(t, "", err)
		},
	})
	tests.Add("success", mockTest{
		setup: func(m *Client) {
			m.ExpectDB().WillReturn(m.NewDB())
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.DB(context.TODO(), "foo")
			err := db.Err()
			testy.Error(t, "", err)
			if db.Name() != "foo" {
				t.Errorf("Unexpected db name: %s", db.Name())
			}
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			m.ExpectDB().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(newCanceledContext(), "foo").Err()
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Run(t, testMock)
}

func TestCreateDB(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			m.ExpectCreateDB().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.CreateDB(context.TODO(), "foo").Err()
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("name", mockTest{
		setup: func(m *Client) {
			m.ExpectCreateDB().WithName("foo")
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.CreateDB(context.TODO(), "foo").Err()
			testy.Error(t, "", err)
		},
	})
	tests.Add("unexpected", mockTest{
		test: func(t *testing.T, c *kivik.Client) {
			err := c.CreateDB(context.TODO(), "foo").Err()
			testy.Error(t, "call to CreateDB() was not expected, all expectations already fulfilled", err)
		},
	})
	tests.Add("options", mockTest{
		setup: func(m *Client) {
			m.ExpectCreateDB().WithOptions(map[string]interface{}{"foo": 123})
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.CreateDB(context.TODO(), "foo", kivik.Options{"foo": 123}).Err()
			testy.Error(t, "", err)
		},
	})
	tests.Add("success", mockTest{
		setup: func(m *Client) {
			m.ExpectCreateDB().WillReturn(m.NewDB())
		},
		test: func(t *testing.T, c *kivik.Client) {
			db := c.CreateDB(context.TODO(), "foo")
			err := db.Err()
			testy.Error(t, "", err)
			if db.Name() != "foo" {
				t.Errorf("Unexpected db name: %s", db.Name())
			}
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			m.ExpectCreateDB().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.CreateDB(newCanceledContext(), "foo").Err()
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("name confusion", mockTest{
		setup: func(m *Client) {
			db := m.NewDB()
			m.ExpectCreateDB().WithName("bundle-foo").WillReturn(db)
			db.ExpectSetSecurity().WithSecurity(&driver.Security{
				Admins: driver.Members{Names: []string{"bob"}},
			}).WillReturnError(errors.New("security fail"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			ctx := context.Background()
			db := c.CreateDB(ctx, "bundle-foo")
			security := &kivik.Security{
				Admins: kivik.Members{
					Names: []string{"bob"},
				},
			}
			err := db.SetSecurity(ctx, security)
			testy.Error(t, "security fail", err)
		},
	})
	tests.Add("cleanup expectations", mockTest{
		setup: func(m *Client) {
			m.ExpectCreateDB().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.CreateDB(context.TODO(), "foo").Err()
			if err == nil {
				t.Fatal("expected error")
			}
		},
	})
	tests.Add("callback", mockTest{
		setup: func(m *Client) {
			m.ExpectCreateDB().WillExecute(func(_ context.Context, _ string, _ map[string]interface{}) error {
				return errors.New("custom error")
			})
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.CreateDB(context.TODO(), "foo").Err()
			testy.Error(t, "custom error", err)
		},
	})
	tests.Run(t, testMock)
}

func TestDBUpdates(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m *Client) {
			m.ExpectDBUpdates().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DBUpdates(context.TODO())
			testy.Error(t, "foo err", err)
		},
	})
	tests.Add("unexpected", mockTest{
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DBUpdates(context.TODO())
			testy.Error(t, "call to DBUpdates() was not expected, all expectations already fulfilled", err)
		},
	})
	tests.Add("close error", mockTest{
		setup: func(m *Client) {
			m.ExpectDBUpdates().WillReturn(NewDBUpdates().CloseError(errors.New("bar err")))
		},
		test: func(t *testing.T, c *kivik.Client) {
			updates, err := c.DBUpdates(context.TODO())
			testy.Error(t, "", err)
			testy.Error(t, "bar err", updates.Close())
		},
	})
	tests.Add("updates", mockTest{
		setup: func(m *Client) {
			m.ExpectDBUpdates().WillReturn(NewDBUpdates().
				AddUpdate(&driver.DBUpdate{DBName: "foo"}).
				AddUpdate(&driver.DBUpdate{DBName: "bar"}).
				AddUpdate(&driver.DBUpdate{DBName: "baz"}))
		},
		test: func(t *testing.T, c *kivik.Client) {
			updates, err := c.DBUpdates(context.TODO())
			testy.Error(t, "", err)
			names := []string{}
			for updates.Next() {
				names = append(names, updates.DBName())
			}
			expected := []string{"foo", "bar", "baz"}
			if d := diff.Interface(expected, names); d != nil {
				t.Error(d)
			}
		},
	})
	tests.Add("iter error", mockTest{
		setup: func(m *Client) {
			m.ExpectDBUpdates().WillReturn(NewDBUpdates().
				AddUpdate(&driver.DBUpdate{DBName: "foo"}).
				AddUpdateError(errors.New("foo err")))
		},
		test: func(t *testing.T, c *kivik.Client) {
			updates, err := c.DBUpdates(context.TODO())
			testy.Error(t, "", err)
			names := []string{}
			for updates.Next() {
				names = append(names, updates.DBName())
			}
			expected := []string{"foo"}
			if d := diff.Interface(expected, names); d != nil {
				t.Error(d)
			}
			testy.Error(t, "foo err", updates.Err())
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m *Client) {
			m.ExpectDBUpdates().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			_, err := c.DBUpdates(newCanceledContext())
			testy.Error(t, "context canceled", err)
		},
	})
	tests.Add("update delay", mockTest{
		setup: func(m *Client) {
			m.ExpectDBUpdates().WillReturn(NewDBUpdates().
				AddDelay(time.Millisecond).
				AddUpdate(&driver.DBUpdate{DBName: "foo"}).
				AddDelay(time.Second).
				AddUpdate(&driver.DBUpdate{DBName: "bar"}))
		},
		test: func(t *testing.T, c *kivik.Client) {
			ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
			defer cancel()
			updates, err := c.DBUpdates(ctx)
			testy.Error(t, "", err)
			names := []string{}
			for updates.Next() {
				names = append(names, updates.DBName())
			}
			expected := []string{"foo"}
			if d := diff.Interface(expected, names); d != nil {
				t.Error(d)
			}
			testy.Error(t, "context deadline exceeded", updates.Err())
		},
	})
	tests.Run(t, testMock)
}
