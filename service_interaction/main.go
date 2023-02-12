package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"interaction/config"
	"interaction/db"
	"interaction/discover"
	"interaction/handle"
	"interaction/middleware/rabbitmq"
	"interaction/middleware/redis"
	"interaction/service"
	"net"
)

func main() {
	config.InitConfig()
	db.InitDB()
	redis.InitRedis()
	rabbitmq.InitRabbitMQ()
	rabbitmq.InitLikeRabbitMQ()
	fmt.Println("init success")

	etcdAddress := []string{viper.GetString("etcd.address")}
	etcdRegister := discovery.NewRegister(etcdAddress, logrus.New())
	grpcAddress := viper.GetString("server.grpcAddress")
	defer etcdRegister.Stop()
	fmt.Println("etcd register")

	interactionNode := discovery.Server{
		Name: viper.GetString("server.domain"),
		Addr: grpcAddress,
	}
	server := grpc.NewServer()
	defer server.Stop()
	fmt.Println("grpc server")

	// 绑定service
	service.RegisterLikeServiceServer(server, handle.NewLikeService())

	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}

	if _, err := etcdRegister.Register(interactionNode, 10); err != nil {
		panic(fmt.Sprintf("start server failed, err: %v", err))
	}
	logrus.Info("server started listen on ", grpcAddress)
	fmt.Println("server started listen on ", grpcAddress)
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
