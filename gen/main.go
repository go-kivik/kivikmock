package main

import (
	"os"
	"reflect"
	"sort"

	"github.com/go-kivik/kivik"
	"github.com/go-kivik/kivik/driver"
)

var clientSkips = map[string]struct{}{
	"Driver":       struct{}{},
	"DSN":          struct{}{},
	"CreateDB":     struct{}{},
	"Authenticate": struct{}{},
	// to consider later
	"GetReplications": struct{}{},
	"Replicate":       struct{}{},
}
var dbSkips = map[string]struct{}{
	"Close":  struct{}{},
	"Client": struct{}{},
	"Err":    struct{}{},
	"Name":   struct{}{},
}

func main() {
	initTemplates(os.Args[1])
	if err := os.Mkdir("./other", 0777); err != nil && !os.IsExist(err) {
		panic(err)
	}
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
	dMethods, err := parseMethods(struct{ X fullClient }{}, false, clientSkips)
	if err != nil {
		return err
	}

	client, err := parseMethods(struct{ X *kivik.Client }{}, true, clientSkips)
	if err != nil {
		return err
	}
	same, cm, dm := compareMethods(client, dMethods)

	if err := RenderExpectationsGo("clientexpectations_gen.go", append(same, dm...)); err != nil {
		return err
	}
	if err := RenderClientGo("client_gen.go", append(same, dm...)); err != nil {
		return err
	}
	if err := RenderMockGo("clientmock_gen.go", append(same, cm...)); err != nil {
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
	dMethods, err := parseMethods(struct{ X fullDB }{}, false, dbSkips)
	if err != nil {
		return err
	}

	client, err := parseMethods(struct{ X *kivik.DB }{}, true, dbSkips)
	if err != nil {
		return err
	}
	same, cm, dm := compareMethods(client, dMethods)

	for _, method := range same {
		method.DBMethod = true
	}
	for _, method := range dm {
		method.DBMethod = true
	}
	for _, method := range cm {
		method.DBMethod = true
	}

	if err := RenderExpectationsGo("dbexpectations_gen.go", append(same, dm...)); err != nil {
		return err
	}
	if err := RenderClientGo("db_gen.go", append(same, dm...)); err != nil {
		return err
	}
	if err := RenderMockGo("dbmock_gen.go", append(same, cm...)); err != nil {
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
		} else {
			delete(dMethods, name)
			delete(cMethods, name)
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
