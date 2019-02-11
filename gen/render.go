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

func RenderMockGo(filename string, same []*Method) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	for _, m := range same {
		fmt.Printf("%s", m.Name)
	}
	return tmpl.ExecuteTemplate(file, "mockgo.tmpl", same)
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
		args = append(args, fmt.Sprintf("arg%d %s", i, arg.Name()))
	}
	if m.AcceptsOptions {
		args = append(args, "options map[string]interface{}")
	}
	return strings.Join(args, ", ")
}

func (m *Method) ReturnArgs() string {
	args := make([]string, 0, len(m.Returns)+1)
	for _, arg := range m.Returns {
		args = append(args, arg.Name())
	}
	if m.ReturnsError {
		args = append(args, "error")
	}
	if len(args) > 1 {
		return `(` + strings.Join(args, ", ") + `)`
	}
	return args[0]
}

func (m *Method) VariableDefinitions(indent int) string {
	args := make([]string, 0, len(m.Accepts)+len(m.Returns)+2)
	types := make([]string, 0, len(args))
	if m.DBMethod {
		args = append(args, "db")
		types = append(types, "*MockDB")
	}
	for i, arg := range m.Accepts {
		args = append(args, fmt.Sprintf("arg%d", i))
		types = append(types, fmt.Sprintf("%s", arg))
	}
	if m.AcceptsOptions {
		args = append(args, "options")
		types = append(types, "map[string]interface{}")
	}
	for i, ret := range m.Returns {
		args = append(args, fmt.Sprintf("ret%d", i))
		types = append(types, fmt.Sprintf("%s", ret))
	}
	var maxLen int
	for _, arg := range args {
		if l := len(arg); l > maxLen {
			maxLen = l
		}
	}
	final := make([]string, len(args))
	for i, arg := range args {
		final[i] = fmt.Sprintf("%s%*s %s", strings.Repeat("\t", indent), -maxLen, arg, types[i])
	}
	return strings.Join(final, "\n")
}

func (m *Method) Variables(indent int) string {
	args := make([]string, 0, len(m.Accepts)+1)
	for i := range m.Accepts {
		args = append(args, fmt.Sprintf("arg%d", i))
	}
	if m.AcceptsOptions {
		args = append(args, "options")
	}
	for i := range m.Returns {
		args = append(args, fmt.Sprintf("ret%d", i))
	}
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
		args = append(args, fmt.Sprintf("%q", reflect.Zero(arg).Interface()))
	}
	args = append(args, "err")
	return strings.Join(args, ", ")
}

func (m *Method) ExpectedReturns() string {
	args := make([]string, 0, len(m.Returns))
	for i := range m.Returns {
		args = append(args, fmt.Sprintf("arg%d", i))
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
		args[i] = fmt.Sprintf("ret%d %s", i, ret.String())
	}
	return strings.Join(args, ", ")
}
