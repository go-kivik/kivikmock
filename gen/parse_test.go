package main

import (
	"context"
	"reflect"
	"testing"

	"github.com/flimzy/diff"
	"github.com/flimzy/testy"
)

type testDriver interface {
	WithCtx(context.Context) error
	NoCtx(string) error
	WithOptions(string, map[string]interface{})
}

type empty interface{}

func TestDriverMethods(t *testing.T) {
	type tst struct {
		input    interface{}
		expected []*DriverMethod
		err      string
	}
	tests := testy.NewTable()
	tests.Add("non-struct", tst{
		input: 123,
		err:   "input must be struct",
	})
	tests.Add("wrong field name", tst{
		input: struct{ Y int }{},
		err:   "wrapper struct must have a single field: X",
	})
	tests.Add("non-interface", tst{
		input: struct{ X int }{},
		err:   "field X must be of type interface",
	})
	tests.Add("testDriver", tst{
		input: struct{ X testDriver }{},
		expected: []*DriverMethod{
			{
				Name:         "NoCtx",
				ReturnsError: true,
				Accepts:      []reflect.Type{typeString},
			},
			{
				Name:           "WithCtx",
				AcceptsContext: true,
				ReturnsError:   true,
			},
			{
				Name:           "WithOptions",
				AcceptsOptions: true,
				Accepts:        []reflect.Type{typeString},
			},
		},
	})

	tests.Run(t, func(t *testing.T, test tst) {
		result, err := parseDriverMethods(test.input)
		testy.Error(t, test.err, err)
		if d := diff.Interface(test.expected, result); d != nil {
			t.Error(d)
		}
	})
}
