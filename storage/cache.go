package storage

import "sync"

type Cache struct {
	mapping map[string][]byte
	size    int
	mu      *sync.Mutex
}

func (c *Cache) Set(key string, data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.mapping[key] = data
	c.size += len(data)
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	data, ok := c.mapping[key]
	return data, ok
}

func (c *Cache) Size() int {
	return c.size
}

func (c *Cache) Reset() map[string][]byte {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.size = 0
	temp := c.mapping
	c.mapping = map[string][]byte{}
	return temp
}
