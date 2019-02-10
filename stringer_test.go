package kivikmock

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/flimzy/testy"
	"github.com/go-kivik/kivik"
	"github.com/go-kivik/kivik/driver"
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
	tests.Add("delay", stringerTest{
		input: &ExpectedClose{commonExpectation{delay: time.Second}},
		expected: `call to Close() which:
	- should delay for: 1s`,
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
	tests.Add("delay", stringerTest{
		input: &ExpectedAllDBs{commonExpectation: commonExpectation{delay: time.Second}},
		expected: `call to AllDBs() which:
	- has any options
	- should delay for: 1s`,
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
	tests.Add("delay", stringerTest{
		input: &ExpectedAuthenticate{commonExpectation: commonExpectation{delay: time.Second}},
		expected: `call to Authenticate() which:
	- has any authenticator
	- should delay for: 1s`,
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
	tests.Add("delay", stringerTest{
		input: &ExpectedClusterSetup{commonExpectation: commonExpectation{delay: time.Second}},
		expected: `call to ClusterSetup() which:
	- has any action
	- should delay for: 1s`,
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
	tests.Add("delay", stringerTest{
		input: &ExpectedClusterStatus{commonExpectation: commonExpectation{delay: time.Second}},
		expected: `call to ClusterStatus() which:
	- has any options
	- should delay for: 1s`,
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
	tests.Add("delay", stringerTest{
		input: &ExpectedDBExists{commonExpectation: commonExpectation{delay: time.Second}},
		expected: `call to DBExists() which:
	- has any name
	- has any options
	- should delay for: 1s
	- should return: false`,
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
	tests.Add("delay", stringerTest{
		input: &ExpectedDestroyDB{commonExpectation: commonExpectation{delay: time.Second}},
		expected: `call to DestroyDB() which:
	- has any name
	- has any options
	- should delay for: 1s`,
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
	tests.Add("delay", stringerTest{
		input: &ExpectedDBsStats{commonExpectation: commonExpectation{delay: time.Second}},
		expected: `call to DBsStats() which:
	- has any names
	- should delay for: 1s`,
	})
	tests.Add("error", stringerTest{
		input: &ExpectedDBsStats{commonExpectation: commonExpectation{err: errors.New("foo err")}},
		expected: `call to DBsStats() which:
	- has any names
	- should return error: foo err`,
	})
	tests.Run(t, testStringer)
}

func TestPingString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input:    &ExpectedPing{},
		expected: `call to Ping()`,
	})
	tests.Add("error", stringerTest{
		input: &ExpectedPing{commonExpectation: commonExpectation{err: errors.New("foo err")}},
		expected: `call to Ping() which:
	- should return error: foo err`,
	})
	tests.Add("delay", stringerTest{
		input: &ExpectedPing{commonExpectation: commonExpectation{delay: time.Second}},
		expected: `call to Ping() which:
	- should delay for: 1s`,
	})
	tests.Run(t, testStringer)
}

func TestSessionString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input:    &ExpectedSession{},
		expected: `call to Session()`,
	})
	tests.Add("session", stringerTest{
		input: &ExpectedSession{session: &kivik.Session{Name: "bob"}},
		expected: `call to Session() which:
	- should return: &{bob []   [] []}`,
	})
	tests.Run(t, testStringer)
}

func TestVersionString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input:    &ExpectedVersion{},
		expected: `call to Version()`,
	})
	tests.Add("session", stringerTest{
		input: &ExpectedVersion{version: &kivik.Version{Version: "1.2"}},
		expected: `call to Version() which:
	- should return: &{1.2  [] []}`,
	})
	tests.Run(t, testStringer)
}

func TestCreateDBString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedCreateDB{},
		expected: `call to CreateDB() which:
	- has any name
	- has any options`,
	})
	tests.Add("name", stringerTest{
		input: &ExpectedCreateDB{name: "foo"},
		expected: `call to CreateDB() which:
	- has name: foo
	- has any options`,
	})
	tests.Add("db", stringerTest{
		input: &ExpectedCreateDB{db: &MockDB{count: 50}},
		expected: `call to CreateDB() which:
	- has any name
	- has any options
	- should return database with 50 expectations`,
	})
	tests.Run(t, testStringer)
}

func TestDBString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedDB{},
		expected: `call to DB() which:
	- has any name
	- has any options`,
	})
	tests.Add("name", stringerTest{
		input: &ExpectedDB{name: "foo"},
		expected: `call to DB() which:
	- has name: foo
	- has any options`,
	})
	tests.Add("db", stringerTest{
		input: &ExpectedDB{db: &MockDB{count: 50}},
		expected: `call to DB() which:
	- has any name
	- has any options
	- should return database with 50 expectations`,
	})
	tests.Run(t, testStringer)
}

func TestDBCloseString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("standard", stringerTest{
		input:    &ExpectedDBClose{},
		expected: "call to DB.Close()",
	})
	tests.Add("error", stringerTest{
		input: &ExpectedDBClose{commonExpectation{err: errors.New("foo error")}},
		expected: `call to DB.Close() which:
	- should return error: foo error`,
	})
	tests.Add("delay", stringerTest{
		input: &ExpectedDBClose{commonExpectation{delay: time.Second}},
		expected: `call to DB.Close() which:
	- should delay for: 1s`,
	})
	tests.Run(t, testStringer)
}

func TestAllDocsString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedAllDocs{},
		expected: `call to DB.AllDocs() which:
	- has any options
	- should return: 0 results`,
	})
	tests.Add("results", stringerTest{
		input: &ExpectedAllDocs{
			rows: &Rows{results: []*delayedRow{
				{Row: &driver.Row{}},
				{Row: &driver.Row{}},
				{delay: 15},
				{Row: &driver.Row{}},
				{Row: &driver.Row{}},
			},
			},
		},
		expected: `call to DB.AllDocs() which:
	- has any options
	- should return: 4 results`,
	})
	tests.Run(t, testStringer)
}

func TestBulkGetString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedBulkGet{},
		expected: `call to DB.BulkGet() which:
	- has any doc references
	- has any options
	- should return: 0 results`,
	})
	tests.Add("docs", stringerTest{
		input: &ExpectedBulkGet{docs: []driver.BulkGetReference{
			{ID: "foo"},
			{ID: "bar"},
		}},
		expected: `call to DB.BulkGet() which:
	- has doc references: [{foo  } {bar  }]
	- has any options
	- should return: 0 results`,
	})
	tests.Add("results", stringerTest{
		input: &ExpectedBulkGet{
			rows: &Rows{results: []*delayedRow{
				{Row: &driver.Row{}},
				{Row: &driver.Row{}},
				{delay: 15},
				{Row: &driver.Row{}},
				{Row: &driver.Row{}},
			},
			},
		},
		expected: `call to DB.BulkGet() which:
	- has any doc references
	- has any options
	- should return: 4 results`,
	})
	tests.Run(t, testStringer)
}

func TestFindString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedFind{},
		expected: `call to DB.Find() which:
	- has any query
	- should return: 0 results`,
	})
	tests.Add("query", stringerTest{
		input: &ExpectedFind{query: map[string]string{"foo": "bar"}},
		expected: `call to DB.Find() which:
	- has query: map[foo:bar]
	- should return: 0 results`,
	})
	tests.Add("results", stringerTest{
		input: &ExpectedFind{
			rows: &Rows{results: []*delayedRow{
				{Row: &driver.Row{}},
				{Row: &driver.Row{}},
				{delay: 15},
				{Row: &driver.Row{}},
				{Row: &driver.Row{}},
			},
			},
		},
		expected: `call to DB.Find() which:
	- has any query
	- should return: 4 results`,
	})
	tests.Run(t, testStringer)
}

func TestCreateIndexString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedCreateIndex{},
		expected: `call to DB.CreateIndex() which:
	- has any ddoc
	- has any name
	- has any index`,
	})
	tests.Add("ddoc", stringerTest{
		input: &ExpectedCreateIndex{ddoc: "foo"},
		expected: `call to DB.CreateIndex() which:
	- has ddoc: foo
	- has any name
	- has any index`,
	})
	tests.Add("name", stringerTest{
		input: &ExpectedCreateIndex{name: "foo"},
		expected: `call to DB.CreateIndex() which:
	- has any ddoc
	- has name: foo
	- has any index`,
	})
	tests.Add("index", stringerTest{
		input: &ExpectedCreateIndex{index: map[string]string{"foo": "bar"}},
		expected: `call to DB.CreateIndex() which:
	- has any ddoc
	- has any name
	- has index: map[foo:bar]`,
	})
	tests.Run(t, testStringer)
}

func TestExpectedGetIndexesString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input:    &ExpectedGetIndexes{},
		expected: `call to DB.GetIndexes()`,
	})
	tests.Add("error", stringerTest{
		input: &ExpectedGetIndexes{commonExpectation: commonExpectation{err: errors.New("foo err")}},
		expected: `call to DB.GetIndexes() which:
	- should return error: foo err`,
	})
	tests.Add("indexes", stringerTest{
		input: &ExpectedGetIndexes{indexes: []kivik.Index{{Name: "foo"}}},
		expected: `call to DB.GetIndexes() which:
	- should return indexes: [{ foo  <nil>}]`,
	})
	tests.Run(t, testStringer)
}
