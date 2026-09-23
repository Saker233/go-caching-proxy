package internal

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client
var Ctx = context.Background()

func RedisClient() {
	Rdb =  redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB: 0,
	})

	// defer rdb.Close()
}