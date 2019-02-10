package kivikmock

import (
	"context"
	"errors"

	"github.com/go-kivik/kivik/driver"
)

type driverDB struct {
	*MockDB
}

var _ driver.DB = &driverDB{}
var _ driver.BulkGetter = &driverDB{}
var _ driver.Finder = &driverDB{}

func (db *driverDB) Close(ctx context.Context) error {
	expected := &ExpectedDBClose{}
	if err := db.client.nextExpectation(expected); err != nil {
		return err
	}

	return expected.wait(ctx)
}

func (db *driverDB) AllDocs(ctx context.Context, options map[string]interface{}) (driver.Rows, error) {
	expected := &ExpectedAllDocs{
		options: options,
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return nil, err
	}
	return &driverRows{
		Context: ctx,
		Rows:    expected.rows,
	}, expected.wait(ctx)
}

func (db *driverDB) BulkGet(ctx context.Context, docs []driver.BulkGetReference, options map[string]interface{}) (driver.Rows, error) {
	expected := &ExpectedBulkGet{
		options: options,
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return nil, err
	}
	return &driverRows{
		Context: ctx,
		Rows:    expected.rows,
	}, expected.wait(ctx)
}

func (db *driverDB) Find(ctx context.Context, query interface{}) (driver.Rows, error) {
	expected := &ExpectedFind{
		query: query,
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return nil, err
	}
	return &driverRows{
		Context: ctx,
		Rows:    expected.rows,
	}, expected.wait(ctx)
}

func (db *driverDB) CreateIndex(ctx context.Context, ddoc, name string, index interface{}) error {
	expected := &ExpectedCreateIndex{
		ddoc:  ddoc,
		name:  name,
		index: index,
	}
	if err := db.client.nextExpectation(expected); err != nil {
		return err
	}
	return expected.wait(ctx)
}

func (db *driverDB) GetIndexes(ctx context.Context) ([]driver.Index, error) {
	return nil, errors.New("unimplemented")
}

func (db *driverDB) DeleteIndex(ctx context.Context, ddoc, name string) error {
	return errors.New("unimplemented")
}

func (db *driverDB) Explain(ctx context.Context, query interface{}) (*driver.QueryPlan, error) {
	return nil, errors.New("unimplemented")
}

func (db *driverDB) Changes(ctx context.Context, options map[string]interface{}) (driver.Changes, error) {
	return nil, errors.New("unimplemented")
}

func (db *driverDB) Compact(ctx context.Context) error {
	return errors.New("unimplemented")
}

func (db *driverDB) CompactView(ctx context.Context, view string) error {
	return errors.New("unimplemented")
}

func (db *driverDB) CreateDoc(ctx context.Context, doc interface{}, options map[string]interface{}) (string, string, error) {
	return "", "", errors.New("unimplemented")
}

func (db *driverDB) Delete(ctx context.Context, _, _ string, options map[string]interface{}) (string, error) {
	return "", errors.New("unimplemented")
}

func (db *driverDB) DeleteAttachment(ctx context.Context, _, _, _ string, options map[string]interface{}) (string, error) {
	return "", errors.New("unimplemented")
}

func (db *driverDB) Get(ctx context.Context, _ string, options map[string]interface{}) (*driver.Document, error) {
	return nil, errors.New("unimplemented")
}

func (db *driverDB) GetAttachment(ctx context.Context, _, _, _ string, options map[string]interface{}) (*driver.Attachment, error) {
	return nil, errors.New("unimplemented")
}

func (db *driverDB) Put(ctx context.Context, _ string, doc interface{}, options map[string]interface{}) (string, error) {
	return "", errors.New("unimplemented")
}

func (db *driverDB) PutAttachment(ctx context.Context, _, _ string, att *driver.Attachment, options map[string]interface{}) (string, error) {
	return "", errors.New("unimplemented")
}

func (db *driverDB) Query(ctx context.Context, _, _ string, options map[string]interface{}) (driver.Rows, error) {
	return nil, errors.New("unimplemented")
}

func (db *driverDB) Security(ctx context.Context) (*driver.Security, error) {
	return nil, errors.New("unimplemented")
}

func (db *driverDB) SetSecurity(ctx context.Context, sec *driver.Security) error {
	return errors.New("unimplemented")
}

func (db *driverDB) Stats(ctx context.Context) (*driver.DBStats, error) {
	return nil, errors.New("unimplemented")
}

func (db *driverDB) ViewCleanup(ctx context.Context) error {
	return errors.New("unimplemented")
}
