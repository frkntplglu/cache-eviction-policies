package lru

import (
	"container/list"
	"fmt"
)

type LRU[K comparable, V any] struct {
	capacity int
	cache    map[K]*list.Element
	list     *list.List
}

type entry[K comparable, V any] struct {
	key   K
	value V
}

func New[K comparable, V any](cap int) *LRU[K, V] {

	return &LRU[K, V]{
		capacity: cap,
		cache:    make(map[K]*list.Element),
		list:     list.New(),
	}
}

func (c *LRU[K, V]) Get(key K) (V, bool) {
	if el, ok := c.cache[key]; ok {
		c.list.MoveToFront(el)
		return el.Value.(entry[K, V]).value, true
	}
	var zero V
	return zero, false
}

func (c *LRU[K, V]) Put(key K, value V) {
	if el, ok := c.cache[key]; ok {
		c.list.MoveToFront(el)
		el.Value = entry[K, V]{key, value}
		return
	}

	if c.list.Len() == c.capacity {
		last := c.list.Back()
		delete(c.cache, last.Value.(entry[K, V]).key)
		c.list.Remove(last)
	}

	el := c.list.PushFront(entry[K, V]{key, value})
	c.cache[key] = el

}

func (c *LRU[K, V]) PrintAll() {
	for e := c.list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
