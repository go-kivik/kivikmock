package kivikmock

// MockDB serves to create expectations for database actions to
// mock and test real database behavior.
type MockDB struct {
	name   string
	id     int
	client *Client
	count  int
}

func (db *MockDB) expectations() int {
	return db.count
}

// ExpectClose queues an expectation for DB.Close() to be called.
func (db *MockDB) ExpectClose() *ExpectedDBClose {
	e := &ExpectedDBClose{
		commonExpectation: commonExpectation{db: db},
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}
