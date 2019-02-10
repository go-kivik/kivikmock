package kivikmock

import "github.com/go-kivik/kivik/driver"

// MockDB serves to create expectations for database actions to
// mock and test real database behavior.
type MockDB struct {
	client *kivikmock
	count  int
	driver.DB
}

var _ driver.DB = &MockDB{}

func (db *MockDB) expectations() int {
	return db.count
}

// ExpectClose queues an expectation for DB.Close() to be called.
func (db *MockDB) ExpectClose() *ExpectedDBClose {
	e := &ExpectedDBClose{}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectAllDocs queues an expectation that DB.AllDocs() will be called.
func (db *MockDB) ExpectAllDocs() *ExpectedAllDocs {
	e := &ExpectedAllDocs{}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectBulkGet queues an expectation that DB.BulkGet() will be called.
func (db *MockDB) ExpectBulkGet() *ExpectedBulkGet {
	e := &ExpectedBulkGet{}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectFind queues an expectation that DB.Find() will be called.
func (db *MockDB) ExpectFind() *ExpectedFind {
	e := &ExpectedFind{}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectCreateIndex queues an expectation that DB.CreateIndex will be called.
func (db *MockDB) ExpectCreateIndex() *ExpectedCreateIndex {
	e := &ExpectedCreateIndex{}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// NewRows returns a new, empty set of rows, which can be returned by any of
// the row-returning expectations.
func (db *MockDB) NewRows() *Rows {
	return &Rows{
		results: make([]*delayedRow, 0),
	}
}
