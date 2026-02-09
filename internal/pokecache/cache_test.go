package pokecache_test

import (
	"github.com/rQxwX3/pokedex/internal/pokecache"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second

	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for _, c := range cases {
		cache := pokecache.NewCache(interval)
		cache.Add(c.key, c.val)

		val, ok := cache.Get(c.key)
		if !ok {
			t.Errorf("expected to find key %s", c.key)
			return
		}

		if string(val) != string(c.val) {
			t.Errorf("expected value %s, got: %s", string(c.val), string(val))
			return
		}
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = 2 * baseTime

	var key = "https://example.com"
	var val = []byte("testdata")

	cache := pokecache.NewCache(baseTime)
	cache.Add(key, val)

	_, ok := cache.Get(key)
	if !ok {
		t.Errorf("expected to find key %s", key)
		return
	}

	time.Sleep(waitTime)

	if _, ok = cache.Get(key); ok {
		t.Errorf("expected to not find key %s", key)
		return
	}
}
