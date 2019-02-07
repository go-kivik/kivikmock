package kivikmock

import (
	"context"
	"fmt"
	"testing"

	"github.com/flimzy/testy"
	"github.com/go-kivik/kivik"
)

func TestExpectedCloseError(t *testing.T) {
	// Open new mock database
	client, mock, err := New()
	if err != nil {
		fmt.Println("error creating mock database")
		return
	}
	mock.ExpectClose().WillReturnError(fmt.Errorf("Close failed"))
	err = client.Close(context.TODO())
	testy.Error(t, "Close failed", err)
	err = mock.ExpectationsWereMet()
	testy.Error(t, "", err)
}

// func TestExpectedCloseOrder(t *testing.T) {
// 	// Open new mock database
// 	client, mock, err := New()
// 	if err != nil {
// 		fmt.Println("error creating mock database")
// 		return
// 	}
// 	defer client.Close(context.TODO())
// 	mock.ExpectClose().WillReturnError(fmt.Errorf("Close failed"))
// 	client.Begin()
// 	if err := mock.ExpectationsWereMet(); err == nil {
// 		t.Error("expected error on ExpectationsWereMet")
// 	}
// }

func TestExpectedAllDBsError(t *testing.T) {
	// Open new mock database
	client, mock, err := New()
	if err != nil {
		fmt.Println("error creating mock database")
		return
	}
	mock.ExpectAllDBs().WillReturnError(fmt.Errorf("AllDBs failed"))
	expectedErr := "AllDBs failed"
	_, err = client.AllDBs(context.TODO())
	testy.Error(t, expectedErr, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestExpectedAllDBsOrder(t *testing.T) {
	// Open new mock database
	client, mock, err := New()
	if err != nil {
		fmt.Println("error creating mock database")
		return
	}
	defer client.Close(context.TODO()) // nolint: errcheck
	mock.ExpectAllDBs().WillReturn([]string{"a", "b"})
	err = mock.ExpectationsWereMet()
	testy.ErrorRE(t, `should return: \[a b\]`, err)
}

func TestExpectedAllDBsUnexpected(t *testing.T) {
	// Open new mock database
	client, _, err := New()
	if err != nil {
		fmt.Println("error creating mock database")
		return
	}
	defer client.Close(context.TODO()) // nolint: errcheck
	_, err = client.AllDBs(context.TODO(), kivik.Options{"Foo": 123})
	expectedErr := `all expectations were already fulfilled, call to AllDBs with options map[Foo:123] was not expected`
	testy.Error(t, expectedErr, err)
}

func TestExpectedAllDBsUnexpected_out_of_order(t *testing.T) {
	// Open new mock database
	client, mock, err := New()
	if err != nil {
		fmt.Println("error creating mock database")
		return
	}
	defer client.Close(context.TODO()) // nolint: errcheck
	mock.ExpectClose()
	_, err = client.AllDBs(context.TODO(), kivik.Options{"Foo": 123})
	expectedErr := `call to AllDBs with options map[Foo:123] was not expected. Next expectation is: ExpectedClose => expecting client Close`
	testy.Error(t, expectedErr, err)
}

func TestExpectedAllDBsUnexpectedUnorderedError(t *testing.T) {
	// Open new mock database
	client, mock, err := New()
	if err != nil {
		fmt.Println("error creating mock database")
		return
	}
	defer client.Close(context.TODO()) // nolint: errcheck
	mock.MatchExpectationsInOrder(false)
	mock.ExpectAllDBs().WithOptions(kivik.Options{"foo": 321})
	_, err = client.AllDBs(context.TODO(), kivik.Options{"Foo": 123})
	expectedErr := `call to AllDBs with options map[Foo:123] was not expected`
	testy.Error(t, expectedErr, err)
	expectedErr = ""
	err = mock.ExpectationsWereMet()
	testy.Error(t, expectedErr, err)
}

func TestExpectedAllDBsUnexpectedUnorderedSuccess(t *testing.T) {
	// Open new mock database
	client, mock, err := New()
	if err != nil {
		fmt.Println("error creating mock database")
		return
	}
	defer client.Close(context.TODO()) // nolint: errcheck
	mock.MatchExpectationsInOrder(false)
	mock.ExpectAllDBs().WithOptions(kivik.Options{"foo": 321})
	mock.ExpectAllDBs().WithOptions(kivik.Options{"Foo": 123})
	_, err = client.AllDBs(context.TODO(), kivik.Options{"Foo": 123})
	testy.Error(t, "", err)
	err = mock.ExpectationsWereMet()
	testy.ErrorRE(t, `map\[foo:321\]`, err)
}
