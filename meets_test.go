package kivikmock

import (
	"testing"

	"github.com/flimzy/testy"
)

func TestDBMeetsExpectation(t *testing.T) {
	type tst struct {
		exp      *MockDB
		act      *MockDB
		expected bool
	}
	tests := testy.NewTable()
	tests.Add("different name", tst{
		exp:      &MockDB{name: "foo"},
		act:      &MockDB{name: "bar"},
		expected: false,
	})
	tests.Add("different id", tst{
		exp:      &MockDB{name: "foo", id: 123},
		act:      &MockDB{name: "foo", id: 321},
		expected: false,
	})
	tests.Add("no db", tst{
		expected: true,
	})
	tests.Add("match", tst{
		exp:      &MockDB{name: "foo", id: 123},
		act:      &MockDB{name: "foo", id: 123},
		expected: true,
	})
	tests.Run(t, func(t *testing.T, test tst) {
		result := dbMeetsExpectation(test.exp, test.act)
		if result != test.expected {
			t.Errorf("Unexpected result: %T", result)
		}
	})
}
