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
	if _, ok := c.Map[key]; !ok {
		c.Map[key] = CacheEntry{time.Now(), val}
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	if entry, ok := c.Map[key]; !ok {
		return nil, false
	} else {
		return entry.val, true
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		for key, entry := range c.Map {
			if time.Since(entry.createdAt) > interval {
				delete(c.Map, key)
			}
		}
	}

}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{}

	go cache.reapLoop(interval)

	return cache
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}
