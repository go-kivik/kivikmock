package kivikmock

import (
	"fmt"
	"reflect"

	"github.com/go-kivik/kivik"
)

// MockClient allows configuring the mock kivik client.
type MockClient struct {
	ordered    bool
	dsn        string
	opened     int
	drv        *mockDriver
	expected   []expectation
	newdbcount int
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

// ExpectAuthenticate queues an expectation for an Authenticate() call.
func (c *MockClient) ExpectAuthenticate() *ExpectedAuthenticate {
	e := &ExpectedAuthenticate{}
	c.expected = append(c.expected, e)
	return e
}

// ExpectCreateDB queues an expectation for a CreateDB() call.
func (c *MockClient) ExpectCreateDB() *ExpectedCreateDB {
	e2 := &ExpectedDB{ret0: &MockDB{}}
	e := &ExpectedCreateDB{
		expectedDB: e2,
	}
	c.expected = append(c.expected, e, e2)
	return e
}

// NewDB creates a new mock DB object, which can be used along with ExpectDB()
// or ExpectCreateDB() calls to mock database actions.
func (c *MockClient) NewDB() *MockDB {
	c.newdbcount++
	return &MockDB{
		client: c,
		id:     c.newdbcount,
	}
}

// NewRows returns a new, empty set of rows, which can be returned by any of
// the row-returning expectations.
func NewRows() *Rows {
	return &Rows{}
}

// NewBulkResults returns a new, empty set of bulk results, which can be
// returned by the DB.BulkDocs() expectation.
func NewBulkResults() *BulkResults {
	return &BulkResults{}
}

// NewChanges returns a new, empty changes set, which can be returned by the
// DB.Changes() expectation.
func NewChanges() *Changes {
	return &Changes{}
}

// NewDBUpdates returns a new, empty update set, which can be returned by the
// DBUpdates() expectation.
func NewDBUpdates() *Updates {
	return &Updates{}
}
