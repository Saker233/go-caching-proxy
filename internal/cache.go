package internal

import (
	"log"

	"github.com/redis/go-redis/v9"
)

func checkCache(origin string) ([]byte, bool) {
	body, err := Rdb.Get(Ctx, origin).Bytes()

	if err == redis.Nil {
		return nil, false
	}

	if err != nil {
		log.Println("Redis error:", err)
		return nil, false
	}

	return body, true
}
