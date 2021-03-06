package cache

import "time"

func NOOPCache() Cache {
	return &noopCache{}
}

type noopCache struct{}

func (c *noopCache) Put(key string, value interface{}, ttl time.Duration) {
	return
}

func (c *noopCache) Get(key string) (value interface{}, status Result) {
	return nil, NotFound
}

func (c *noopCache) Size() int {
	return 0
}

func (c *noopCache) Remove(key string) {
	return
}

func (c *noopCache) Len() int {
	return 0
}
