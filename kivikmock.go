package kivikmock

import (
	"fmt"
	"reflect"

	"github.com/go-kivik/kivik"
	"github.com/go-kivik/kivik/driver"
)

// Mock interface serves to create expectations for database actions to
// mock and test real database behavior.
type Mock interface {
	// ExpectClose queues an expectation for this client action to be triggered.
	// *ExpectedClose allows mocking the response.
	ExpectClose() *ExpectedClose

	// ExpectationsWereMet returns an error if any expectations were not met.
	ExpectationsWereMet() error

	// ExpectAllDBs queues an expectation for this client action to be
	// triggered. *ExpectedAllDBs allows mocking the response.
	ExpectAllDBs() *ExpectedAllDBs

	// ExpectAuthenticate queues an expectation for this client action to be
	// triggered. *ExpectAuthenticate allows mocking the response.
	ExpectAuthenticate() *ExpectedAuthenticate

	// ExpectClusterSetup queues an expectation for this client action to be
	// triggered.
	ExpectClusterSetup() *ExpectedClusterSetup

	// ExpectClusterStatus queues an expectation for this client action to be
	// triggered.
	ExpectClusterStatus() *ExpectedClusterStatus

	// MatchExpectationsInOrder indicates whether to match expectations in the
	// order they were set.
	//
	// By default this is true, but if you use goroutines to parallelize
	// executation, that option may be handy.
	//
	// This option may be turned on any time during tests. As soon
	// as it is switched to false, expectations will be matched
	// in any order. Or otherwise if switched to true, any unmatched
	// expectations will be expected in order
	MatchExpectationsInOrder(bool)
}

type kivikmock struct {
	ordered  bool
	dsn      string
	opened   int
	drv      *mockDriver
	expected []expectation
	driver.Client
}

var _ driver.Client = &kivikmock{}

func (c *kivikmock) open() (*kivik.Client, Mock, error) {
	client, err := kivik.New("kivikmock", c.dsn)
	return client, c, err
}

func (c *kivikmock) ExpectationsWereMet() error {
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

func (c *kivikmock) MatchExpectationsInOrder(b bool) {
	c.ordered = b
}

func (c *kivikmock) ExpectClose() *ExpectedClose {
	e := &ExpectedClose{}
	c.expected = append(c.expected, e)
	return e
}

func (c *kivikmock) ExpectAllDBs() *ExpectedAllDBs {
	e := &ExpectedAllDBs{}
	c.expected = append(c.expected, e)
	return e
}

func (c *kivikmock) ExpectAuthenticate() *ExpectedAuthenticate {
	e := &ExpectedAuthenticate{}
	c.expected = append(c.expected, e)
	return e
}

func (c *kivikmock) ExpectClusterSetup() *ExpectedClusterSetup {
	e := &ExpectedClusterSetup{}
	c.expected = append(c.expected, e)
	return e
}

func (c *kivikmock) ExpectClusterStatus() *ExpectedClusterStatus {
	e := &ExpectedClusterStatus{}
	c.expected = append(c.expected, e)
	return e
}

// nextExpectation accepts the expected value `e`, checks that this is a valid
// expectation, and if so, populates e with the matching expectation. If the
// expectation is not expected, an error is returned.
func (c *kivikmock) nextExpectation(actual expectation) error {
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
