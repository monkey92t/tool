package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v7"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	planP := []int{1, 2, 4, 8, 16}
	planC := []int{1, 2, 4, 8, 16, 32, 64}
	planN := []int{1, 2, 4, 8, 16}

	ctx := context.TODO()
	_ = ctx


	fmt.Printf("pool size\t concurrent\t plan req\t actual req\t fail\t time\n")

	for _, p := range planP {
		client := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs: []string{"127.0.0.1:1001", "127.0.0.1:1002", "127.0.0.1:1003"},
			PoolSize: p,
		})
		for _, c := range planC {
			for _, n := range planN {
				var (
					wg sync.WaitGroup
					fail int32
					count int32
				)
				start := time.Now()
				wg.Add(c)
				for i := 0; i < c; i++ {
					go func() {
						for x := 0; x < p * n; x++ {
							atomic.AddInt32(&count, 1)
							if err := client.Get("key").Err(); err != nil {
								atomic.AddInt32(&fail, 1)
							}
						}
						wg.Done()
					}()
				}
				wg.Wait()
				fmt.Printf("%d\t%d\t%d\t%d\t%d\t%v\n", p, c, p*c*n, atomic.LoadInt32(&count), atomic.LoadInt32(&fail), time.Now().Sub(start))
			}
		}
	}
}
