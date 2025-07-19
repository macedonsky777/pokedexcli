package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mu    sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cache: make(map[string]cacheEntry),
		mu:    sync.Mutex{},
	}
	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.cache[key] = entry
	defer c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.cache[key]
	if exists {
		return entry.val, true
	} else {
		return nil, false
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.cache {
			if time.Since(entry.createdAt) > interval {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}
