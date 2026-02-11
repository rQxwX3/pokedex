package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Map map[string]CacheEntry
	mt  sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}
