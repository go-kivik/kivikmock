package kivikmock

import (
	"errors"
	"fmt"
	"testing"

	"github.com/flimzy/diff"
	"github.com/flimzy/kivik"
	"github.com/flimzy/testy"
)

type fmtTest struct {
	input    fmt.Stringer
	format   string
	expected string
}

func testFmt(t *testing.T, test fmtTest) {
	result := fmt.Sprintf(test.format, test.input)
	if d := diff.Text(test.expected, result); d != nil {
		t.Error(d)
	}
}

func TestClusterSetupFormatter(t *testing.T) {
	tests := testy.NewTable()

	tests.Add("string empty", fmtTest{
		input:    &ExpectedClusterSetup{},
		format:   "%s",
		expected: `ExpectedClusterSetup => expecting ClusterSetup`,
	})
	tests.Add("default empty", fmtTest{
		input:    &ExpectedClusterSetup{},
		format:   "%v",
		expected: `ExpectedClusterSetup => expecting ClusterSetup`,
	})
	tests.Add("string error", fmtTest{
		input:    &ExpectedClusterSetup{commonExpectation: commonExpectation{err: errors.New("foo")}},
		format:   "%s",
		expected: `ExpectedClusterSetup => expecting ClusterSetup which should return an error`,
	})
	tests.Add("string action", fmtTest{
		input:    &ExpectedClusterSetup{action: map[string]string{"foo": "bar"}},
		format:   "%s",
		expected: `ExpectedClusterSetup => expecting ClusterSetup which has the desired action`,
	})
	tests.Add("string unmarshalable action and error", fmtTest{
		input:    &ExpectedClusterSetup{action: func() {}, commonExpectation: commonExpectation{err: errors.New("foo")}},
		format:   "%s",
		expected: `ExpectedClusterSetup => expecting ClusterSetup which has the desired action and should return an error`,
	})
	tests.Add("verbose empty", fmtTest{
		input:  &ExpectedClusterSetup{},
		format: "%+v",
		expected: `ExpectedClusterSetup => expecting ClusterSetup which:
	- expects any action`,
	})
	tests.Add("verbose error", fmtTest{
		input:  &ExpectedClusterSetup{commonExpectation: commonExpectation{err: errors.New("foo")}},
		format: "%+v",
		expected: `ExpectedClusterSetup => expecting ClusterSetup which:
	- expects any action
	- should return error: foo`,
	})
	tests.Add("verbose action and error", fmtTest{
		input:  &ExpectedClusterSetup{action: map[string]string{"foo": "bar"}, commonExpectation: commonExpectation{err: errors.New("foo")}},
		format: "%+v",
		expected: `ExpectedClusterSetup => expecting ClusterSetup which:
	- expects the following action:
		{
		  "foo": "bar"
		}
	- should return error: foo`,
	})
	tests.Add("verbose unmarshalable action and error", fmtTest{
		input:  &ExpectedClusterSetup{action: func() {}},
		format: "%+v",
		expected: `ExpectedClusterSetup => expecting ClusterSetup which:
	- expects the following action:
		<<unmarshalable: json: unsupported type: func()>>`,
	})
	tests.Run(t, testFmt)
}

func TestAuthenticateFormatter(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("string empty", fmtTest{
		input:    &ExpectedAuthenticate{},
		format:   "%s",
		expected: `ExpectedAuthenticate => expecting Authenticate`,
	})
	tests.Add("default empty", fmtTest{
		input:    &ExpectedAuthenticate{},
		format:   "%v",
		expected: `ExpectedAuthenticate => expecting Authenticate`,
	})
	tests.Add("string error", fmtTest{
		input:    &ExpectedAuthenticate{commonExpectation: commonExpectation{err: errors.New("foo")}},
		format:   "%s",
		expected: `ExpectedAuthenticate => expecting Authenticate which should return an error`,
	})
	tests.Add("string authenticator", fmtTest{
		input:    &ExpectedAuthenticate{authType: "foo"},
		format:   "%s",
		expected: `ExpectedAuthenticate => expecting Authenticate which expects an authenticator of type 'foo'`,
	})
	tests.Add("verbose empty", fmtTest{
		input:  &ExpectedAuthenticate{},
		format: "%+v",
		expected: `ExpectedAuthenticate => expecting Authenticate which:
	- expects any authenticator`,
	})
	tests.Add("verbose error", fmtTest{
		input:  &ExpectedAuthenticate{commonExpectation: commonExpectation{err: errors.New("foo")}},
		format: "%+v",
		expected: `ExpectedAuthenticate => expecting Authenticate which:
	- expects any authenticator
	- should return error: foo`,
	})
	tests.Add("verbose authenticator", fmtTest{
		input:  &ExpectedAuthenticate{authType: "foo"},
		format: "%+v",
		expected: `ExpectedAuthenticate => expecting Authenticate which:
	- expects an authenticator of type: foo`,
	})

	tests.Run(t, testFmt)
}

func TestFormatters(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("string ExpectedAllDBs empty", fmtTest{
		input:  &ExpectedAllDBs{},
		format: "%s",
		expected: `ExpectedAllDBs => expecting AllDBs which:
	- is without options`,
	})
	tests.Add("string ExpectedAllDBs options", fmtTest{
		input:  &ExpectedAllDBs{options: kivik.Options{"foo": "bar"}},
		format: "%s",
		expected: `ExpectedAllDBs => expecting AllDBs which:
	- is with options map[foo:bar]`,
	})
	tests.Add("string ExpectedAllDBs return", fmtTest{
		input:  &ExpectedAllDBs{results: []string{"foo", "bar"}},
		format: "%s",
		expected: `ExpectedAllDBs => expecting AllDBs which:
	- is without options
	- should return: [foo bar]`,
	})
	tests.Add("string ExpectedAllDBs error", fmtTest{
		input:  &ExpectedAllDBs{commonExpectation: commonExpectation{err: errors.New("foo")}},
		format: "%s",
		expected: `ExpectedAllDBs => expecting AllDBs which:
	- is without options
	- should return error: foo`,
	})
	tests.Add("string ExpectedClose empty", fmtTest{
		input:    &ExpectedClose{},
		format:   "%s",
		expected: `ExpectedClose => expecting client Close`,
	})
	tests.Add("string ExpectedClose error", fmtTest{
		input:    &ExpectedClose{commonExpectation: commonExpectation{err: errors.New("foo")}},
		format:   "%s",
		expected: `ExpectedClose => expecting client Close, which should return error: foo`,
	})
	tests.Add("default ExpectedClose error", fmtTest{
		input:    &ExpectedClose{commonExpectation: commonExpectation{err: errors.New("foo")}},
		format:   "%v",
		expected: `ExpectedClose => expecting client Close, which should return error: foo`,
	})
	tests.Add("verbose ExpectedClose error", fmtTest{
		input:    &ExpectedClose{commonExpectation: commonExpectation{err: errors.New("foo")}},
		format:   "%+v",
		expected: `ExpectedClose => expecting client Close, which should return error: foo`,
	})

	tests.Run(t, testFmt)
}
