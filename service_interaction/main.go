package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/spf13/viper"
	"interaction/config"
	"interaction/db"
	"interaction/handle"
	"interaction/middleware/rabbitmq"
	"interaction/middleware/redis"
	"interaction/service/service"
)

func main() {
	config.InitConfig()
	db.InitDB()
	redis.InitRedis()
	rabbitmq.InitRabbitMQ()
	rabbitmq.InitLikeRabbitMQ()
	fmt.Println("init success")

	//etcdAddr := viper.GetString("etcd.addr")

	etcdReg := etcd.NewRegistry(
		registry.Addrs("159.27.184.52:2379"),
	)
	serviceName := viper.GetString("server.domain")
	serviceAddr := viper.GetString("server.addr")

	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name(serviceName), // 微服务名字
		micro.Address(serviceAddr),
		micro.Registry(etcdReg), // etcd注册件
	)

	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	service.RegisterLikeServiceHandler(microService.Server(), new(handle.LikeService))
	// 启动微服务
	_ = microService.Run()
}
