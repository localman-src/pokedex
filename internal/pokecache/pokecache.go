package pokecache

import (
	"sync"
	"time"
)

type pokecache struct {
	mu     sync.Mutex
	cache  map[string]cacheEntry
	expiry time.Duration
}
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (p *pokecache) Add(key string, val []byte) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (p *pokecache) Get(key string) ([]byte, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	var output []byte
	entry, ok := p.cache[key]
	if !ok {
		return output, false
	}

	return entry.val, true
}

func (p *pokecache) Set(key string) {
	if entry, ok := p.cache[key]; ok {
		entry.createdAt = time.Now().UTC()
		p.cache[key] = entry
	}

}

func (p *pokecache) ReapLoop(t time.Duration) {
	reapTime := time.NewTicker(t)
	for ; ; <-reapTime.C {
		p.mu.Lock()
		for key, entry := range p.cache {
			reapEntry := entry.createdAt.Add(t).Before(time.Now().UTC())
			if reapEntry {
				delete(p.cache, key)
			}
		}
		p.mu.Unlock()
	}
}

func NewCache(duration time.Duration) *pokecache {

	cache := pokecache{
		mu:     sync.Mutex{},
		cache:  map[string]cacheEntry{},
		expiry: duration,
	}

	go cache.ReapLoop(cache.expiry)

	return &cache
}
