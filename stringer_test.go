package kivikmock

import (
	"errors"
	"fmt"
	"testing"

	"github.com/flimzy/diff"
	"github.com/flimzy/kivik"
	"github.com/flimzy/testy"
)

func TestStringers(t *testing.T) {
	type tst struct {
		input    fmt.Stringer
		expected string
	}
	tests := testy.NewTable()
	tests.Add("empty ExpectedClusterSetup", tst{
		input: &ExpectedClusterSetup{},
		expected: `ExpectedClusterSetup => expecting ClusterSetup which:
	- has any action`,
	})
	tests.Add("error ExpectedClusterSetup", tst{
		input: &ExpectedClusterSetup{commonExpectation: commonExpectation{err: errors.New("foo")}},
		expected: `ExpectedClusterSetup => expecting ClusterSetup which:
	- has any action
	- should return error: foo`,
	})
	tests.Add("action ExpectedClusterSetup", tst{
		input: &ExpectedClusterSetup{action: map[string]string{"foo": "bar"}},
		expected: `ExpectedClusterSetup => expecting ClusterSetup which:
	- has the action:
		{
		  "foo": "bar"
		}`,
	})
	tests.Add("unmarshalable action ExpectedClusterSetup", tst{
		input: &ExpectedClusterSetup{action: func() {}},
		expected: `ExpectedClusterSetup => expecting ClusterSetup which:
	- has the action:
		<<unmarshalable object: json: unsupported type: func()>>`,
	})
	tests.Add("ExpectedAuthenticate empty", tst{
		input: &ExpectedAuthenticate{},
		expected: `ExpectedAuthenticate => expecting Authenticate which:
	- has any authenticator`,
	})
	tests.Add("ExpectedAuthenticate error", tst{
		input: &ExpectedAuthenticate{commonExpectation: commonExpectation{err: errors.New("foo")}},
		expected: `ExpectedAuthenticate => expecting Authenticate which:
	- has any authenticator
	- should return error: foo`,
	})
	tests.Add("ExpectedAuthenticate authenticator", tst{
		input: &ExpectedAuthenticate{authType: "foo"},
		expected: `ExpectedAuthenticate => expecting Authenticate which:
	- has authenticator of type foo`,
	})
	tests.Add("ExpectedAllDBs empty", tst{
		input: &ExpectedAllDBs{},
		expected: `ExpectedAllDBs => expecting AllDBs which:
	- is without options`,
	})
	tests.Add("ExpectedAllDBs options", tst{
		input: &ExpectedAllDBs{options: kivik.Options{"foo": "bar"}},
		expected: `ExpectedAllDBs => expecting AllDBs which:
	- is with options map[foo:bar]`,
	})
	tests.Add("ExpectedAllDBs return", tst{
		input: &ExpectedAllDBs{results: []string{"foo", "bar"}},
		expected: `ExpectedAllDBs => expecting AllDBs which:
	- is without options
	- should return: [foo bar]`,
	})
	tests.Add("ExpectedAllDBs error", tst{
		input: &ExpectedAllDBs{commonExpectation: commonExpectation{err: errors.New("foo")}},
		expected: `ExpectedAllDBs => expecting AllDBs which:
	- is without options
	- should return error: foo`,
	})
	tests.Add("ExpectedClose empty", tst{
		input:    &ExpectedClose{},
		expected: `ExpectedClose => expecting client Close`,
	})
	tests.Add("ExpectedClose error", tst{
		input:    &ExpectedClose{commonExpectation: commonExpectation{err: errors.New("foo")}},
		expected: `ExpectedClose => expecting client Close, which should return error: foo`,
	})

	tests.Run(t, func(t *testing.T, test tst) {
		result := test.input.String()
		if d := diff.Text(test.expected, result); d != nil {
			t.Error(d)
		}
	})
}
