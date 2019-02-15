package kivikmock

import "time"

type item struct {
	delay time.Duration
	item  interface{}
}

type iter struct {
	items     []*item
	closeErr  error
	resultErr error
}

func (i *iter) Close() error { return i.closeErr }

func (i *iter) push(item *item) {
	i.items = append(i.items, item)
}

func (i *iter) unshift() (item *item) {
	if len(i.items) == 0 {
		return nil
	}
	item, i.items = i.items[0], i.items[1:]
	return item
}

func (i *iter) count() int {
	if len(i.items) == 0 {
		return 0
	}
	var count int
	for _, result := range i.items {
		if result != nil && result.item != nil {
			count++
		}
	}

	return count
}
