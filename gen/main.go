package main

import (
	"os"
	"reflect"

	"github.com/flimzy/kivik"
	"github.com/flimzy/kivik/driver"
)

var clientSkips = map[string]struct{}{"CreateDB": struct{}{}}

func main() {
	initTemplates(os.Args[1])
	if err := client(); err != nil {
		panic(err)
	}
}

func client() error {
	driver, err := parseMethods(struct{ X driver.Client }{}, false)
	if err != nil {
		return err
	}
	client, err := parseMethods(struct{ X *kivik.Client }{}, true)
	if err != nil {
		return err
	}
	for i, method := range client {
		if _, ok := clientSkips[method.Name]; ok {
			client[i].Name += "_skipped"
		}
	}
	same, _, _ := compareMethods(client, driver)

	if err := RenderMockGo("expectations_gen.go", same); err != nil {
		return err
	}
	return nil
}

func compareMethods(client, driver []*Method) (same []*Method, differentClient []*Method, differentDriver []*Method) {
	dMethods := toMap(driver)
	cMethods := toMap(client)
	sameMethods := make(map[string]*Method, 0)
	for name, method := range dMethods {
		if cMethod, ok := cMethods[name]; ok {
			if reflect.DeepEqual(cMethod, method) {
				sameMethods[name] = method
				delete(dMethods, name)
				delete(cMethods, name)
			}
		}
	}
	return toSlice(sameMethods), toSlice(cMethods), toSlice(dMethods)
}

func toSlice(methods map[string]*Method) []*Method {
	result := make([]*Method, 0, len(methods))
	for _, method := range methods {
		result = append(result, method)
	}
	return result
}

func toMap(methods []*Method) map[string]*Method {
	result := make(map[string]*Method, len(methods))
	for _, method := range methods {
		result[method.Name] = method
	}
	return result
}
