package kivikmock

import "context"

func (db *db) Close(ctx context.Context) error {
	expected := &ExpectedDBClose{}
	if err := db.client.nextExpectation(expected); err != nil {
		return err
	}

	return expected.wait(ctx)
}
