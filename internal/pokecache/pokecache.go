package pokecache

import (
	"sync"
	"time"
)

type pokecache struct {
	mu    sync.Mutex
	cache map[string]cacheEntry
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

func (p *pokecache) ReapLoop(t time.Duration) {
	p.mu.Lock()
}

func NewCache() pokecache {

	return pokecache{}
}
