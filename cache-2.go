package main

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{})

	cc := rdb.ClientCache()

	var userName string
	var orderInfo string
	err := cc.AddCache([]*redis.Item{
		{
			Val: &userName,
			Cmd: rdb.Get(ctx, "user"),
		},
		{
			Val: &orderInfo,
			Cmd: rdb.Get(ctx, "order"),
		},
	})

	fmt.Println(userName, orderInfo)
}


// ---------------------------------
// go-redis
func (c *Client) ClientCache(invalidate Invalidate) *ClientCache {
	//...
}

type Item struct {
	Val interface{}	// ptr
	Cmd redis.Cmder
}

func (c *ClientCache) AddCache(items []*Item) error {
	info := c.cmdsInfoCache.Get(item.Cmd.Name())
	if !info.ReadOnly {
		return errors.New("only supports read-only commands")
	}
	keys := info.Keys()
	if len(keys) > 1 {
		return errors.New("only supports one key")
	}

  item.Val = item.cmd.Process().Result()
	c.items.Store(key, item)
}

func (c *ClientCache) healthy() {
	for {
		if c.Ping() != nil {
			// Multiple attempts to connect
			if c.reconnect() {
				for _, item := range c.items {
					item.Val = item.Cmd.Process().Result()
				}
			}
		}
	}
}

func (c *ClientCache) Receive() {
	for {
		key := c.Read()
		item := c.items[key]
		item.Val = item.Cmd.Process().Result()
	}
}
