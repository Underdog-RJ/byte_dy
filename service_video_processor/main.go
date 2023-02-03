package main

import (
	"service_video_processor/conf"
	"service_video_processor/service"
)

func main() {
	forever := make(chan bool)
	conf.Init()
	service.Consumer()
	<-forever
}
