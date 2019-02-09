package kivikmock

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/flimzy/testy"
	"github.com/go-kivik/kivik"
)

func TestCloseDB(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("error", mockTest{
		setup: func(m Mock) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectClose().WillReturnError(errors.New("foo err"))
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").Close(context.TODO())
			testy.Error(t, "foo err", err)
		},
		err: "",
	})
	tests.Add("unexpected", mockTest{
		setup: func(m Mock) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").Close(context.TODO())
			testy.Error(t, "call to DB.Close() was not expected, all expectations already fulfilled", err)
		},
	})
	tests.Add("delay", mockTest{
		setup: func(m Mock) {
			db := m.NewDB()
			m.ExpectDB().WillReturn(db)
			db.ExpectClose().WillDelay(time.Second)
		},
		test: func(t *testing.T, c *kivik.Client) {
			err := c.DB(context.TODO(), "foo").Close(newCanceledContext())
			testy.Error(t, "context canceled", err)
		},
	})

	tests.Run(t, testMock)
}
