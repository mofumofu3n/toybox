package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := setupRedisClient()

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("key", val)
}

func setupRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	return client
}
