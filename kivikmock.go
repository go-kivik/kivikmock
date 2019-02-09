package kivikmock

import (
	"fmt"

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

	// ExpectDBExists queues an expectation for this client action to be
	// triggered.
	ExpectDBExists() *ExpectedDBExists

	// ExpectDestroyDB queues an expectation for this client action to be
	// triggered.
	ExpectDestroyDB() *ExpectedDestroyDB

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

func (c *kivikmock) ExpectDBExists() *ExpectedDBExists {
	e := &ExpectedDBExists{}
	c.expected = append(c.expected, e)
	return e
}

func (c *kivikmock) ExpectDestroyDB() *ExpectedDestroyDB {
	e := &ExpectedDestroyDB{}
	c.expected = append(c.expected, e)
	return e
}
