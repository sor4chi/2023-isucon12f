package main

import "sync"

type cache[K comparable, V any] struct {
	sync.RWMutex
	items map[K]V
}

func NewCache[K comparable, V any]() *cache[K, V] {
	m := make(map[K]V)
	c := &cache[K, V]{
		items: m,
	}
	return c
}

func (c *cache[K, V]) Set(key K, value V) {
	c.Lock()
	c.items[key] = value
	c.Unlock()
}

func (c *cache[K, V]) Get(key K) (V, bool) {
	c.RLock()
	v, found := c.items[key]
	c.RUnlock()
	return v, found
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

type cacheInteger[K comparable, V Integer] struct {
	// Setが多いならsync.Mutex
	sync.RWMutex
	items map[K]V
}

func NewCacheInteger[K comparable, V Integer]() *cacheInteger[K, V] {
	m := make(map[K]V)
	c := &cacheInteger[K, V]{
		items: m,
	}
	return c
}

func (c *cacheInteger[K, V]) Set(key K, value V) {
	c.Lock()
	c.items[key] = value
	c.Unlock()
}

func (c *cacheInteger[K, V]) Get(key K) (V, bool) {
	c.RLock()
	v, found := c.items[key]
	c.RUnlock()
	return v, found
}

func (c *cacheInteger[K, V]) Incr(key K, value V) {
	c.Lock()
	v, found := c.items[key]
	if found {
		c.items[key] = v + value
	} else {
		c.items[key] = value
	}
	c.Unlock()
}
