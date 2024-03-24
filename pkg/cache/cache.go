package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client
var Ctx = context.Background()

func InitRedis(uri string) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     uri, // Redis server address
		Password: "",  // Password, leave empty if none
		DB:       0,   // Default DB
	})
}
