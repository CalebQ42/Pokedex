package cache

import (
	"sync"
	"time"
)

type cacheValue struct {
	t    time.Time
	data []byte
}

type Cache struct {
	cMut *sync.RWMutex
	c    map[string]cacheValue
}

func NewCache(clearInterval time.Duration) *Cache {
	c := &Cache{
		cMut: &sync.RWMutex{},
		c:    make(map[string]cacheValue),
	}
	go c.reapLoop(clearInterval)
	return c
}

func (c *Cache) reapLoop(clearInterval time.Duration) {
	t := time.NewTicker(clearInterval)
	for cur := range t.C {
		cur = cur.Add(-1 * clearInterval)
		for i := range c.c {
			if c.c[i].t.Before(cur) {
				c.cMut.Lock()
				delete(c.c, i)
				c.cMut.Unlock()
			}
		}
	}
}

func (c *Cache) Get(url string) ([]byte, bool) {
	c.cMut.RLock()
	defer c.cMut.RUnlock()
	res, ok := c.c[url]
	if !ok {
		return nil, false
	}
	return res.data, true
}

func (c *Cache) Add(url string, data []byte) {
	c.cMut.Lock()
	defer c.cMut.Unlock()
	c.c[url] = cacheValue{
		t:    time.Now(),
		data: data,
	}
}
