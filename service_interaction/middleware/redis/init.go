package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var RedisCtx = context.Background()

var RdbLike *redis.Client //key:userId,value:VideoId

func InitRedis() {
	addr := viper.GetString("redis.addr")
	port := viper.GetString("redis.port")
	password := viper.GetString("redis.password")
	RdbLike = redis.NewClient(&redis.Options{
		Addr:     addr + ":" + port,
		Password: password,
		DB:       5,
	})

}
