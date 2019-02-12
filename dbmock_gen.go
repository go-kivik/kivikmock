/* This file is auto-generated. Do not edit it! */

package kivikmock

import (
	"github.com/go-kivik/kivik"
	"github.com/go-kivik/kivik/driver"
)

var _ = kivik.EndKeySuffix // To ensure a reference to kivik package
var _ = &driver.Attachment{}

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

// ExpectAllDocs queues an expectation that DB.AllDocs will be called.
func (db *MockDB) ExpectAllDocs() *ExpectedAllDocs {
	e := &ExpectedAllDocs{
		db:   db,
		ret0: &Rows{},
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectBulkDocs queues an expectation that DB.BulkDocs will be called.
func (db *MockDB) ExpectBulkDocs() *ExpectedBulkDocs {
	e := &ExpectedBulkDocs{
		db: db,
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectBulkGet queues an expectation that DB.BulkGet will be called.
func (db *MockDB) ExpectBulkGet() *ExpectedBulkGet {
	e := &ExpectedBulkGet{
		db:   db,
		ret0: &Rows{},
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectChanges queues an expectation that DB.Changes will be called.
func (db *MockDB) ExpectChanges() *ExpectedChanges {
	e := &ExpectedChanges{
		db: db,
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectDesignDocs queues an expectation that DB.DesignDocs will be called.
func (db *MockDB) ExpectDesignDocs() *ExpectedDesignDocs {
	e := &ExpectedDesignDocs{
		db:   db,
		ret0: &Rows{},
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectExplain queues an expectation that DB.Explain will be called.
func (db *MockDB) ExpectExplain() *ExpectedExplain {
	e := &ExpectedExplain{
		db:   db,
		ret0: &driver.QueryPlan{},
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectFind queues an expectation that DB.Find will be called.
func (db *MockDB) ExpectFind() *ExpectedFind {
	e := &ExpectedFind{
		db:   db,
		ret0: &Rows{},
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectGet queues an expectation that DB.Get will be called.
func (db *MockDB) ExpectGet() *ExpectedGet {
	e := &ExpectedGet{
		db: db,
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectGetAttachment queues an expectation that DB.GetAttachment will be called.
func (db *MockDB) ExpectGetAttachment() *ExpectedGetAttachment {
	e := &ExpectedGetAttachment{
		db: db,
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectGetAttachmentMeta queues an expectation that DB.GetAttachmentMeta will be called.
func (db *MockDB) ExpectGetAttachmentMeta() *ExpectedGetAttachmentMeta {
	e := &ExpectedGetAttachmentMeta{
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

// ExpectLocalDocs queues an expectation that DB.LocalDocs will be called.
func (db *MockDB) ExpectLocalDocs() *ExpectedLocalDocs {
	e := &ExpectedLocalDocs{
		db:   db,
		ret0: &Rows{},
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectPurge queues an expectation that DB.Purge will be called.
func (db *MockDB) ExpectPurge() *ExpectedPurge {
	e := &ExpectedPurge{
		db:   db,
		ret0: &driver.PurgeResult{},
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectPutAttachment queues an expectation that DB.PutAttachment will be called.
func (db *MockDB) ExpectPutAttachment() *ExpectedPutAttachment {
	e := &ExpectedPutAttachment{
		db: db,
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectQuery queues an expectation that DB.Query will be called.
func (db *MockDB) ExpectQuery() *ExpectedQuery {
	e := &ExpectedQuery{
		db:   db,
		ret0: &Rows{},
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectSecurity queues an expectation that DB.Security will be called.
func (db *MockDB) ExpectSecurity() *ExpectedSecurity {
	e := &ExpectedSecurity{
		db: db,
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectSetSecurity queues an expectation that DB.SetSecurity will be called.
func (db *MockDB) ExpectSetSecurity() *ExpectedSetSecurity {
	e := &ExpectedSetSecurity{
		db: db,
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}

// ExpectStats queues an expectation that DB.Stats will be called.
func (db *MockDB) ExpectStats() *ExpectedStats {
	e := &ExpectedStats{
		db: db,
	}
	db.count++
	db.client.expected = append(db.client.expected, e)
	return e
}
