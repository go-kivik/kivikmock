package kivikmock

// MockDB serves to create expectations for database actions to
// mock and test real database behavior.
type MockDB struct {
	name   string
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

// ExpectAllDocs queues an expectation that DB.AllDocs() will be called.
func (db *MockDB) ExpectAllDocs() *ExpectedAllDocs {
	e := &ExpectedAllDocs{
		db: db,
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectBulkGet queues an expectation that DB.BulkGet() will be called.
func (db *MockDB) ExpectBulkGet() *ExpectedBulkGet {
	e := &ExpectedBulkGet{
		db: db,
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectFind queues an expectation that DB.Find() will be called.
func (db *MockDB) ExpectFind() *ExpectedFind {
	e := &ExpectedFind{
		db: db,
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectCreateIndex queues an expectation that DB.CreateIndex will be called.
func (db *MockDB) ExpectCreateIndex() *ExpectedCreateIndex {
	e := &ExpectedCreateIndex{
		db: db,
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectGetIndexes queues an expectation that DB.GetIndexes will be called.
func (db *MockDB) ExpectGetIndexes() *ExpectedGetIndexes {
	e := &ExpectedGetIndexes{
		db: db,
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectDeleteIndex queues an expectation that DB.DeleteIndex will be called.
func (db *MockDB) ExpectDeleteIndex() *ExpectedDeleteIndex {
	e := &ExpectedDeleteIndex{
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
