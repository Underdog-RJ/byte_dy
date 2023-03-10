package main

import (
	"api-gateway/weblib"
	"api-gateway/wrappers"
	"service_common/services"
	"service_common/services/service"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
)

func main() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("159.27.184.52:2379"),
	)
	// 用户
	userMicroService := micro.NewService(
		micro.Name("userService.client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)

	// 用户服务调用实例
	userService := services.NewUserService("rpcUserService", userMicroService.Client())

	// 视频
	videoMicroService := micro.NewService(
		micro.Name("videoService.client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)
	videoService := services.NewVideoService("rpcVideoService", videoMicroService.Client())

	// 互动
	interactionMicroService := micro.NewService(
		micro.Name("interactionService.client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)

	likeService := services.NewLikeService("rpcInteractionService", interactionMicroService.Client())

	// 社交
	socialMicroService := micro.NewService(
		micro.Name("socialService.client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)
	socialService := services.NewSocialService("rpcSocialService", socialMicroService.Client())
	commentService := service.NewCommentService("rpcInteractionService", interactionMicroService.Client())
	//创建微服务实例，使用gin暴露http接口并注册到etcd
	server := web.NewService(
		web.Name("httpService"),
		web.Address("0.0.0.0:4000"),
		//将服务调用实例使用gin处理
		web.Handler(weblib.NewRouter(userService, videoService, socialService, likeService, commentService)),
		web.Registry(etcdReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	//接收命令行参数
	_ = server.Init()
	_ = server.Run()
}
