package kivikmock

import "time"

type delayedItem struct {
	delay time.Duration
	item  interface{}
}

type items []*delayedItem

func (i *items) push(item *delayedItem) {
	*i = append(*i, item)
}

func (i *items) unshift() (item *delayedItem) {
	itemsAry := *i
	item, *i = itemsAry[0], itemsAry[1:]
	return item
}

func (i items) count() int {
	if len(i) == 0 {
		return 0
	}
	var count int
	for _, result := range i {
		if result != nil && result.item != nil {
			count++
		}
	}

	return count
}
