package main

import (
	"testing"

	"github.com/flimzy/diff"
	"github.com/flimzy/testy"
)

func TestCompareMethods(t *testing.T) {
	type tst struct {
		client    []*Method
		driver    []*Method
		expSame   []*Method
		expClient []*Method
		expDriver []*Method
	}
	tests := testy.NewTable()
	tests.Add("one identical", tst{
		client: []*Method{
			{Name: "Foo"},
		},
		driver: []*Method{
			{Name: "Foo"},
		},
		expSame: []*Method{
			{Name: "Foo"},
		},
		expClient: []*Method{},
		expDriver: []*Method{},
	})
	tests.Add("same name", tst{
		client: []*Method{
			{Name: "Foo", ReturnsError: true},
		},
		driver: []*Method{
			{Name: "Foo"},
		},
		expSame: []*Method{},
		expClient: []*Method{
			{Name: "Foo", ReturnsError: true},
		},
		expDriver: []*Method{
			{Name: "Foo"},
		},
	})

	tests.Run(t, func(t *testing.T, test tst) {
		same, client, driver := compareMethods(test.client, test.driver)
		if d := diff.Interface(test.expSame, same); d != nil {
			t.Errorf("Same:\n%s\n", d)
		}
		if d := diff.Interface(test.expClient, client); d != nil {
			t.Errorf("Same:\n%s\n", d)
		}
		if d := diff.Interface(test.expDriver, driver); d != nil {
			t.Errorf("Same:\n%s\n", d)
		}
	})
}
