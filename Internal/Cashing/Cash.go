package cashing

import (
	"sync"

	"github.com/coocood/freecache"
)

type Cache struct {
	store     *freecache.Cache
	hits      map[string]uint32
	hitsMutex sync.RWMutex
}

func NewCache(size int) *Cache {
	return &Cache{
		store: freecache.NewCache(size),
		hits:  make(map[string]uint32),
	}
}

func (c *Cache) Set(key string, value []byte, baseTTL int) {
	c.hitsMutex.RLock()
	hits := c.hits[key]
	c.hitsMutex.RUnlock()

	ttl := baseTTL

	if hits >= 10 {
		ttl *= 2
	}

	c.store.Set([]byte(key), value, ttl)
}

func (c *Cache) Get(key string) ([]byte, error) {
	val, err := c.store.Get([]byte(key))
	if err != nil {
		return nil, err
	}

	c.hitsMutex.Lock()

	if c.hits[key] < 11 {
		c.hits[key]++
	}

	c.hitsMutex.Unlock()

	return val, nil
}

func (c *Cache) Delete(key string) {
	c.store.Del([]byte(key))

	c.hitsMutex.Lock()
	delete(c.hits, key)
	c.hitsMutex.Unlock()
}

func (c *Cache) ClearHits(key string) {
	c.hitsMutex.Lock()
	delete(c.hits, key)
	c.hitsMutex.Unlock()
}
