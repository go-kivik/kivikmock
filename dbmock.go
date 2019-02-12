package kivikmock

// MockDB serves to create expectations for database actions to
// mock and test real database behavior.
type MockDB struct {
	name   string
	id     int
	client *MockClient
	count  int
}

func (db *MockDB) expectations() int {
	return db.count
}

// ExpectClose queues an expectation for DB.Close() to be called.
func (db *MockDB) ExpectClose() *ExpectedDBClose {
	e := &ExpectedDBClose{
		db: db,
	}
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
