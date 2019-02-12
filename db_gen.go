/* This file is auto-generated. Do not edit it! */

package kivikmock

import (
	"context"

	"github.com/go-kivik/kivik/driver"
)

var _ = &driver.Attachment{}

func (db *driverDB) Compact(ctx context.Context) error {
	expected := &ExpectedCompact{
		db: db.MockDB,
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return err
	}
	return expected.wait(ctx)
}

func (db *driverDB) CompactView(ctx context.Context, arg0 string) error {
	expected := &ExpectedCompactView{
		db:   db.MockDB,
		arg0: arg0,
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return err
	}
	return expected.wait(ctx)
}

func (db *driverDB) Copy(ctx context.Context, arg0 string, arg1 string, options map[string]interface{}) (string, error) {
	expected := &ExpectedCopy{
		db:                db.MockDB,
		arg0:              arg0,
		arg1:              arg1,
		commonExpectation: commonExpectation{options: options},
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return "", err
	}
	return expected.ret0, expected.wait(ctx)
}

func (db *driverDB) CreateDoc(ctx context.Context, arg0 interface{}, options map[string]interface{}) (string, string, error) {
	expected := &ExpectedCreateDoc{
		db:                db.MockDB,
		arg0:              arg0,
		commonExpectation: commonExpectation{options: options},
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return "", "", err
	}
	return expected.ret0, expected.ret1, expected.wait(ctx)
}

func (db *driverDB) CreateIndex(ctx context.Context, arg0 string, arg1 string, arg2 interface{}) error {
	expected := &ExpectedCreateIndex{
		db:   db.MockDB,
		arg0: arg0,
		arg1: arg1,
		arg2: arg2,
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return err
	}
	return expected.wait(ctx)
}

func (db *driverDB) Delete(ctx context.Context, arg0 string, arg1 string, options map[string]interface{}) (string, error) {
	expected := &ExpectedDelete{
		db:                db.MockDB,
		arg0:              arg0,
		arg1:              arg1,
		commonExpectation: commonExpectation{options: options},
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return "", err
	}
	return expected.ret0, expected.wait(ctx)
}

func (db *driverDB) DeleteAttachment(ctx context.Context, arg0 string, arg1 string, arg2 string, options map[string]interface{}) (string, error) {
	expected := &ExpectedDeleteAttachment{
		db:                db.MockDB,
		arg0:              arg0,
		arg1:              arg1,
		arg2:              arg2,
		commonExpectation: commonExpectation{options: options},
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return "", err
	}
	return expected.ret0, expected.wait(ctx)
}

func (db *driverDB) DeleteIndex(ctx context.Context, arg0 string, arg1 string) error {
	expected := &ExpectedDeleteIndex{
		db:   db.MockDB,
		arg0: arg0,
		arg1: arg1,
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return err
	}
	return expected.wait(ctx)
}

func (db *driverDB) Flush(ctx context.Context) error {
	expected := &ExpectedFlush{
		db: db.MockDB,
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return err
	}
	return expected.wait(ctx)
}

func (db *driverDB) GetMeta(ctx context.Context, arg0 string, options map[string]interface{}) (int64, string, error) {
	expected := &ExpectedGetMeta{
		db:                db.MockDB,
		arg0:              arg0,
		commonExpectation: commonExpectation{options: options},
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return 0, "", err
	}
	return expected.ret0, expected.ret1, expected.wait(ctx)
}

func (db *driverDB) Put(ctx context.Context, arg0 string, arg1 interface{}, options map[string]interface{}) (string, error) {
	expected := &ExpectedPut{
		db:                db.MockDB,
		arg0:              arg0,
		arg1:              arg1,
		commonExpectation: commonExpectation{options: options},
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return "", err
	}
	return expected.ret0, expected.wait(ctx)
}

func (db *driverDB) ViewCleanup(ctx context.Context) error {
	expected := &ExpectedViewCleanup{
		db: db.MockDB,
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return err
	}
	return expected.wait(ctx)
}

func (db *driverDB) AllDocs(ctx context.Context, options map[string]interface{}) (driver.Rows, error) {
	expected := &ExpectedAllDocs{
		db:                db.MockDB,
		commonExpectation: commonExpectation{options: options},
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return nil, err
	}
	return &driverRows{Context: ctx, Rows: expected.ret0}, expected.wait(ctx)
}

func (db *driverDB) BulkDocs(ctx context.Context, arg0 []interface{}, options map[string]interface{}) (driver.BulkResults, error) {
	expected := &ExpectedBulkDocs{
		db:                db.MockDB,
		arg0:              arg0,
		commonExpectation: commonExpectation{options: options},
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return nil, err
	}
	return expected.ret0, expected.wait(ctx)
}

func (db *driverDB) BulkGet(ctx context.Context, arg0 []driver.BulkGetReference, options map[string]interface{}) (driver.Rows, error) {
	expected := &ExpectedBulkGet{
		db:                db.MockDB,
		arg0:              arg0,
		commonExpectation: commonExpectation{options: options},
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return nil, err
	}
	return &driverRows{Context: ctx, Rows: expected.ret0}, expected.wait(ctx)
}

func (db *driverDB) Changes(ctx context.Context, options map[string]interface{}) (driver.Changes, error) {
	expected := &ExpectedChanges{
		db:                db.MockDB,
		commonExpectation: commonExpectation{options: options},
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return nil, err
	}
	return expected.ret0, expected.wait(ctx)
}

func (db *driverDB) DesignDocs(ctx context.Context, options map[string]interface{}) (driver.Rows, error) {
	expected := &ExpectedDesignDocs{
		db:                db.MockDB,
		commonExpectation: commonExpectation{options: options},
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return nil, err
	}
	return &driverRows{Context: ctx, Rows: expected.ret0}, expected.wait(ctx)
}

func (db *driverDB) Explain(ctx context.Context, arg0 interface{}) (*driver.QueryPlan, error) {
	expected := &ExpectedExplain{
		db:   db.MockDB,
		arg0: arg0,
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return nil, err
	}
	return expected.ret0, expected.wait(ctx)
}

func (db *driverDB) Find(ctx context.Context, arg0 interface{}) (driver.Rows, error) {
	expected := &ExpectedFind{
		db:   db.MockDB,
		arg0: arg0,
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return nil, err
	}
	return &driverRows{Context: ctx, Rows: expected.ret0}, expected.wait(ctx)
}

func (db *driverDB) Get(ctx context.Context, arg0 string, options map[string]interface{}) (*driver.Document, error) {
	expected := &ExpectedGet{
		db:                db.MockDB,
		arg0:              arg0,
		commonExpectation: commonExpectation{options: options},
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return nil, err
	}
	return expected.ret0, expected.wait(ctx)
}

func (db *driverDB) GetAttachment(ctx context.Context, arg0 string, arg1 string, arg2 string, options map[string]interface{}) (*driver.Attachment, error) {
	expected := &ExpectedGetAttachment{
		db:                db.MockDB,
		arg0:              arg0,
		arg1:              arg1,
		arg2:              arg2,
		commonExpectation: commonExpectation{options: options},
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return nil, err
	}
	return expected.ret0, expected.wait(ctx)
}

func (db *driverDB) GetAttachmentMeta(ctx context.Context, arg0 string, arg1 string, arg2 string, options map[string]interface{}) (*driver.Attachment, error) {
	expected := &ExpectedGetAttachmentMeta{
		db:                db.MockDB,
		arg0:              arg0,
		arg1:              arg1,
		arg2:              arg2,
		commonExpectation: commonExpectation{options: options},
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return nil, err
	}
	return expected.ret0, expected.wait(ctx)
}

func (db *driverDB) GetIndexes(ctx context.Context) ([]driver.Index, error) {
	expected := &ExpectedGetIndexes{
		db: db.MockDB,
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return nil, err
	}
	return expected.ret0, expected.wait(ctx)
}

func (db *driverDB) LocalDocs(ctx context.Context, options map[string]interface{}) (driver.Rows, error) {
	expected := &ExpectedLocalDocs{
		db:                db.MockDB,
		commonExpectation: commonExpectation{options: options},
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return nil, err
	}
	return &driverRows{Context: ctx, Rows: expected.ret0}, expected.wait(ctx)
}

func (db *driverDB) Purge(ctx context.Context, arg0 map[string][]string) (*driver.PurgeResult, error) {
	expected := &ExpectedPurge{
		db:   db.MockDB,
		arg0: arg0,
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return nil, err
	}
	return expected.ret0, expected.wait(ctx)
}

func (db *driverDB) PutAttachment(ctx context.Context, arg0 string, arg1 string, arg2 *driver.Attachment, options map[string]interface{}) (string, error) {
	expected := &ExpectedPutAttachment{
		db:                db.MockDB,
		arg0:              arg0,
		arg1:              arg1,
		arg2:              arg2,
		commonExpectation: commonExpectation{options: options},
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return "", err
	}
	return expected.ret0, expected.wait(ctx)
}

func (db *driverDB) Query(ctx context.Context, arg0 string, arg1 string, options map[string]interface{}) (driver.Rows, error) {
	expected := &ExpectedQuery{
		db:                db.MockDB,
		arg0:              arg0,
		arg1:              arg1,
		commonExpectation: commonExpectation{options: options},
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return nil, err
	}
	return &driverRows{Context: ctx, Rows: expected.ret0}, expected.wait(ctx)
}

func (db *driverDB) Security(ctx context.Context) (*driver.Security, error) {
	expected := &ExpectedSecurity{
		db: db.MockDB,
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return nil, err
	}
	return expected.ret0, expected.wait(ctx)
}

func (db *driverDB) SetSecurity(ctx context.Context, arg0 *driver.Security) error {
	expected := &ExpectedSetSecurity{
		db:   db.MockDB,
		arg0: arg0,
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return err
	}
	return expected.wait(ctx)
}

func (db *driverDB) Stats(ctx context.Context) (*driver.DBStats, error) {
	expected := &ExpectedStats{
		db: db.MockDB,
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return nil, err
	}
	return expected.ret0, expected.wait(ctx)
}
