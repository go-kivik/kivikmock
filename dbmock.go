package kivikmock

import "github.com/go-kivik/kivik/driver"

// MockDB serves to create expectations for database actions to
// mock and test real database behavior.
type MockDB interface {
	// ExpectClose queues an expectation Close() to be called on this database.
	ExpectClose() *ExpectedDBClose
	driver.DB

	// expectations returns the number of expectations registered in this db.
	expectations() int
}

type db struct {
	client *kivikmock
	count  int
	driver.DB
}

var _ MockDB = &db{}
var _ driver.DB = &db{}

func (db *db) expectations() int {
	return db.count
}

func (db *db) ExpectClose() *ExpectedDBClose {
	e := &ExpectedDBClose{}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}
