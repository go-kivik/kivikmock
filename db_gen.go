/* This file is auto-generated. Do not edit it! */

package kivikmock

import "context"

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
