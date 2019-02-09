package kivikmock

import (
	"errors"
	"fmt"
	"testing"

	"github.com/flimzy/testy"
)

type stringerTest struct {
	input    fmt.Stringer
	expected string
}

func testStringer(t *testing.T, test stringerTest) {
	result := test.input.String()
	if test.expected != result {
		t.Errorf("Unexpected String() output.\nWant: %s\n Got: %s\n", test.expected, result)
	}
}

func TestCloseString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("standard", stringerTest{
		input:    &ExpectedClose{},
		expected: "call to Close()",
	})
	tests.Add("error", stringerTest{
		input: &ExpectedClose{commonExpectation{err: errors.New("foo error")}},
		expected: `call to Close() which:
	- should return error: foo error`,
	})
	tests.Run(t, testStringer)
}

func TestAllDBsString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("standard", stringerTest{
		input: &ExpectedAllDBs{},
		expected: `call to AllDBs() which:
	- has any options`,
	})
	tests.Add("options", stringerTest{
		input: &ExpectedAllDBs{options: map[string]interface{}{"foo": 123}},
		expected: `call to AllDBs() which:
	- has options: map[foo:123]`,
	})
	tests.Add("error", stringerTest{
		input: &ExpectedAllDBs{commonExpectation: commonExpectation{err: errors.New("foo err")}},
		expected: `call to AllDBs() which:
	- has any options
	- should return error: foo err`,
	})
	tests.Run(t, testStringer)
}

func TestAuthenticateString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedAuthenticate{},
		expected: `call to Authenticate() which:
	- has any authenticator`,
	})
	tests.Add("authenticator", stringerTest{
		input: &ExpectedAuthenticate{authType: "foo"},
		expected: `call to Authenticate() which:
	- has an authenticator of type: foo`,
	})
	tests.Add("error", stringerTest{
		input: &ExpectedAuthenticate{commonExpectation: commonExpectation{err: errors.New("foo err")}},
		expected: `call to Authenticate() which:
	- has any authenticator
	- should return error: foo err`,
	})
	tests.Run(t, testStringer)
}

func TestClusterSetupString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedClusterSetup{},
		expected: `call to ClusterSetup() which:
	- has any action`,
	})
	tests.Add("action", stringerTest{
		input: &ExpectedClusterSetup{action: map[string]string{"foo": "bar"}},
		expected: `call to ClusterSetup() which:
	- has the action: map[foo:bar]`,
	})
	tests.Add("error", stringerTest{
		input: &ExpectedClusterSetup{commonExpectation: commonExpectation{err: errors.New("foo error")}},
		expected: `call to ClusterSetup() which:
	- has any action
	- should return error: foo error`,
	})
	tests.Run(t, testStringer)
}

func TestClusterStatusString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedClusterStatus{},
		expected: `call to ClusterStatus() which:
	- has any options`,
	})
	tests.Add("options", stringerTest{
		input: &ExpectedClusterStatus{options: map[string]interface{}{"foo": 123}},
		expected: `call to ClusterStatus() which:
	- has options: map[foo:123]`,
	})
	tests.Add("error", stringerTest{
		input: &ExpectedClusterStatus{commonExpectation: commonExpectation{err: errors.New("foo error")}},
		expected: `call to ClusterStatus() which:
	- has any options
	- should return error: foo error`,
	})
	tests.Run(t, testStringer)
}

func TestDBExistsString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedDBExists{},
		expected: `call to DBExists() which:
	- has any name
	- has any options
	- should return: false`,
	})
	tests.Add("full", stringerTest{
		input: &ExpectedDBExists{name: "foo", exists: true, options: map[string]interface{}{"foo": 123}},
		expected: `call to DBExists() which:
	- has name: foo
	- has options: map[foo:123]
	- should return: true`,
	})
	tests.Add("error", stringerTest{
		input: &ExpectedDBExists{commonExpectation: commonExpectation{err: errors.New("foo err")}},
		expected: `call to DBExists() which:
	- has any name
	- has any options
	- should return error: foo err`,
	})
	tests.Run(t, testStringer)
}

func TestDestroyDBString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedDestroyDB{},
		expected: `call to DestroyDB() which:
	- has any name
	- has any options`,
	})
	tests.Add("name", stringerTest{
		input: &ExpectedDestroyDB{name: "foo"},
		expected: `call to DestroyDB() which:
	- has name: foo
	- has any options`,
	})
	tests.Run(t, testStringer)
}

func TestDBsStatsString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedDBsStats{},
		expected: `call to DBsStats() which:
	- has any names`,
	})
	tests.Add("names", stringerTest{
		input: &ExpectedDBsStats{names: []string{"a", "b"}},
		expected: `call to DBsStats() which:
	- has names: [a b]`,
	})
	tests.Add("error", stringerTest{
		input: &ExpectedDBsStats{commonExpectation: commonExpectation{err: errors.New("foo err")}},
		expected: `call to DBsStats() which:
	- has any names
	- should return error: foo err`,
	})
	tests.Run(t, testStringer)
}
