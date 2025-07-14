package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func RedisRunner() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Increment the counter
	val, err := rdb.Incr(ctx, "pageviews").Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Pageviews: %d\n", val) // Pageviews: 1 (then 2, 3, ...)
}
