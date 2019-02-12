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
		input: &ExpectedAllDBs{commonExpectation: commonExpectation{options: map[string]interface{}{"foo": 123}}},
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
		input: &ExpectedClusterSetup{arg0: map[string]string{"foo": "bar"}},
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
		input: &ExpectedClusterStatus{commonExpectation: commonExpectation{options: map[string]interface{}{"foo": 123}}},
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
		input: &ExpectedDBExists{arg0: "foo", ret0: true, commonExpectation: commonExpectation{options: map[string]interface{}{"foo": 123}}},
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
		input: &ExpectedDestroyDB{arg0: "foo"},
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
		input: &ExpectedCreateDB{arg0: "foo"},
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
		input: &ExpectedDB{arg0: "foo"},
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
		input:    &ExpectedDBClose{db: &MockDB{name: "foo"}},
		expected: "call to DB(foo#0).Close()",
	})
	tests.Add("error", stringerTest{
		input: &ExpectedDBClose{db: &MockDB{name: "foo"}, commonExpectation: commonExpectation{err: errors.New("foo error")}},
		expected: `call to DB(foo#0).Close() which:
	- should return error: foo error`,
	})
	tests.Add("delay", stringerTest{
		input: &ExpectedDBClose{db: &MockDB{name: "foo"}, commonExpectation: commonExpectation{delay: time.Second}},
		expected: `call to DB(foo#0).Close() which:
	- should delay for: 1s`,
	})
	tests.Run(t, testStringer)
}

func TestAllDocsString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedAllDocs{db: &MockDB{name: "foo"}},
		expected: `call to DB(foo#0).AllDocs() which:
	- has any options
	- should return: 0 results`,
	})
	tests.Add("results", stringerTest{
		input: &ExpectedAllDocs{
			db: &MockDB{name: "foo"},
			ret0: &Rows{results: []*delayedRow{
				{Row: &driver.Row{}},
				{Row: &driver.Row{}},
				{delay: 15},
				{Row: &driver.Row{}},
				{Row: &driver.Row{}},
			},
			},
		},
		expected: `call to DB(foo#0).AllDocs() which:
	- has any options
	- should return: 4 results`,
	})
	tests.Run(t, testStringer)
}

func TestBulkGetString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedBulkGet{db: &MockDB{name: "foo"}},
		expected: `call to DB(foo#0).BulkGet() which:
	- has any doc references
	- has any options
	- should return: 0 results`,
	})
	tests.Add("docs", stringerTest{
		input: &ExpectedBulkGet{db: &MockDB{name: "foo"}, arg0: []driver.BulkGetReference{
			{ID: "foo"},
			{ID: "bar"},
		}},
		expected: `call to DB(foo#0).BulkGet() which:
	- has doc references: [{foo  } {bar  }]
	- has any options
	- should return: 0 results`,
	})
	tests.Add("results", stringerTest{
		input: &ExpectedBulkGet{
			db: &MockDB{name: "foo"},
			ret0: &Rows{results: []*delayedRow{
				{Row: &driver.Row{}},
				{Row: &driver.Row{}},
				{delay: 15},
				{Row: &driver.Row{}},
				{Row: &driver.Row{}},
			},
			},
		},
		expected: `call to DB(foo#0).BulkGet() which:
	- has any doc references
	- has any options
	- should return: 4 results`,
	})
	tests.Run(t, testStringer)
}

func TestFindString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedFind{db: &MockDB{name: "foo"}},
		expected: `call to DB(foo#0).Find() which:
	- has any query
	- should return: 0 results`,
	})
	tests.Add("query", stringerTest{
		input: &ExpectedFind{db: &MockDB{name: "foo"}, arg0: map[string]string{"foo": "bar"}},
		expected: `call to DB(foo#0).Find() which:
	- has query: map[foo:bar]
	- should return: 0 results`,
	})
	tests.Add("results", stringerTest{
		input: &ExpectedFind{
			db: &MockDB{name: "foo"},
			ret0: &Rows{results: []*delayedRow{
				{Row: &driver.Row{}},
				{Row: &driver.Row{}},
				{delay: 15},
				{Row: &driver.Row{}},
				{Row: &driver.Row{}},
			},
			},
		},
		expected: `call to DB(foo#0).Find() which:
	- has any query
	- should return: 4 results`,
	})
	tests.Run(t, testStringer)
}

func TestCreateIndexString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedCreateIndex{db: &MockDB{name: "foo"}},
		expected: `call to DB(foo#0).CreateIndex() which:
	- has any ddoc
	- has any name
	- has any index`,
	})
	tests.Add("ddoc", stringerTest{
		input: &ExpectedCreateIndex{db: &MockDB{name: "foo"}, arg0: "foo"},
		expected: `call to DB(foo#0).CreateIndex() which:
	- has ddoc: foo
	- has any name
	- has any index`,
	})
	tests.Add("name", stringerTest{
		input: &ExpectedCreateIndex{db: &MockDB{name: "foo"}, arg1: "foo"},
		expected: `call to DB(foo#0).CreateIndex() which:
	- has any ddoc
	- has name: foo
	- has any index`,
	})
	tests.Add("index", stringerTest{
		input: &ExpectedCreateIndex{db: &MockDB{name: "foo"}, arg2: map[string]string{"foo": "bar"}},
		expected: `call to DB(foo#0).CreateIndex() which:
	- has any ddoc
	- has any name
	- has index: map[foo:bar]`,
	})
	tests.Run(t, testStringer)
}

func TestExpectedGetIndexesString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input:    &ExpectedGetIndexes{db: &MockDB{name: "foo"}},
		expected: `call to DB(foo#0).GetIndexes()`,
	})
	tests.Add("error", stringerTest{
		input: &ExpectedGetIndexes{db: &MockDB{name: "foo"}, commonExpectation: commonExpectation{err: errors.New("foo err")}},
		expected: `call to DB(foo#0).GetIndexes() which:
	- should return error: foo err`,
	})
	tests.Add("indexes", stringerTest{
		input: &ExpectedGetIndexes{db: &MockDB{name: "foo"}, ret0: []driver.Index{{Name: "foo"}}},
		expected: `call to DB(foo#0).GetIndexes() which:
	- should return indexes: [{ foo  <nil>}]`,
	})
	tests.Run(t, testStringer)
}

func TestDeleteIndexString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedDeleteIndex{db: &MockDB{name: "foo"}},
		expected: `call to DB(foo#0).DeleteIndex() which:
	- has any ddoc
	- has any name`,
	})
	tests.Add("ddoc", stringerTest{
		input: &ExpectedDeleteIndex{db: &MockDB{name: "foo"}, arg0: "foo"},
		expected: `call to DB(foo#0).DeleteIndex() which:
	- has ddoc: foo
	- has any name`,
	})
	tests.Add("name", stringerTest{
		input: &ExpectedDeleteIndex{db: &MockDB{name: "foo"}, arg1: "foo"},
		expected: `call to DB(foo#0).DeleteIndex() which:
	- has any ddoc
	- has name: foo`,
	})
	tests.Run(t, testStringer)
}

func TestExplainString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedExplain{db: &MockDB{name: "foo"}},
		expected: `call to DB(foo#0).Explain() which:
	- has any query`,
	})
	tests.Add("query", stringerTest{
		input: &ExpectedExplain{db: &MockDB{name: "foo"}, arg0: map[string]string{"foo": "bar"}},
		expected: `call to DB(foo#0).Explain() which:
	- has query: map[foo:bar]`,
	})
	tests.Add("plan", stringerTest{
		input: &ExpectedExplain{db: &MockDB{name: "foo"}, ret0: &driver.QueryPlan{DBName: "foo"}},
		expected: `call to DB(foo#0).Explain() which:
	- has any query
	- should return query plan: &{foo map[] map[] map[] 0 0 [] map[]}`,
	})
	tests.Add("error", stringerTest{
		input: &ExpectedExplain{db: &MockDB{name: "foo"}, commonExpectation: commonExpectation{err: errors.New("foo err")}},
		expected: `call to DB(foo#0).Explain() which:
	- has any query
	- should return error: foo err`,
	})
	tests.Run(t, testStringer)
}

func TestCreateDocString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedCreateDoc{db: &MockDB{name: "foo"}},
		expected: `call to DB(foo#0).CreateDoc() which:
	- has any doc
	- has any options`,
	})
	tests.Add("doc", stringerTest{
		input: &ExpectedCreateDoc{db: &MockDB{name: "foo"}, arg0: map[string]string{"foo": "bar"}},
		expected: `call to DB(foo#0).CreateDoc() which:
	- has doc: map[foo:bar]
	- has any options`,
	})
	tests.Add("options", stringerTest{
		input: &ExpectedCreateDoc{db: &MockDB{name: "foo"}, commonExpectation: commonExpectation{options: map[string]interface{}{"foo": "bar"}}},
		expected: `call to DB(foo#0).CreateDoc() which:
	- has any doc
	- has options: map[foo:bar]`,
	})
	tests.Add("error", stringerTest{
		input: &ExpectedCreateDoc{db: &MockDB{name: "foo"}, commonExpectation: commonExpectation{err: errors.New("foo err")}},
		expected: `call to DB(foo#0).CreateDoc() which:
	- has any doc
	- has any options
	- should return error: foo err`,
	})
	tests.Add("delay", stringerTest{
		input: &ExpectedCreateDoc{db: &MockDB{name: "foo"}, commonExpectation: commonExpectation{delay: time.Second}},
		expected: `call to DB(foo#0).CreateDoc() which:
	- has any doc
	- has any options
	- should delay for: 1s`,
	})
	tests.Add("docID", stringerTest{
		input: &ExpectedCreateDoc{db: &MockDB{name: "foo"}, ret0: "foo"},
		expected: `call to DB(foo#0).CreateDoc() which:
	- has any doc
	- has any options
	- should return docID: foo`,
	})
	tests.Add("rev", stringerTest{
		input: &ExpectedCreateDoc{db: &MockDB{name: "foo"}, ret1: "1-foo"},
		expected: `call to DB(foo#0).CreateDoc() which:
	- has any doc
	- has any options
	- should return rev: 1-foo`,
	})
	tests.Run(t, testStringer)
}

func TestCompactString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input:    &ExpectedCompact{db: &MockDB{name: "foo"}},
		expected: `call to DB(foo#0).Compact()`,
	})
	tests.Add("error", stringerTest{
		input: &ExpectedCompact{db: &MockDB{name: "foo"}, commonExpectation: commonExpectation{err: errors.New("foo err")}},
		expected: `call to DB(foo#0).Compact() which:
	- should return error: foo err`,
	})
	tests.Add("delay", stringerTest{
		input: &ExpectedCompact{db: &MockDB{name: "foo"}, commonExpectation: commonExpectation{delay: time.Second}},
		expected: `call to DB(foo#0).Compact() which:
	- should delay for: 1s`,
	})

	tests.Run(t, testStringer)
}

func TestViewCleanupString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input:    &ExpectedViewCleanup{db: &MockDB{name: "foo"}},
		expected: `call to DB(foo#0).ViewCleanup()`,
	})
	tests.Add("error", stringerTest{
		input: &ExpectedViewCleanup{db: &MockDB{name: "foo"}, commonExpectation: commonExpectation{err: errors.New("foo err")}},
		expected: `call to DB(foo#0).ViewCleanup() which:
	- should return error: foo err`,
	})
	tests.Add("delay", stringerTest{
		input: &ExpectedViewCleanup{db: &MockDB{name: "foo"}, commonExpectation: commonExpectation{delay: time.Second}},
		expected: `call to DB(foo#0).ViewCleanup() which:
	- should delay for: 1s`,
	})

	tests.Run(t, testStringer)
}

func TestPutString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedPut{db: &MockDB{name: "foo"}},
		expected: `call to DB(foo#0).Put() which:
	- has any docID
	- has any doc
	- has any options`,
	})
	tests.Add("docID", stringerTest{
		input: &ExpectedPut{db: &MockDB{name: "foo"}, arg0: "foo"},
		expected: `call to DB(foo#0).Put() which:
	- has docID: foo
	- has any doc
	- has any options`,
	})
	tests.Add("doc", stringerTest{
		input: &ExpectedPut{db: &MockDB{name: "foo"}, arg1: map[string]string{"foo": "bar"}},
		expected: `call to DB(foo#0).Put() which:
	- has any docID
	- has doc: map[foo:bar]
	- has any options`,
	})
	tests.Add("error", stringerTest{
		input: &ExpectedPut{db: &MockDB{name: "foo"}, commonExpectation: commonExpectation{err: errors.New("foo err")}},
		expected: `call to DB(foo#0).Put() which:
	- has any docID
	- has any doc
	- has any options
	- should return error: foo err`,
	})
	tests.Add("delay", stringerTest{
		input: &ExpectedPut{db: &MockDB{name: "foo"}, commonExpectation: commonExpectation{delay: time.Second}},
		expected: `call to DB(foo#0).Put() which:
	- has any docID
	- has any doc
	- has any options
	- should delay for: 1s`,
	})
	tests.Run(t, testStringer)
}

func TestGetMetaString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedGetMeta{db: &MockDB{name: "foo"}},
		expected: `call to DB(foo#0).GetMeta() which:
	- has any docID
	- has any options`,
	})
	tests.Add("docID", stringerTest{
		input: &ExpectedGetMeta{db: &MockDB{name: "foo"}, arg0: "foo"},
		expected: `call to DB(foo#0).GetMeta() which:
	- has docID: foo
	- has any options`,
	})
	tests.Add("size", stringerTest{
		input: &ExpectedGetMeta{db: &MockDB{name: "foo"}, ret0: 123},
		expected: `call to DB(foo#0).GetMeta() which:
	- has any docID
	- has any options
	- should return size: 123`,
	})
	tests.Add("rev", stringerTest{
		input: &ExpectedGetMeta{db: &MockDB{name: "foo"}, ret1: "1-xxx"},
		expected: `call to DB(foo#0).GetMeta() which:
	- has any docID
	- has any options
	- should return rev: 1-xxx`,
	})
	tests.Add("error", stringerTest{
		input: &ExpectedGetMeta{db: &MockDB{name: "foo"}, commonExpectation: commonExpectation{err: errors.New("foo err")}},
		expected: `call to DB(foo#0).GetMeta() which:
	- has any docID
	- has any options
	- should return error: foo err`,
	})
	tests.Add("delay", stringerTest{
		input: &ExpectedGetMeta{db: &MockDB{name: "foo"}, commonExpectation: commonExpectation{delay: time.Second}},
		expected: `call to DB(foo#0).GetMeta() which:
	- has any docID
	- has any options
	- should delay for: 1s`,
	})
	tests.Run(t, testStringer)
}

func TestCompactViewString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedCompactView{db: &MockDB{name: "foo"}},
		expected: `call to DB(foo#0).CompactView() which:
	- has any ddocID`,
	})
	tests.Add("ddocID", stringerTest{
		input: &ExpectedCompactView{db: &MockDB{name: "foo"}, arg0: "foo"},
		expected: `call to DB(foo#0).CompactView() which:
	- has ddocID: foo`,
	})
	tests.Add("error", stringerTest{
		input: &ExpectedCompactView{db: &MockDB{name: "foo"}, commonExpectation: commonExpectation{err: errors.New("foo err")}},
		expected: `call to DB(foo#0).CompactView() which:
	- has any ddocID
	- should return error: foo err`,
	})
	tests.Add("delay", stringerTest{
		input: &ExpectedCompactView{db: &MockDB{name: "foo"}, commonExpectation: commonExpectation{delay: time.Second}},
		expected: `call to DB(foo#0).CompactView() which:
	- has any ddocID
	- should delay for: 1s`,
	})
	tests.Run(t, testStringer)
}

func TestFlushString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input:    &ExpectedFlush{db: &MockDB{name: "foo"}},
		expected: `call to DB(foo#0).Flush()`,
	})
	tests.Run(t, testStringer)
}

func TestDeleteAttachmentString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedDeleteAttachment{db: &MockDB{name: "foo"}},
		expected: `call to DB(foo#0).DeleteAttachment() which:
	- has any docID
	- has any rev
	- has any filename
	- has any options`,
	})
	tests.Add("docID", stringerTest{
		input: &ExpectedDeleteAttachment{db: &MockDB{name: "foo"}, arg0: "foo"},
		expected: `call to DB(foo#0).DeleteAttachment() which:
	- has docID: foo
	- has any rev
	- has any filename
	- has any options`,
	})
	tests.Add("rev", stringerTest{
		input: &ExpectedDeleteAttachment{db: &MockDB{name: "foo"}, arg1: "1-foo"},
		expected: `call to DB(foo#0).DeleteAttachment() which:
	- has any docID
	- has rev: 1-foo
	- has any filename
	- has any options`,
	})
	tests.Add("filename", stringerTest{
		input: &ExpectedDeleteAttachment{db: &MockDB{name: "foo"}, arg2: "foo.txt"},
		expected: `call to DB(foo#0).DeleteAttachment() which:
	- has any docID
	- has any rev
	- has filename: foo.txt
	- has any options`,
	})
	tests.Add("return", stringerTest{
		input: &ExpectedDeleteAttachment{db: &MockDB{name: "foo"}, ret0: "2-bar"},
		expected: `call to DB(foo#0).DeleteAttachment() which:
	- has any docID
	- has any rev
	- has any filename
	- has any options
	- should return rev: 2-bar`,
	})
	tests.Run(t, testStringer)
}

func TestDeleteString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedDelete{db: &MockDB{name: "foo"}},
		expected: `call to DB(foo#0).Delete() which:
	- has any docID
	- has any rev
	- has any options`,
	})
	tests.Add("docID", stringerTest{
		input: &ExpectedDelete{db: &MockDB{name: "foo"}, arg0: "foo"},
		expected: `call to DB(foo#0).Delete() which:
	- has docID: foo
	- has any rev
	- has any options`,
	})
	tests.Add("rev", stringerTest{
		input: &ExpectedDelete{db: &MockDB{name: "foo"}, arg1: "1-foo"},
		expected: `call to DB(foo#0).Delete() which:
	- has any docID
	- has rev: 1-foo
	- has any options`,
	})
	tests.Add("options", stringerTest{
		input: &ExpectedDelete{db: &MockDB{name: "foo"}, commonExpectation: commonExpectation{options: map[string]interface{}{"foo": "bar"}}},
		expected: `call to DB(foo#0).Delete() which:
	- has any docID
	- has any rev
	- has options: map[foo:bar]`,
	})
	tests.Add("return", stringerTest{
		input: &ExpectedDelete{db: &MockDB{name: "foo"}, ret0: "2-bar"},
		expected: `call to DB(foo#0).Delete() which:
	- has any docID
	- has any rev
	- has any options
	- should return rev: 2-bar`,
	})
	tests.Run(t, testStringer)
}

func TestCopyString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedCopy{db: &MockDB{name: "foo"}},
		expected: `call to DB(foo#0).Copy() which:
	- has any targetID
	- has any sourceID
	- has any options`,
	})
	tests.Add("targetID", stringerTest{
		input: &ExpectedCopy{db: &MockDB{name: "foo"}, arg0: "foo"},
		expected: `call to DB(foo#0).Copy() which:
	- has targetID: foo
	- has any sourceID
	- has any options`,
	})
	tests.Add("sourceID", stringerTest{
		input: &ExpectedCopy{db: &MockDB{name: "foo"}, arg1: "foo"},
		expected: `call to DB(foo#0).Copy() which:
	- has any targetID
	- has sourceID: foo
	- has any options`,
	})
	tests.Add("return value", stringerTest{
		input: &ExpectedCopy{db: &MockDB{name: "foo"}, ret0: "1-foo"},
		expected: `call to DB(foo#0).Copy() which:
	- has any targetID
	- has any sourceID
	- has any options
	- should return rev: 1-foo`,
	})
	tests.Run(t, testStringer)
}

func TestGetString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedGet{db: &MockDB{name: "foo"}},
		expected: `call to DB(foo#0).Get() which:
	- has any docID
	- has any options`,
	})
	tests.Add("docID", stringerTest{
		input: &ExpectedGet{db: &MockDB{name: "foo"}, arg0: "foo"},
		expected: `call to DB(foo#0).Get() which:
	- has docID: foo
	- has any options`,
	})
	tests.Add("return value", stringerTest{
		input: &ExpectedGet{db: &MockDB{name: "foo"}, ret0: &driver.Document{Rev: "1-foo"}},
		expected: `call to DB(foo#0).Get() which:
	- has any docID
	- has any options
	- should return document with rev: 1-foo`,
	})
	tests.Run(t, testStringer)
}

func TestGetAttachmentMetaString(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", stringerTest{
		input: &ExpectedGetAttachmentMeta{db: &MockDB{name: "foo"}},
		expected: `call to DB(foo#0).GetAttachmentMeta() which:
	- has any docID
	- has any filename
	- has any options`,
	})
	tests.Add("docID", stringerTest{
		input: &ExpectedGetAttachmentMeta{db: &MockDB{name: "foo"}, arg0: "foo"},
		expected: `call to DB(foo#0).GetAttachmentMeta() which:
	- has docID: foo
	- has any filename
	- has any options`,
	})
	tests.Add("filename", stringerTest{
		input: &ExpectedGetAttachmentMeta{db: &MockDB{name: "foo"}, arg1: "foo.txt"},
		expected: `call to DB(foo#0).GetAttachmentMeta() which:
	- has any docID
	- has filename: foo.txt
	- has any options`,
	})
	tests.Add("return value", stringerTest{
		input: &ExpectedGetAttachmentMeta{db: &MockDB{name: "foo"}, ret0: &driver.Attachment{Filename: "foo.txt"}},
		expected: `call to DB(foo#0).GetAttachmentMeta() which:
	- has any docID
	- has any filename
	- has any options
	- should return attachment: foo.txt`,
	})
	tests.Run(t, testStringer)
}
