package main

import "github.com/go-redis/redis/v8"

type Cache struct {
	data map[string]interface{}
}

func (c *Cache) RemoveKey(key string) {
	delete(c.data, key)
}

func (c *Cache) RemoveAll() {
	c.data = make(map[string]interface{})
}

func main() {
	rdb := redis.NewClient(&redis.Options{})

	cache := &Cache{
		data: make(map[string]interface{}),
	}
	cc := rdb.ClientCache(cache)
	val, err := cc.Get(ctx, "key").Result()
	if err != nil {
		//..
	}
	cache.data["key"] = val
}

// ---------------------------------
// ---------------------------------
// go-redis
type Invalidate interface {
	RemoveKey(key string)
	RemoveAll()
}
func (c *Client) ClientCache(invalidate Invalidate) *ClientCache {
	//...
}

func (c *ClientCache) healthy() {
	for {
		if c.Ping() != nil {
			c.invalidate.RemoveAll()
		}
	}
}

func (c *ClientCache) Receive() {
	for {
		key := c.Read()
		c.invalidate.RemoveKey(key)
	}
}
