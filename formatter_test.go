package kivikmock

import (
	"errors"
	"fmt"
	"testing"

	"github.com/flimzy/diff"
	"github.com/flimzy/kivik"
	"github.com/flimzy/testy"
)

func TestFormatters(t *testing.T) {
	type tst struct {
		input    fmt.Stringer
		format   string
		expected string
	}
	tests := testy.NewTable()
	tests.Add("string ExpectedClusterSetup empty", tst{
		input:    &ExpectedClusterSetup{},
		format:   "%s",
		expected: `ExpectedClusterSetup => expecting ClusterSetup`,
	})
	tests.Add("default ExpectedClusterSetup empty", tst{
		input:    &ExpectedClusterSetup{},
		format:   "%v",
		expected: `ExpectedClusterSetup => expecting ClusterSetup`,
	})
	tests.Add("string ExpectedClusterSetup error", tst{
		input:    &ExpectedClusterSetup{commonExpectation: commonExpectation{err: errors.New("foo")}},
		format:   "%s",
		expected: `ExpectedClusterSetup => expecting ClusterSetup which should return an error`,
	})
	tests.Add("string ExpectedClusterSetup action", tst{
		input:    &ExpectedClusterSetup{action: map[string]string{"foo": "bar"}},
		format:   "%s",
		expected: `ExpectedClusterSetup => expecting ClusterSetup which has the desired action`,
	})
	tests.Add("string ExpectedClusterSetup unmarshalable action and error", tst{
		input:    &ExpectedClusterSetup{action: func() {}, commonExpectation: commonExpectation{err: errors.New("foo")}},
		format:   "%s",
		expected: `ExpectedClusterSetup => expecting ClusterSetup which has the desired action and should return an error`,
	})
	tests.Add("verbose ExpectedClusterSetup empty", tst{
		input:  &ExpectedClusterSetup{},
		format: "%+v",
		expected: `ExpectedClusterSetup => expecting ClusterSetup which:
	- expects any action`,
	})
	tests.Add("verbose ExpectedClusterSetup error", tst{
		input:  &ExpectedClusterSetup{commonExpectation: commonExpectation{err: errors.New("foo")}},
		format: "%+v",
		expected: `ExpectedClusterSetup => expecting ClusterSetup which:
	- expects any action
	- should return error: foo`,
	})
	tests.Add("verbose ExpectedClusterSetup action and error", tst{
		input:  &ExpectedClusterSetup{action: map[string]string{"foo": "bar"}, commonExpectation: commonExpectation{err: errors.New("foo")}},
		format: "%+v",
		expected: `ExpectedClusterSetup => expecting ClusterSetup which:
	- expects the following action:
		{
		  "foo": "bar"
		}
	- should return error: foo`,
	})
	tests.Add("verbose ExpectedClusterSetup unmarshalable action and error", tst{
		input:  &ExpectedClusterSetup{action: func() {}},
		format: "%+v",
		expected: `ExpectedClusterSetup => expecting ClusterSetup which:
	- expects the following action:
		<<unmarshalable: json: unsupported type: func()>>`,
	})

	tests.Add("string ExpectedAuthenticate empty", tst{
		input:  &ExpectedAuthenticate{},
		format: "%s",
		expected: `ExpectedAuthenticate => expecting Authenticate which:
	- has any authenticator`,
	})
	tests.Add("string ExpectedAuthenticate error", tst{
		input:  &ExpectedAuthenticate{commonExpectation: commonExpectation{err: errors.New("foo")}},
		format: "%s",
		expected: `ExpectedAuthenticate => expecting Authenticate which:
	- has any authenticator
	- should return error: foo`,
	})
	tests.Add("string ExpectedAuthenticate authenticator", tst{
		input:  &ExpectedAuthenticate{authType: "foo"},
		format: "%s",
		expected: `ExpectedAuthenticate => expecting Authenticate which:
	- has authenticator of type foo`,
	})
	tests.Add("string ExpectedAllDBs empty", tst{
		input:  &ExpectedAllDBs{},
		format: "%s",
		expected: `ExpectedAllDBs => expecting AllDBs which:
	- is without options`,
	})
	tests.Add("string ExpectedAllDBs options", tst{
		input:  &ExpectedAllDBs{options: kivik.Options{"foo": "bar"}},
		format: "%s",
		expected: `ExpectedAllDBs => expecting AllDBs which:
	- is with options map[foo:bar]`,
	})
	tests.Add("string ExpectedAllDBs return", tst{
		input:  &ExpectedAllDBs{results: []string{"foo", "bar"}},
		format: "%s",
		expected: `ExpectedAllDBs => expecting AllDBs which:
	- is without options
	- should return: [foo bar]`,
	})
	tests.Add("string ExpectedAllDBs error", tst{
		input:  &ExpectedAllDBs{commonExpectation: commonExpectation{err: errors.New("foo")}},
		format: "%s",
		expected: `ExpectedAllDBs => expecting AllDBs which:
	- is without options
	- should return error: foo`,
	})
	tests.Add("string ExpectedClose empty", tst{
		input:    &ExpectedClose{},
		format:   "%s",
		expected: `ExpectedClose => expecting client Close`,
	})
	tests.Add("string ExpectedClose error", tst{
		input:    &ExpectedClose{commonExpectation: commonExpectation{err: errors.New("foo")}},
		format:   "%s",
		expected: `ExpectedClose => expecting client Close, which should return error: foo`,
	})
	tests.Add("default ExpectedClose error", tst{
		input:    &ExpectedClose{commonExpectation: commonExpectation{err: errors.New("foo")}},
		format:   "%v",
		expected: `ExpectedClose => expecting client Close, which should return error: foo`,
	})
	tests.Add("verbose ExpectedClose error", tst{
		input:    &ExpectedClose{commonExpectation: commonExpectation{err: errors.New("foo")}},
		format:   "%+v",
		expected: `ExpectedClose => expecting client Close, which should return error: foo`,
	})

	tests.Run(t, func(t *testing.T, test tst) {
		result := fmt.Sprintf(test.format, test.input)
		if d := diff.Text(test.expected, result); d != nil {
			t.Error(d)
		}
	})
}
