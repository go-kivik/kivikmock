package kivikmock

import (
	"fmt"
	"reflect"

	"github.com/go-kivik/kivik"
)

// MockClient allows configuring the mock kivik client.
type MockClient struct {
	ordered  bool
	dsn      string
	opened   int
	drv      *mockDriver
	expected []expectation
}

// nextExpectation accepts the expected value `e`, checks that this is a valid
// expectation, and if so, populates e with the matching expectation. If the
// expectation is not expected, an error is returned.
func (c *MockClient) nextExpectation(actual expectation) error {
	c.drv.Lock()
	defer c.drv.Unlock()

	var expected expectation
	var fulfilled int
	for _, next := range c.expected {
		next.Lock()
		if next.fulfilled() {
			next.Unlock()
			fulfilled++
			continue
		}

		if c.ordered {
			if reflect.TypeOf(actual).Elem().Name() == reflect.TypeOf(next).Elem().Name() {
				if meets(actual, next) {
					expected = next
					break
				}
				next.Unlock()
				return fmt.Errorf("Expectation not met:\nExpected: %s\n  Actual: %s",
					next, actual)
			}
			next.Unlock()
			return fmt.Errorf("call to %s was not expected. Next expectation is: %s", actual.method(false), next.method(false))
		}
		if meets(actual, next) {
			expected = next
			break
		}

		next.Unlock()
	}

	if expected == nil {
		if fulfilled == len(c.expected) {
			return fmt.Errorf("call to %s was not expected, all expectations already fulfilled", actual.method(false))
		}
		return fmt.Errorf("call to %s was not expected", actual.method(!c.ordered))
	}

	defer expected.Unlock()
	expected.fulfill()

	reflect.ValueOf(actual).Elem().Set(reflect.ValueOf(expected).Elem())
	return nil
}

func meets(a, e expectation) bool {
	if reflect.TypeOf(a).Elem().Name() != reflect.TypeOf(e).Elem().Name() {
		return false
	}
	return a.met(e)
}

func (c *MockClient) open() (*kivik.Client, *MockClient, error) {
	client, err := kivik.New("kivikmock", c.dsn)
	return client, c, err
}

// ExpectationsWereMet returns an error if any outstanding expectatios were
// not met.
func (c *MockClient) ExpectationsWereMet() error {
	for _, e := range c.expected {
		e.Lock()
		fulfilled := e.fulfilled()
		e.Unlock()

		if !fulfilled {
			return fmt.Errorf("there is a remaining unmet expectation: %s", e)
		}
	}
	return nil
}

// MatchExpectationsInOrder sets whether expectations should occur in the
// precise order in which they were defined.
func (c *MockClient) MatchExpectationsInOrder(b bool) {
	c.ordered = b
}

// ExpectClose queues an expectation for a Close() call.
func (c *MockClient) ExpectClose() *ExpectedClose {
	e := &ExpectedClose{}
	c.expected = append(c.expected, e)
	return e
}

// ExpectAllDBs queues an expectation for an AllDBs() call.
func (c *MockClient) ExpectAllDBs() *ExpectedAllDBs {
	e := &ExpectedAllDBs{}
	c.expected = append(c.expected, e)
	return e
}

// ExpectAuthenticate queues an expectation for an Authenticate() call.
func (c *MockClient) ExpectAuthenticate() *ExpectedAuthenticate {
	e := &ExpectedAuthenticate{}
	c.expected = append(c.expected, e)
	return e
}

// ExpectClusterSetup queues an expectation for a ClusterSetup() call.
func (c *MockClient) ExpectClusterSetup() *ExpectedClusterSetup {
	e := &ExpectedClusterSetup{}
	c.expected = append(c.expected, e)
	return e
}

// ExpectClusterStatus queues an expectation for a ClusterStatus() call.
func (c *MockClient) ExpectClusterStatus() *ExpectedClusterStatus {
	e := &ExpectedClusterStatus{}
	c.expected = append(c.expected, e)
	return e
}

// ExpectDBExists queues an expectation for a DBExists() call.
func (c *MockClient) ExpectDBExists() *ExpectedDBExists {
	e := &ExpectedDBExists{}
	c.expected = append(c.expected, e)
	return e
}

// ExpectDestroyDB queues an expectation for a DestroyDB() call.
func (c *MockClient) ExpectDestroyDB() *ExpectedDestroyDB {
	e := &ExpectedDestroyDB{}
	c.expected = append(c.expected, e)
	return e
}

// ExpectDBsStats queues an expectation for a DBsStats() call.
func (c *MockClient) ExpectDBsStats() *ExpectedDBsStats {
	e := &ExpectedDBsStats{}
	c.expected = append(c.expected, e)
	return e
}

// ExpectPing queues an expectation for a Ping() call.
func (c *MockClient) ExpectPing() *ExpectedPing {
	e := &ExpectedPing{}
	c.expected = append(c.expected, e)
	return e
}

// ExpectSession queues an expectation for a Session() call.
func (c *MockClient) ExpectSession() *ExpectedSession {
	e := &ExpectedSession{}
	c.expected = append(c.expected, e)
	return e
}

// ExpectVersion queues an expectation for a Version() call.
func (c *MockClient) ExpectVersion() *ExpectedVersion {
	e := &ExpectedVersion{}
	c.expected = append(c.expected, e)
	return e
}

// ExpectCreateDB queues an expectation for a CreateDB() call.
func (c *MockClient) ExpectCreateDB() *ExpectedCreateDB {
	e2 := &ExpectedDB{}
	e := &ExpectedCreateDB{
		expectedDB: e2,
	}
	c.expected = append(c.expected, e, e2)
	return e
}

// ExpectDB queues an expectation for a DB() call.
func (c *MockClient) ExpectDB() *ExpectedDB {
	e := &ExpectedDB{}
	c.expected = append(c.expected, e)
	return e
}

// NewDB creates a new mock DB object, which can be used along with ExpectDB()
// or ExpectCreateDB() calls to mock database actions.
func (c *MockClient) NewDB() *MockDB {
	return &MockDB{
		client: c,
	}
}
