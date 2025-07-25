package database

import (
	"context"
	c "mind-set/config"

	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

func initRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     c.Config.Redis.Host + ":" + c.Config.Redis.Port,
		Password: c.Config.Redis.Password,
		DB:       c.Config.Redis.Database,
	})
	var ctx = context.Background()
	_, err := Rdb.Ping(ctx).Result()

	if err != nil {
		panic("Redis connection failed：" + err.Error())
	}
}
