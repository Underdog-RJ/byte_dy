package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"service_common/services"
	"service_social/conf"
	"service_social/core"
)

func main() {
	conf.Init()
	etcdReg := etcd.NewRegistry(
		registry.Addrs("159.27.184.52:2379"),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcSocialService"), // 微服务名字
		micro.Address("127.0.0.1:8084"),
		micro.Registry(etcdReg), // etcd注册件
	)

	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	services.RegisterSocialServiceHandler(microService.Server(), new(core.SocialService))
	// 启动微服务
	_ = microService.Run()
}
