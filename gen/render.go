package main

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/template"
)

var tmpl *template.Template

func initTemplates(root string) {
	var err error
	tmpl, err = template.ParseGlob(root + "/*")
	if err != nil {
		panic(err)
	}
}

func RenderExpectationsGo(filename string, methods []*Method) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	return tmpl.ExecuteTemplate(file, "expectations.go.tmpl", methods)
}

func RenderClientGo(filename string, methods []*Method) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	return tmpl.ExecuteTemplate(file, "client.go.tmpl", methods)
}

func RenderMockGo(filename string, methods []*Method) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	return tmpl.ExecuteTemplate(file, "mock.go.tmpl", methods)
}

func RenderDriverMethod(m *Method) (string, error) {
	buf := &bytes.Buffer{}
	err := tmpl.ExecuteTemplate(buf, "drivermethod.tmpl", m)
	return buf.String(), err
}

func RenderExpectedType(m *Method) (string, error) {
	buf := &bytes.Buffer{}
	err := tmpl.ExecuteTemplate(buf, "expectedtype.tmpl", m)
	return buf.String(), err
}

func RenderMock(m *Method) (string, error) {
	buf := &bytes.Buffer{}
	err := tmpl.ExecuteTemplate(buf, "mock.tmpl", m)
	return buf.String(), err
}

func (m *Method) DriverArgs() string {
	args := make([]string, 0, len(m.Accepts)+2)
	if m.AcceptsContext {
		args = append(args, "ctx context.Context")
	}
	for i, arg := range m.Accepts {
		args = append(args, fmt.Sprintf("arg%d %s", i, typeName(arg)))
	}
	if m.AcceptsOptions {
		args = append(args, "options map[string]interface{}")
	}
	return strings.Join(args, ", ")
}

func (m *Method) ReturnArgs() string {
	args := make([]string, 0, len(m.Returns)+1)
	for _, arg := range m.Returns {
		args = append(args, arg.String())
	}
	if m.ReturnsError {
		args = append(args, "error")
	}
	if len(args) > 1 {
		return `(` + strings.Join(args, ", ") + `)`
	}
	return args[0]
}

func (m *Method) VariableDefinitions() string {
	var result []string
	if m.DBMethod {
		result = append(result, "\tdb *MockDB\n")
	}
	for i, arg := range m.Accepts {
		result = append(result, fmt.Sprintf("\targ%d %s\n", i, typeName(arg)))
	}
	for i, ret := range m.Returns {
		result = append(result, fmt.Sprintf("\tret%d %s\n", i, typeName(ret)))
	}
	return strings.Join(result, "")
}

func (m *Method) inputVars() []string {
	args := make([]string, 0, len(m.Accepts)+1)
	for i := range m.Accepts {
		args = append(args, fmt.Sprintf("arg%d", i))
	}
	if m.AcceptsOptions {
		args = append(args, "options")
	}
	return args
}

func (m *Method) ExpectedVariables() string {
	args := []string{}
	if m.DBMethod {
		args = append(args, "db")
	}
	args = append(args, m.inputVars()...)
	return alignVars(0, args)
}

func (m *Method) InputVariables() string {
	var result []string
	if m.DBMethod {
		result = append(result, "\t\tdb: db.MockDB,\n")
	}
	for i := range m.Accepts {
		result = append(result, fmt.Sprintf("\t\targ%d: arg%d,\n", i, i))
	}
	if m.AcceptsOptions {
		result = append(result, fmt.Sprintf("\t\tcommonExpectation: commonExpectation{options:options},\n"))
	}
	return strings.Join(result, "")
}

func (m *Method) Variables(indent int) string {
	args := m.inputVars()
	for i := range m.Returns {
		args = append(args, fmt.Sprintf("ret%d", i))
	}
	return alignVars(indent, args)
}

func alignVars(indent int, args []string) string {
	var maxLen int
	for _, arg := range args {
		if l := len(arg); l > maxLen {
			maxLen = l
		}
	}
	final := make([]string, len(args))
	for i, arg := range args {
		final[i] = fmt.Sprintf("%s%*s %s,", strings.Repeat("\t", indent), -(maxLen + 1), arg+":", arg)
	}
	return strings.Join(final, "\n")
}

func (m *Method) ZeroReturns() string {
	args := make([]string, 0, len(m.Returns))
	for _, arg := range m.Returns {
		args = append(args, zeroValue(arg))
	}
	args = append(args, "err")
	return strings.Join(args, ", ")
}

func zeroValue(t reflect.Type) string {
	z := fmt.Sprintf("%#v", reflect.Zero(t).Interface())
	if strings.HasSuffix(z, "(nil)") {
		return "nil"
	}
	switch z {
	case "<nil>":
		return "nil"
	}
	return z
}

func (m *Method) ExpectedReturns() string {
	args := make([]string, 0, len(m.Returns))
	for i, arg := range m.Returns {
		if arg.String() == "driver.Rows" {
			args = append(args, fmt.Sprintf("&driverRows{Context: ctx, Rows: expected.ret%d}", i))
		} else {
			args = append(args, fmt.Sprintf("expected.ret%d", i))
		}
	}
	if m.AcceptsContext {
		args = append(args, "expected.wait(ctx)")
	} else {
		args = append(args, "err")
	}
	return strings.Join(args, ", ")
}

func (m *Method) ReturnTypes() string {
	args := make([]string, len(m.Returns))
	for i, ret := range m.Returns {
		args[i] = fmt.Sprintf("ret%d %s", i, typeName(ret))
	}
	return strings.Join(args, ", ")
}

func typeName(t reflect.Type) string {
	name := t.String()
	switch name {
	case "interface {}":
		return "interface{}"
	case "driver.Rows":
		return "*Rows"
	}
	return name
}

func (m *Method) SetExpectations() string {
	var args []string
	if m.DBMethod {
		args = append(args, "db: db,\n")
	}
	for i, ret := range m.Returns {
		var zero string
		switch ret.String() {
		case "*kivik.Rows":
			zero = "&Rows{}"
		case "*kivik.QueryPlan":
			zero = "&driver.QueryPlan{}"
		case "*kivik.PurgeResult":
			zero = "&driver.PurgeResult{}"
		}
		if zero != "" {
			args = append(args, fmt.Sprintf("ret%d: %s,\n", i, zero))
		}
	}
	return strings.Join(args, "")
}
