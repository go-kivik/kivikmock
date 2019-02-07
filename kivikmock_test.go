package kivikmock

import (
	"context"
	"fmt"
	"testing"
)

func TestExpectedCloseError(t *testing.T) {
	// Open new mock database
	client, mock, err := New()
	if err != nil {
		fmt.Println("error creating mock database")
		return
	}
	mock.ExpectClose().WillReturnError(fmt.Errorf("Close failed"))
	if err := client.Close(context.TODO()); err == nil {
		t.Error("an error was expected when calling close, but got none")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
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
