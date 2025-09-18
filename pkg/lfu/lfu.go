package lfu

import (
	"container/list"
	"fmt"
)

type LFU[K comparable, V any] struct {
	capacity int
	cache    map[K]*list.Element
	freqMap  map[int]*list.List
	minFreq  int
}

type entry[K comparable, V any] struct {
	key   K
	value V
	freq  int
}

func New[K comparable, V any](cap int) *LFU[K, V] {
	return &LFU[K, V]{
		capacity: cap,
		cache:    make(map[K]*list.Element),
		freqMap:  make(map[int]*list.List),
		minFreq:  0,
	}
}

func (c *LFU[K, V]) Get(key K) (V, bool) {
	if el, ok := c.cache[key]; ok {
		c.increment(el)
		return el.Value.(*entry[K, V]).value, true
	}
	var zero V
	return zero, false
}

func (c *LFU[K, V]) Put(key K, value V) {
	if c.capacity == 0 {
		return
	}

	if el, ok := c.cache[key]; ok {
		// update value
		el.Value.(*entry[K, V]).value = value
		c.increment(el)
		return
	}

	if len(c.cache) >= c.capacity {
		lst := c.freqMap[c.minFreq]
		toRemove := lst.Back()
		if toRemove != nil {
			e := toRemove.Value.(*entry[K, V])
			delete(c.cache, e.key)
			lst.Remove(toRemove)
		}
	}

	entry := &entry[K, V]{key: key, value: value, freq: 1}
	lst, ok := c.freqMap[1]
	if !ok {
		lst = list.New()
		c.freqMap[1] = lst
	}
	el := lst.PushFront(entry)
	c.cache[key] = el
	c.minFreq = 1
}

func (c *LFU[K, V]) increment(el *list.Element) {
	entry := el.Value.(*entry[K, V])
	oldFreq := entry.freq
	entry.freq++

	oldList := c.freqMap[oldFreq]
	oldList.Remove(el)
	if oldList.Len() == 0 && oldFreq == c.minFreq {
		c.minFreq++
	}

	newList, ok := c.freqMap[entry.freq]
	if !ok {
		newList = list.New()
		c.freqMap[entry.freq] = newList
	}
	newList.PushFront(entry)
}

func (c *LFU[K, V]) PrintAll() {
	for _, lst := range c.freqMap {
		for e := lst.Front(); e != nil; e = e.Next() {
			entry := e.Value.(*entry[K, V])
			fmt.Printf("  Key:%v, Value:%v\n", entry.key, entry.value)
		}
	}
}
