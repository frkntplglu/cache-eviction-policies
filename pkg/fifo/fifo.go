package fifo

import (
	"container/list"
	"fmt"
)

type FIFO[K comparable, V any] struct {
	capacity int
	cache    map[K]*list.Element
	queue    *list.List
}

type entry[K comparable, V any] struct {
	key   K
	value V
}

func New[K comparable, V any](cap int) *FIFO[K, V] {
	return &FIFO[K, V]{
		capacity: cap,
		cache:    make(map[K]*list.Element),
		queue:    list.New(),
	}
}

func (c *FIFO[K, V]) Get(key K) (V, bool) {
	if el, ok := c.cache[key]; ok {
		return el.Value.(*entry[K, V]).value, true
	}
	var zero V
	return zero, false
}

func (c *FIFO[K, V]) Put(key K, value V) {
	if el, ok := c.cache[key]; ok {
		el.Value.(*entry[K, V]).value = value
		return
	}

	if c.queue.Len() == c.capacity {
		oldest := c.queue.Back()
		if oldest != nil {
			e := oldest.Value.(*entry[K, V])
			delete(c.cache, e.key)
			c.queue.Remove(oldest)
		}
	}

	entry := &entry[K, V]{key: key, value: value}
	el := c.queue.PushFront(entry)
	c.cache[key] = el
}

func (c *FIFO[K, V]) PrintAll() {
	for e := c.queue.Back(); e != nil; e = e.Prev() { // en eski -> en yeni
		entry := e.Value.(*entry[K, V])
		fmt.Printf("Key:%v, Value:%v\n", entry.key, entry.value)
	}
}
