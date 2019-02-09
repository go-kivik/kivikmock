package kivikmock

import "github.com/go-kivik/kivik/driver"

// MockDB serves to create expectations for database actions to
// mock and test real database behavior.
type MockDB interface {
	// ExpectClose queues an expectation Close() to be called on this database.
	ExpectClose() *ExpectedDBClose
	driver.DB
}

type db struct {
	client *kivikmock
	count  int
	driver.DB
}

var _ MockDB = &db{}
var _ driver.DB = &db{}

func (db *db) ExpectClose() *ExpectedDBClose {
	e := &ExpectedDBClose{}
	db.client.expected = append(db.client.expected, e)
	return e
}
