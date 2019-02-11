package main

import (
	"reflect"
	"testing"

	"github.com/flimzy/diff"
	"github.com/flimzy/testy"
)

func init() {
	initTemplates("templates")
}

func TestRenderExpectedType(t *testing.T) {
	tests := testy.NewTable()
	tests.Add("CreateDoc", &Method{
		Name:           "CreateDoc",
		DBMethod:       true,
		AcceptsContext: true,
		AcceptsOptions: true,
		ReturnsError:   true,
		Accepts:        []reflect.Type{reflect.TypeOf((*interface{})(nil)).Elem()},
		Returns:        []reflect.Type{typeString, typeString},
	})

	tests.Run(t, func(t *testing.T, m *Method) {
		result, err := RenderExpectedType(m)
		if err != nil {
			t.Fatal(err)
		}
		if d := diff.Text(&diff.File{Path: "testdata/" + testy.Stub(t)}, result); d != nil {
			t.Error(d)
		}
	})
}

func TestRenderMock(t *testing.T) {

}

func TestRenderDriverMethod(t *testing.T) {
	type tst struct {
		method *Method
		err    string
	}
	tests := testy.NewTable()
	tests.Add("CreateDB", tst{
		method: &Method{
			Name:           "CreateDB",
			Accepts:        []reflect.Type{typeString},
			AcceptsContext: true,
			AcceptsOptions: true,
			ReturnsError:   true,
		},
	})
	tests.Add("No context", tst{
		method: &Method{
			Name:           "NoCtx",
			AcceptsOptions: true,
			ReturnsError:   true,
		},
	})
	tests.Run(t, func(t *testing.T, test tst) {
		result, err := RenderDriverMethod(test.method)
		testy.Error(t, test.err, err)
		if d := diff.Text(&diff.File{Path: "testdata/" + testy.Stub(t)}, result); d != nil {
			t.Error(d)
		}
	})
}

func TestVariables(t *testing.T) {
	type tst struct {
		method   *Method
		indent   int
		expected string
	}
	tests := testy.NewTable()
	tests.Add("no args", tst{
		method:   &Method{},
		expected: "",
	})
	tests.Add("one arg", tst{
		method:   &Method{Accepts: []reflect.Type{typeString}},
		expected: "arg0: arg0,",
	})
	tests.Add("one arg + options", tst{
		method: &Method{Accepts: []reflect.Type{typeString}, AcceptsOptions: true},
		expected: `arg0:    arg0,
options: options,`,
	})
	tests.Add("indent", tst{
		method: &Method{Accepts: []reflect.Type{typeString, typeString}, AcceptsOptions: true},
		indent: 2,
		expected: `		arg0:    arg0,
		arg1:    arg1,
		options: options,`,
	})

	tests.Run(t, func(t *testing.T, test tst) {
		result := test.method.Variables(test.indent)
		if d := diff.Text(test.expected, result); d != nil {
			t.Error(d)
		}
	})
}
