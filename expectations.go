package kivikmock

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/flimzy/diff"
	"github.com/go-kivik/kivik"
)

type expectation interface {
	fulfill()
	fulfilled() bool
	Lock()
	Unlock()
	// fmt.Formatter
	// method should return the name of the method that would trigger this
	// condition. If verbose is true, the output should disambiguate between
	// different calls to the same method.
	method(verbose bool) string
	error() error
	wait(context.Context) error
	// met is called on the actual value, and returns true if the expectation
	// is met.
	met(expectation) bool
}

// commonExpectation satisfies the expectation interface, except the String()
// and method() methods.
type commonExpectation struct {
	sync.Mutex
	triggered bool
	err       error // nolint: structcheck
	delay     time.Duration
}

func (e *commonExpectation) fulfill() {
	e.triggered = true
}

func (e *commonExpectation) fulfilled() bool {
	return e.triggered
}

func (e *commonExpectation) error() error {
	return e.err
}

// wait blocks until e.delay expires, or ctx is cancelled. If e.delay expires,
// e.err is returned, otherwise ctx.Err() is returned.
func (e *commonExpectation) wait(ctx context.Context) error {
	if e.delay == 0 {
		return e.err
	}
	t := time.NewTimer(e.delay)
	defer t.Stop()
	select {
	case <-t.C:
		return e.err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ExpectedClose is used to manage *kivik.Client.Close expectation returned
// by Mock.ExpectClose.
type ExpectedClose struct {
	commonExpectation
}

func (e *ExpectedClose) method(v bool) string {
	if v {
		return "Close(ctx)"
	}
	return "Close()"
}

func (e *ExpectedClose) met(_ expectation) bool { return true }

// WillReturnError allows setting an error for *kivik.Client.Close action.
func (e *ExpectedClose) WillReturnError(err error) *ExpectedClose {
	e.err = err
	return e
}

func (e *ExpectedClose) String() string {
	msg := "ExpectedClose => expecting client Close"
	if e.err != nil {
		msg += fmt.Sprintf(", which should return error: %s", e.err)
	}
	return msg
}

// WithDelay will cause execution of Close() to delay by duration d.
func (e *ExpectedClose) WithDelay(d time.Duration) *ExpectedClose {
	e.delay = d
	return e
}

// ExpectedAllDBs is used to manage *kivik.Client.AllDBs expectation returned
// by Mock.ExpectAllDBs.
type ExpectedAllDBs struct {
	commonExpectation
	options map[string]interface{}
	results []string
}

func (e *ExpectedAllDBs) method(v bool) string {
	if v {
		if e.options == nil {
			return "AllDBs(ctx, nil)"
		}
		return fmt.Sprintf("AllDBs(ctx, %#v)", e.options)
	}
	return "AllDBs()"
}

func (e *ExpectedAllDBs) met(ex expectation) bool {
	exp := ex.(*ExpectedAllDBs)
	if exp.options == nil {
		return true
	}
	return reflect.DeepEqual(e.options, exp.options)
}

// Format satisfies the fmt.Formatter interface.
func (e *ExpectedAllDBs) Format(f fmt.State, verb rune) {
	msg := e.method(f.Flag('+'))
	switch verb {
	case 's':
		fmt.Fprintf(f, msg)
		return
	case 'v':

	}
}

// WillReturnError allows setting an error for *kivik.Client.Close action.
func (e *ExpectedAllDBs) WillReturnError(err error) *ExpectedAllDBs {
	e.err = err
	return e
}

// WithOptions will match the provided options against actual options passed
// during execution.
func (e *ExpectedAllDBs) WithOptions(options kivik.Options) *ExpectedAllDBs {
	e.options = options
	return e
}

// WillReturn sets the expected results.
func (e *ExpectedAllDBs) WillReturn(results []string) *ExpectedAllDBs {
	e.results = results
	return e
}

// WithDelay will cause execution of AllDBs() to delay by duration d.
func (e *ExpectedAllDBs) WithDelay(d time.Duration) *ExpectedAllDBs {
	e.delay = d
	return e
}

// ExpectedAuthenticate is used to manage *kivik.Client.Authenticate
// expectation returned by Mock.ExpectAuthenticate.
type ExpectedAuthenticate struct {
	commonExpectation
	authType string
}

// Format satisfies the fmt.Formatter interface.
func (e *ExpectedAuthenticate) Format(f fmt.State, verb rune) {
	switch verb {
	case 's':
		fmt.Fprint(f, e.method(f.Flag('+'))) // nolint: errcheck
	case 'v':
		if !f.Flag('+') {
			fmt.Fprintf(f, "%s", e) // nolint: errcheck
			return
		}
		fmt.Fprintf(f, "call to %s which:", e)
		if e.authType == "" {
			fmt.Fprint(f, "\n\t- has any authenticator")
		} else {
			fmt.Fprint(f, "\n\t- has an authenticator of type: "+e.authType)
		}
		if e.err != nil {
			fmt.Fprintf(f, "\n\t- should return error: %s", e.err)
		}
	default:
		fmt.Fprintf(f, "<<unknown format verb:%s>>", string(verb))
	}
}

func (e *ExpectedAuthenticate) method(v bool) string {
	if v {
		if e.authType == "" {
			return "Authenticate(ctx, <T>)"
		}
		return fmt.Sprintf("Authenticate(ctx, <%s>)", e.authType)
	}
	return "Authenticate()"
}

func (e *ExpectedAuthenticate) met(ex expectation) bool {
	exp := ex.(*ExpectedAuthenticate)
	if exp.authType == "" {
		return true
	}
	return e.authType == exp.authType
}

// WillReturnError allows setting an error for *kivik.Client.Authenticate action.
func (e *ExpectedAuthenticate) WillReturnError(err error) *ExpectedAuthenticate {
	e.err = err
	return e
}

// WithAuthenticator will match the the provide authenticator _type_ against
// that provided. There is no way to validate the authenticated credentials
// with this method.
func (e *ExpectedAuthenticate) WithAuthenticator(authenticator interface{}) *ExpectedAuthenticate {
	e.authType = reflect.TypeOf(authenticator).Name()
	return e
}

// WithDelay will cause execution of Authenticate() to delay by duration d.
func (e *ExpectedAuthenticate) WithDelay(d time.Duration) *ExpectedAuthenticate {
	e.delay = d
	return e
}

// ExpectedClusterSetup is used to manage *kivik.Client.ClusterSetup
// expectation returned by Mock.ExpectClusterSetup.
type ExpectedClusterSetup struct {
	commonExpectation
	action interface{}
}

func (e *ExpectedClusterSetup) method(v bool) string {
	if v {
		return fmt.Sprintf("ClusterSetup(ctx, %#v)", e.action)
	}
	return "ClusterSetup()"
}

func (e *ExpectedClusterSetup) met(ex expectation) bool {
	exp := ex.(*ExpectedClusterSetup)
	if exp.action == nil {
		return true
	}
	return diff.AsJSON(e.action, exp.action) == nil
}

// Format satisfies the fmt.Formatter interface
func (e *ExpectedClusterSetup) Format(f fmt.State, verb rune) {
	switch verb {
	case 's':
		fmt.Fprint(f, e.String()) // nolint: errcheck
	case 'v':
		if !f.Flag('+') {
			fmt.Fprintf(f, "%s", e) // nolint: errcheck
			return
		}
		msg := "ExpectedClusterSetup => expecting ClusterSetup which:"
		if e.action == nil {
			msg += "\n\t- expects any action"
		} else {
			msg += "\n\t- expects the following action:"
			b, err := json.MarshalIndent(e.action, "\t\t", "  ")
			if err != nil {
				msg += fmt.Sprintf("\n\t\t<<unmarshalable: %s>>", err)
			} else {
				msg += "\n\t\t" + string(b)
			}
		}
		if e.err != nil {
			msg += fmt.Sprintf("\n\t- should return error: %s", e.err)
		}
		fmt.Fprint(f, msg) // nolint: errcheck
	}
}

func (e *ExpectedClusterSetup) String() string {
	msg := "ExpectedClusterSetup => expecting ClusterSetup"
	modifiers := []string{}
	if e.action != nil {
		modifiers = append(modifiers, "has the desired action")
	}
	if e.err != nil {
		modifiers = append(modifiers, "should return an error")
	}
	if len(modifiers) > 0 {
		msg = msg + " which " + strings.Join(modifiers, " and ")
	}
	return msg
}

// WillReturnError causes ClusterSetup to mock this return error.
func (e *ExpectedClusterSetup) WillReturnError(err error) *ExpectedClusterSetup {
	e.err = err
	return e
}

// WithAction specifies the action to be matched. Note that this expectation
// is compared with the actual action's marshaled JSON output, so it is not
// essential that the data types match exactly, in a Go sense.
func (e *ExpectedClusterSetup) WithAction(action interface{}) *ExpectedClusterSetup {
	e.action = action
	return e
}

// WithDelay will cause execution of ClusterSetups() to delay by duration d.
func (e *ExpectedClusterSetup) WithDelay(d time.Duration) *ExpectedClusterSetup {
	e.delay = d
	return e
}

// ExpectedClusterStatus is used to manage *kivik.Client.ClusterStatus
// expectation returned by Mock.ExpectClusterStatus.
type ExpectedClusterStatus struct {
	commonExpectation
	options map[string]interface{}
	status  string
}

func (e *ExpectedClusterStatus) met(ex expectation) bool {
	exp := ex.(*ExpectedClusterStatus)
	if exp.options == nil {
		return true
	}
	return reflect.DeepEqual(e.options, exp.options)
}

func (e *ExpectedClusterStatus) method(v bool) string {
	if v {
		if e.options == nil {
			return "ClusterStatus(ctx, nil)"
		}
		return fmt.Sprintf("ClusterStatus(ctx, %+v)", e.options)
	}
	return "ClusterStatus()"
}

// Format satisfies the fmt.Formatter interface
func (e *ExpectedClusterStatus) Format(f fmt.State, verb rune) {
	switch verb {
	case 's':
		fmt.Fprint(f, e.method(f.Flag('+'))) // nolint: errcheck
	case 'v':
		if !f.Flag('+') {
			fmt.Fprintf(f, "%s", e) // nolint: errcheck
			return
		}
		msg := "call to ClusterStatus() which:"
		if e.options == nil {
			msg += "\n\t- expects any options"
		} else {
			msg += fmt.Sprintf("\n\t- expects the options: %#v", e.options)
		}
		if e.err != nil {
			msg += fmt.Sprintf("\n\t- should return error: %s", e.err)
		}
		fmt.Fprint(f, msg) // nolint: errcheck
	}
}

// WithOptions sets the expectation that ClusterStatus will be called with the
// provided options.
func (e *ExpectedClusterStatus) WithOptions(options map[string]interface{}) *ExpectedClusterStatus {
	e.options = options
	return e
}

// WillReturn causes ClusterStatus to mock this return value.
func (e *ExpectedClusterStatus) WillReturn(status string) *ExpectedClusterStatus {
	e.status = status
	return e
}

// WillReturnError causes ClusterStatus to mock this return error.
func (e *ExpectedClusterStatus) WillReturnError(err error) *ExpectedClusterStatus {
	e.err = err
	return e
}

// WithDelay will cause execution of ClusterStatus() to delay by duration d.
func (e *ExpectedClusterStatus) WithDelay(d time.Duration) *ExpectedClusterStatus {
	e.delay = d
	return e
}
