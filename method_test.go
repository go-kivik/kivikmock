package kivikmock

import (
	"testing"

	"github.com/flimzy/testy"
	"github.com/go-kivik/kivik/driver"
)

type methodTest struct {
	input    expectation
	standard string
	verbose  string
}

func testMethod(t *testing.T, test methodTest) {
	result := test.input.method(false)
	if result != test.standard {
		t.Errorf("Unexpected method(false) output.\nWant: %s\n Got: %s\n", test.standard, result)
	}
	result = test.input.method(true)
	if result != test.verbose {
		t.Errorf("Unexpected method(true) output.\nWant: %s\n Got: %s\n", test.verbose, result)
	}
}

func TestCloseMethod(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", methodTest{
		input:    &ExpectedClose{},
		standard: "Close()",
		verbose:  "Close(ctx)",
	})
	tests.Run(t, testMethod)
}

func TestAllDBsMethod(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty ", methodTest{
		input:    &ExpectedAllDBs{},
		standard: "AllDBs()",
		verbose:  "AllDBs(ctx, nil)",
	})
	tests.Add("options", methodTest{
		input:    &ExpectedAllDBs{options: map[string]interface{}{"foo": 123}},
		standard: "AllDBs()",
		verbose:  `AllDBs(ctx, map[foo:123])`,
	})
	tests.Run(t, testMethod)
}

func TestAuthenticateMethod(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", methodTest{
		input:    &ExpectedAuthenticate{},
		standard: "Authenticate()",
		verbose:  "Authenticate(ctx, <T>)",
	})
	tests.Add("authenticator", methodTest{
		input:    &ExpectedAuthenticate{authType: "foo"},
		standard: "Authenticate()",
		verbose:  "Authenticate(ctx, <foo>)",
	})
	tests.Run(t, testMethod)
}

func TestClusterSetupMethod(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", methodTest{
		input:    &ExpectedClusterSetup{},
		standard: "ClusterSetup()",
		verbose:  "ClusterSetup(ctx, <T>)",
	})
	tests.Add("action", methodTest{
		input:    &ExpectedClusterSetup{action: map[string]string{"foo": "bar"}},
		standard: "ClusterSetup()",
		verbose:  "ClusterSetup(ctx, map[foo:bar])",
	})
	tests.Run(t, testMethod)
}

func TestClusterStatusMethod(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", methodTest{
		input:    &ExpectedClusterStatus{},
		standard: "ClusterStatus()",
		verbose:  "ClusterStatus(ctx, ?)",
	})
	tests.Add("options", methodTest{
		input:    &ExpectedClusterStatus{options: map[string]interface{}{"foo": 123}},
		standard: "ClusterStatus()",
		verbose:  "ClusterStatus(ctx, map[foo:123])",
	})
	tests.Add("no options", methodTest{
		input:    &ExpectedClusterStatus{options: map[string]interface{}{}},
		standard: "ClusterStatus()",
		verbose:  "ClusterStatus(ctx, map[])",
	})
	tests.Run(t, testMethod)
}

func TestDBExistsMethod(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", methodTest{
		input:    &ExpectedDBExists{},
		standard: "DBExists()",
		verbose:  "DBExists(ctx, ?, ?)",
	})
	tests.Add("name", methodTest{
		input:    &ExpectedDBExists{arg0: "foo"},
		standard: "DBExists()",
		verbose:  `DBExists(ctx, "foo", ?)`,
	})
	tests.Add("options", methodTest{
		input:    &ExpectedDBExists{options: map[string]interface{}{"foo": 321}},
		standard: "DBExists()",
		verbose:  `DBExists(ctx, ?, map[foo:321])`,
	})
	tests.Add("full", methodTest{
		input:    &ExpectedDBExists{arg0: "foo", options: map[string]interface{}{"foo": 321}},
		standard: "DBExists()",
		verbose:  `DBExists(ctx, "foo", map[foo:321])`,
	})
	tests.Run(t, testMethod)
}

func TestDestroyDBMethod(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", methodTest{
		input:    &ExpectedDestroyDB{},
		standard: "DestroyDB()",
		verbose:  "DestroyDB(ctx, ?, ?)",
	})
	tests.Add("name", methodTest{
		input:    &ExpectedDestroyDB{arg0: "foo"},
		standard: "DestroyDB()",
		verbose:  `DestroyDB(ctx, "foo", ?)`,
	})
	tests.Add("options", methodTest{
		input:    &ExpectedDestroyDB{options: map[string]interface{}{"foo": 12}},
		standard: "DestroyDB()",
		verbose:  `DestroyDB(ctx, ?, map[foo:12])`,
	})
	tests.Run(t, testMethod)
}

func TestDBsStatsMethod(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", methodTest{
		input:    &ExpectedDBsStats{},
		standard: "DBsStats()",
		verbose:  "DBsStats(ctx, ?)",
	})
	tests.Add("names", methodTest{
		input:    &ExpectedDBsStats{names: []string{"a", "b"}},
		standard: "DBsStats()",
		verbose:  `DBsStats(ctx, [a b])`,
	})
	tests.Run(t, testMethod)
}

func TestPingMethod(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", methodTest{
		input:    &ExpectedPing{},
		standard: "Ping()",
		verbose:  "Ping(ctx)",
	})
	tests.Run(t, testMethod)
}

func TestSessionMethod(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", methodTest{
		input:    &ExpectedSession{},
		standard: "Session()",
		verbose:  "Session(ctx)",
	})
	tests.Run(t, testMethod)
}

func TestVersionMethod(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", methodTest{
		input:    &ExpectedVersion{},
		standard: "Version()",
		verbose:  "Version(ctx)",
	})
	tests.Run(t, testMethod)
}

func TestCreateDBMethod(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", methodTest{
		input:    &ExpectedCreateDB{},
		standard: "CreateDB()",
		verbose:  "CreateDB(ctx, ?)",
	})
	tests.Add("options", methodTest{
		input:    &ExpectedCreateDB{options: map[string]interface{}{"foo": 123}},
		standard: "CreateDB()",
		verbose:  "CreateDB(ctx, ?, map[foo:123])",
	})
	tests.Add("name", methodTest{
		input:    &ExpectedCreateDB{arg0: "foo", options: map[string]interface{}{"foo": 123}},
		standard: "CreateDB()",
		verbose:  `CreateDB(ctx, "foo", map[foo:123])`,
	})
	tests.Run(t, testMethod)
}

func TestDBMethod(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", methodTest{
		input:    &ExpectedDB{},
		standard: "DB()",
		verbose:  "DB(ctx, ?)",
	})
	tests.Add("options", methodTest{
		input:    &ExpectedDB{options: map[string]interface{}{"foo": 123}},
		standard: "DB()",
		verbose:  "DB(ctx, ?, map[foo:123])",
	})
	tests.Add("name", methodTest{
		input:    &ExpectedDB{arg0: "foo", options: map[string]interface{}{"foo": 123}},
		standard: "DB()",
		verbose:  `DB(ctx, "foo", map[foo:123])`,
	})
	tests.Run(t, testMethod)
}

func TestDBCloseMethod(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", methodTest{
		input:    &ExpectedDBClose{},
		standard: "DB.Close()",
		verbose:  "DB.Close(ctx)",
	})
	tests.Run(t, testMethod)
}

func TestAllDocsMethod(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", methodTest{
		input:    &ExpectedAllDocs{db: &MockDB{name: "foo"}},
		standard: "DB.AllDocs()",
		verbose:  "DB(foo).AllDocs(ctx)",
	})
	tests.Add("options", methodTest{
		input:    &ExpectedAllDocs{db: &MockDB{name: "foo"}, options: map[string]interface{}{"foo": 123}},
		standard: "DB.AllDocs()",
		verbose:  "DB(foo).AllDocs(ctx, map[foo:123])",
	})
	tests.Run(t, testMethod)
}

func TestBulkGetMethod(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", methodTest{
		input:    &ExpectedBulkGet{db: &MockDB{name: "foo"}},
		standard: "DB.BulkGet()",
		verbose:  "DB(foo).BulkGet(ctx, ?)",
	})
	tests.Add("docs", methodTest{
		input:    &ExpectedBulkGet{db: &MockDB{name: "foo"}, docs: []driver.BulkGetReference{{ID: "foo"}}},
		standard: "DB.BulkGet()",
		verbose:  "DB(foo).BulkGet(ctx, [{foo  }])",
	})
	tests.Add("options", methodTest{
		input:    &ExpectedBulkGet{db: &MockDB{name: "foo"}, options: map[string]interface{}{"foo": 123}},
		standard: "DB.BulkGet()",
		verbose:  "DB(foo).BulkGet(ctx, ?, map[foo:123])",
	})
	tests.Run(t, testMethod)
}

func TestFindMethod(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", methodTest{
		input:    &ExpectedFind{db: &MockDB{name: "foo"}},
		standard: "DB.Find()",
		verbose:  "DB(foo).Find(ctx, ?)",
	})
	tests.Add("query", methodTest{
		input:    &ExpectedFind{db: &MockDB{name: "foo"}, query: map[string]string{"foo": "bar"}},
		standard: "DB.Find()",
		verbose:  "DB(foo).Find(ctx, map[foo:bar])",
	})
	tests.Run(t, testMethod)
}

func TestCreateIndexMethod(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", methodTest{
		input:    &ExpectedCreateIndex{db: &MockDB{name: "foo"}},
		standard: "DB.CreateIndex()",
		verbose:  "DB(foo).CreateIndex(ctx, ?, ?, ?)",
	})
	tests.Add("name", methodTest{
		input:    &ExpectedCreateIndex{db: &MockDB{name: "foo"}, name: "foo"},
		standard: "DB.CreateIndex()",
		verbose:  `DB(foo).CreateIndex(ctx, ?, "foo", ?)`,
	})
	tests.Add("ddoc", methodTest{
		input:    &ExpectedCreateIndex{db: &MockDB{name: "foo"}, ddoc: "foo"},
		standard: "DB.CreateIndex()",
		verbose:  `DB(foo).CreateIndex(ctx, "foo", ?, ?)`,
	})
	tests.Add("index", methodTest{
		input:    &ExpectedCreateIndex{db: &MockDB{name: "foo"}, index: map[string]string{"foo": "bar"}},
		standard: "DB.CreateIndex()",
		verbose:  `DB(foo).CreateIndex(ctx, ?, ?, map[foo:bar])`,
	})
	tests.Run(t, testMethod)
}

func TestGetIndexesMethod(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", methodTest{
		input:    &ExpectedGetIndexes{db: &MockDB{name: "foo"}},
		standard: "DB.GetIndexes()",
		verbose:  "DB(foo).GetIndexes(ctx)",
	})
	tests.Run(t, testMethod)
}

func TestDeleteIndexMethod(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", methodTest{
		input:    &ExpectedDeleteIndex{db: &MockDB{name: "foo"}},
		standard: "DB.DeleteIndex()",
		verbose:  "DB(foo).DeleteIndex(ctx, ?, ?)",
	})
	tests.Add("ddoc", methodTest{
		input:    &ExpectedDeleteIndex{db: &MockDB{name: "foo"}, ddoc: "foo"},
		standard: "DB.DeleteIndex()",
		verbose:  `DB(foo).DeleteIndex(ctx, "foo", ?)`,
	})
	tests.Add("name", methodTest{
		input:    &ExpectedDeleteIndex{db: &MockDB{name: "foo"}, name: "foo"},
		standard: "DB.DeleteIndex()",
		verbose:  `DB(foo).DeleteIndex(ctx, ?, "foo")`,
	})
	tests.Run(t, testMethod)
}

func TestExplainMethod(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", methodTest{
		input:    &ExpectedExplain{db: &MockDB{name: "foo"}},
		standard: "DB.Explain()",
		verbose:  "DB(foo).Explain(ctx, ?)",
	})
	tests.Add("query", methodTest{
		input:    &ExpectedExplain{db: &MockDB{name: "foo"}, query: map[string]string{"foo": "bar"}},
		standard: "DB.Explain()",
		verbose:  "DB(foo).Explain(ctx, map[foo:bar])",
	})
	tests.Run(t, testMethod)
}

func TestCreateDocMethod(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("empty", methodTest{
		input:    &ExpectedCreateDoc{db: &MockDB{name: "foo"}},
		standard: "DB.CreateDoc()",
		verbose:  "DB(foo).CreateDoc(ctx, ?)",
	})
	tests.Add("docs", methodTest{
		input:    &ExpectedCreateDoc{db: &MockDB{name: "foo"}, doc: map[string]string{"foo": "bar"}},
		standard: "DB.CreateDoc()",
		verbose:  "DB(foo).CreateDoc(ctx, map[foo:bar])",
	})
	tests.Add("options", methodTest{
		input:    &ExpectedCreateDoc{db: &MockDB{name: "foo"}, options: map[string]interface{}{"foo": "bar"}},
		standard: "DB.CreateDoc()",
		verbose:  "DB(foo).CreateDoc(ctx, ?, map[foo:bar])",
	})
	tests.Run(t, testMethod)
}
