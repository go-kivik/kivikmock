package kivikmock

import (
	"context"

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
