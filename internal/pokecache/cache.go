package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Map map[string]CacheEntry
	mt  sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mt.Lock()

	if _, ok := c.Map[key]; !ok {
		c.Map[key] = CacheEntry{time.Now(), val}
	}

	c.mt.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mt.Lock()

	defer c.mt.Unlock()

	if entry, ok := c.Map[key]; !ok {
		return nil, false
	} else {
		return entry.val, true
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.mt.Lock()

		for key, entry := range c.Map {
			if time.Since(entry.createdAt) > interval {
				delete(c.Map, key)
			}
		}

		c.mt.Unlock()
	}

}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		Map: make(map[string]CacheEntry),
	}

	go cache.reapLoop(interval)

	return cache
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}
