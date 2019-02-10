package kivikmock

import (
	"context"
	"errors"

	"github.com/go-kivik/kivik/driver"
)

func (db *db) Close(ctx context.Context) error {
	expected := &ExpectedDBClose{}
	if err := db.client.nextExpectation(expected); err != nil {
		return err
	}

	return expected.wait(ctx)
}

func (db *db) AllDocs(ctx context.Context, options map[string]interface{}) (driver.Rows, error) {
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

var _ driver.BulkGetter = &db{}

func (db *db) BulkGet(ctx context.Context, docs []driver.BulkGetReference, options map[string]interface{}) (driver.Rows, error) {
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

var _ driver.Finder = &db{}

func (db *db) Find(ctx context.Context, query interface{}) (driver.Rows, error) {
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

func (db *db) CreateIndex(ctx context.Context, ddoc, name string, index interface{}) error {
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

func (db *db) GetIndexes(ctx context.Context) ([]driver.Index, error) {
	return nil, errors.New("unimplemented")
}

func (db *db) DeleteIndex(ctx context.Context, ddoc, name string) error {
	return errors.New("unimplemented")
}

func (db *db) Explain(ctx context.Context, query interface{}) (*driver.QueryPlan, error) {
	return nil, errors.New("unimplemented")
}
