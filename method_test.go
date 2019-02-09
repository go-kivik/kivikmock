package kivikmock

import (
	"testing"

	"github.com/flimzy/testy"
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
		input:    &ExpectedDBExists{name: "foo"},
		standard: "DBExists()",
		verbose:  `DBExists(ctx, "foo", ?)`,
	})
	tests.Add("options", methodTest{
		input:    &ExpectedDBExists{options: map[string]interface{}{"foo": 321}},
		standard: "DBExists()",
		verbose:  `DBExists(ctx, ?, map[foo:321])`,
	})
	tests.Add("full", methodTest{
		input:    &ExpectedDBExists{name: "foo", options: map[string]interface{}{"foo": 321}},
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
		input:    &ExpectedDestroyDB{name: "foo"},
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
