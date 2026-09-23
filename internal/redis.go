package internal

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var ctx = context.Background()

func RedisClient() {
	rdb =  redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB: 0,
	})

	// defer rdb.Close()
}