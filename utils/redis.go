package utils

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Cache struct{}

var ctx = context.Background()
var rdb *redis.Client

func SetupRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func (r Cache) Set(key string, value any) {
	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		Logger(err)
	}
}

func (r Cache) Get(key string) string {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		Logger(err)

		return ""
	}

	return val
}

func (r Cache) Remove(key string) {
	err := rdb.Do(ctx, "DEL", key).Err()
	if err != nil {
		Logger(err)
	}
}
