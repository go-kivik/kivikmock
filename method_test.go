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
	tests.Add("auther", methodTest{
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
