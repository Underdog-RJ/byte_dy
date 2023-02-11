package handle

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"interaction/db"
	"interaction/middleware/rabbitmq"
	"interaction/middleware/redis"
	"interaction/pkg/util"
	"interaction/service"
	"testing"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/Users/hebinglun/Desktop/go/btye_qingxun/service_interaction/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	db.InitDB()
	redis.InitRedis()
	rabbitmq.InitRabbitMQ()
	rabbitmq.InitLikeRabbitMQ()
}

func TestLikeService_LikeAction(t *testing.T) {

	req := service.LikeActionRequest{UserId: 5556, VideoId: 20, ActionType: util.ISLIKE}

	ctx := context.Background()
	l := LikeService{}
	resp, err := l.LikeAction(ctx, &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
