package main

import (
	"github.com/gudongkun/single_ucenter/enlight_ucenter_client"
	"github.com/gudongkun/single_ucenter/enlight_ucenter_client/user_proto"
	"github.com/gudongkun/single_ucenter/handler"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	enlight_ucenter_client.InitService()
	enlight_ucenter_client.UCenterService.Init()
	// Register Handler
	user_proto.RegisterUserHandler(enlight_ucenter_client.UCenterService.Server(), new(handler.User))
	//// Register Struct as Subscriber 发布订阅模式
	//micro.RegisterSubscriber("go.micro.service.hello", service.Server(), new(subscriber.Hello))

	// Run service
	if err := enlight_ucenter_client.UCenterService.Run(); err != nil {
		log.Fatal(err)
	}
}
