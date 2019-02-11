package main

import (
	"os"
	"reflect"
	"sort"

	"github.com/go-kivik/kivik"
	"github.com/go-kivik/kivik/driver"
)

var clientSkips = map[string]struct{}{
	"CreateDB":     struct{}{},
	"Authenticate": struct{}{},
}
var dbSkips = map[string]struct{}{
	"Close": struct{}{},
}

func main() {
	initTemplates(os.Args[1])
	if err := client(); err != nil {
		panic(err)
	}
	if err := db(); err != nil {
		panic(err)
	}
}

type fullClient interface {
	driver.Client
	driver.DBsStatser
	driver.Pinger
	driver.Sessioner
	driver.Cluster
	driver.ClientCloser
	driver.Authenticator
	driver.ClientReplicator
	driver.DBUpdater
}

func client() error {
	dMethods, err := parseMethods(struct{ X fullClient }{}, false)
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
	same, _, _ := compareMethods(client, dMethods)

	if err := RenderExpectationsGo("clientexpectations_gen.go", same); err != nil {
		return err
	}
	if err := RenderMockGo("clientmock_gen.go", same); err != nil {
		return err
	}
	return nil
}

type fullDB interface {
	driver.DB
	driver.AttachmentMetaGetter
	driver.BulkDocer
	driver.BulkGetter
	driver.Copier
	driver.DBCloser
	driver.DesignDocer
	driver.Finder
	driver.Flusher
	driver.LocalDocer
	driver.MetaGetter
	driver.Purger
}

func db() error {
	dMethods, err := parseMethods(struct{ X fullDB }{}, false)
	if err != nil {
		return err
	}

	client, err := parseMethods(struct{ X *kivik.DB }{}, true)
	if err != nil {
		return err
	}
	for i, method := range client {
		if _, ok := dbSkips[method.Name]; ok {
			client[i].Name += "_skipped"
		}
	}
	same, _, _ := compareMethods(client, dMethods)

	for _, method := range same {
		method.DBMethod = true
	}

	if err := RenderExpectationsGo("dbexpectations_gen.go", same); err != nil {
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
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result
}

func toMap(methods []*Method) map[string]*Method {
	result := make(map[string]*Method, len(methods))
	for _, method := range methods {
		result[method.Name] = method
	}
	return result
}
