/* This file is auto-generated. Do not edit it! */

package kivikmock

import "github.com/go-kivik/kivik"

var _ = kivik.EndKeySuffix // To ensure a reference to kivik package

// ExpectCompact queues an expectation that DB.Compact will be called.
func (db *MockDB) ExpectCompact() *ExpectedCompact {
	e := &ExpectedCompact{
		db: db,
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectCompactView queues an expectation that DB.CompactView will be called.
func (db *MockDB) ExpectCompactView() *ExpectedCompactView {
	e := &ExpectedCompactView{
		db: db,
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectCopy queues an expectation that DB.Copy will be called.
func (db *MockDB) ExpectCopy() *ExpectedCopy {
	e := &ExpectedCopy{
		db: db,
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectCreateDoc queues an expectation that DB.CreateDoc will be called.
func (db *MockDB) ExpectCreateDoc() *ExpectedCreateDoc {
	e := &ExpectedCreateDoc{
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

// ExpectDelete queues an expectation that DB.Delete will be called.
func (db *MockDB) ExpectDelete() *ExpectedDelete {
	e := &ExpectedDelete{
		db: db,
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectDeleteAttachment queues an expectation that DB.DeleteAttachment will be called.
func (db *MockDB) ExpectDeleteAttachment() *ExpectedDeleteAttachment {
	e := &ExpectedDeleteAttachment{
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

// ExpectFlush queues an expectation that DB.Flush will be called.
func (db *MockDB) ExpectFlush() *ExpectedFlush {
	e := &ExpectedFlush{
		db: db,
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectGetMeta queues an expectation that DB.GetMeta will be called.
func (db *MockDB) ExpectGetMeta() *ExpectedGetMeta {
	e := &ExpectedGetMeta{
		db: db,
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectPut queues an expectation that DB.Put will be called.
func (db *MockDB) ExpectPut() *ExpectedPut {
	e := &ExpectedPut{
		db: db,
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectViewCleanup queues an expectation that DB.ViewCleanup will be called.
func (db *MockDB) ExpectViewCleanup() *ExpectedViewCleanup {
	e := &ExpectedViewCleanup{
		db: db,
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}
