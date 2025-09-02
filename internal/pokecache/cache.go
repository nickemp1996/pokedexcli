package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

type Cache struct {
	cacheEntries	map[string]cacheEntry
	mu				sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{cacheEntries: make(map[string]cacheEntry)}
    go c.reapLoop(interval)
    return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.cacheEntries[key] = cacheEntry{
		createdAt:	time.Now(),
		val:		val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	elem, ok := c.cacheEntries[key]
	if !ok {
		return nil, false
	}
	return elem.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.cacheEntries {
			if time.Since(v.createdAt) >= interval {
				delete(c.cacheEntries, k)
			}
		}
		c.mu.Unlock()
	}
}