package hw04_lru_cache //nolint:golint,stylecheck

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*listItem
	mx       sync.Mutex
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mx.Lock()
	defer c.mx.Unlock()
	if listItem, ok := c.items[key]; ok {
		listItem.Value.(*cacheItem).value = value
		c.queue.MoveToFront(listItem)
		return true
	}

	cItem := &cacheItem{key: key, value: value}
	listItem := c.queue.PushFront(cItem)
	c.items[key] = listItem

	if c.queue.Len() > c.capacity {
		lastItem := c.queue.Back()
		c.queue.Remove(lastItem)
		key := lastItem.Value.(*cacheItem).key
		delete(c.items, key)
	}

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	if listItem, ok := c.items[key]; ok {
		c.queue.MoveToFront(listItem)
		if listItem.Value.(*cacheItem) == nil {
			return nil, false
		}

		return listItem.Value.(*cacheItem).value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.mx.Lock()
	defer c.mx.Unlock()
	for k, v := range c.items {
		delete(c.items, k)
		c.queue.Remove(v)
	}
	c.queue = NewList()
}

func NewCache(capacity int) Cache {
	l := NewList()
	return &lruCache{
		capacity: capacity,
		queue:    l,
		items:    make(map[Key]*listItem),
	}
}
