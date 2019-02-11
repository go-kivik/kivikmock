package main

import (
	"context"
	"errors"
	"reflect"
)

type DriverMethod struct {
	// The method name
	Name string
	// Accepted values, except for context and options
	Accepts []reflect.Type
	// Return values, except for error
	Returns        []reflect.Type
	AcceptsContext bool
	AcceptsOptions bool
	ReturnsError   bool
}

var (
	typeContext = reflect.TypeOf((*context.Context)(nil)).Elem()
	typeOptions = reflect.TypeOf(map[string]interface{}{})
	typeError   = reflect.TypeOf((*error)(nil)).Elem()
	typeString  = reflect.TypeOf("")
)

func parseDriverMethods(input interface{}) ([]*DriverMethod, error) {
	t := reflect.TypeOf(input)
	if t.Kind() != reflect.Struct {
		return nil, errors.New("input must be struct")
	}
	if t.NumField() != 1 || t.Field(0).Name != "X" {
		return nil, errors.New("wrapper struct must have a single field: X")
	}
	f := t.Field(0)
	if f.Type.Kind() != reflect.Interface {
		return nil, errors.New("field X must be of type interface")
	}
	result := make([]*DriverMethod, 0, f.Type.NumMethod())
	for i := 0; i < f.Type.NumMethod(); i++ {
		m := f.Type.Method(i)
		dm := &DriverMethod{
			Name: m.Name,
		}
		accepts := make([]reflect.Type, m.Type.NumIn())
		for j := 0; j < m.Type.NumIn(); j++ {
			accepts[j] = m.Type.In(j)
		}
		if accepts[0].Kind() == reflect.Interface && accepts[0].Implements(typeContext) {
			dm.AcceptsContext = true
			accepts = accepts[1:]
		}
		if accepts[len(accepts)-1] == typeOptions {
			dm.AcceptsOptions = true
			accepts = accepts[:len(accepts)-1]
		}
		result = append(result, dm)
		if len(accepts) > 0 {
			dm.Accepts = accepts
		}

		returns := make([]reflect.Type, m.Type.NumOut())
		for j := 0; j < m.Type.NumOut(); j++ {
			returns[j] = m.Type.Out(j)
		}
		if returns[len(returns)-1] == typeError {
			dm.ReturnsError = true
			returns = returns[:len(returns)-1]
		}
		if len(returns) > 0 {
			dm.Returns = returns
		}
	}
	return result, nil
}
