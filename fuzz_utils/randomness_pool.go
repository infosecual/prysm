package fuzz_utils

import (
	"math/rand"
	"sync"
)

var cache *Cache

type byteCache struct {
	bytes [][]byte
}

func (b *byteCache) Add(field []byte) {
	b.bytes = append(b.bytes, field)
}

func (b *byteCache) GetRandom(rnd int) []byte {
	return b.bytes[rnd%len(b.bytes)]
}

type Cache struct {
	bytes   map[int]*byteCache
	integer []uint64
	mu      sync.Mutex
}

func (c *Cache) AddBytes(field []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	elem := c.bytes[len(field)]
	if elem == nil {
		elem = new(byteCache)
	}
	elem.Add(field)
}

func (c *Cache) AddUint64(i uint64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.integer = append(c.integer, i)
}

func (c *Cache) GetBytes(size int, rnd int) []byte {
	if cache, ok := c.bytes[size]; ok {
		return cache.GetRandom(rnd)
	}
	bytes := make([]byte, size)
	rng := rand.New(rand.NewSource(int64(rnd)))
	rng.Read(bytes)
	return bytes
}

func (c *Cache) GetUint64(rnd int) uint64 {
	return c.integer[rnd%len(c.integer)]
}
