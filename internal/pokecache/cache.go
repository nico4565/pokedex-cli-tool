package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntryMap map[string]cacheEntry
	mu            *sync.Mutex
	interval      time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interv time.Duration) *Cache {
	newCachePtr := &Cache{
		cacheEntryMap: make(map[string]cacheEntry),
		mu:            &sync.Mutex{},
		interval:      interv,
	}
	go newCachePtr.reapLoop()
	return newCachePtr
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheEntryMap[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cEntry, exists := c.cacheEntryMap[key]
	if !exists {
		return nil, false
	}

	return cEntry.val, exists
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()

		for key, entry := range c.cacheEntryMap {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.cacheEntryMap, key)
			}
		}

		c.mu.Unlock()
	}
}
