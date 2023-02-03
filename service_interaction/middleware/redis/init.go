package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var RdbLikeUserId *redis.Client  //key:userId,value:VideoId
var RdbLikeVideoId *redis.Client //key:VideoId,value:userId

func InitRedis() {
	addr := viper.GetString("redis.addr")
	port := viper.GetString("redis.port")
	password := viper.GetString("redis.password")
	RdbLikeUserId = redis.NewClient(&redis.Options{
		Addr:     addr + ":" + port,
		Password: password,
		DB:       5,
	})

	RdbLikeVideoId = redis.NewClient(&redis.Options{
		Addr:     addr + ":" + port,
		Password: password,
		DB:       5,
	})
}
