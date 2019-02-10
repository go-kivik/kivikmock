package kivikmock

import "github.com/go-kivik/kivik/driver"

// MockDB serves to create expectations for database actions to
// mock and test real database behavior.
type MockDB interface {
	// ExpectClose queues an expectation for Close() to be called on this database.
	ExpectClose() *ExpectedDBClose

	// ExpectAllDocs queues an expectation for AllDocs() to be called.
	ExpectAllDocs() *ExpectedAllDocs
	driver.DB

	// expectations returns the number of expectations registered in this db.
	expectations() int

	NewRows() *Rows
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

func (db *db) ExpectAllDocs() *ExpectedAllDocs {
	e := &ExpectedAllDocs{}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

func (db *db) NewRows() *Rows {
	return &Rows{
		results: make([]*delayedRow, 0),
	}
}
