package main

import (
	"fmt"
	"github.com/gudongkun/single_common"
	"github.com/gudongkun/single_common/custom_gorm"
	"github.com/gudongkun/single_common/jaeger"
	"github.com/gudongkun/single_ucenter/enlight_ucenter_client"
	"github.com/gudongkun/single_ucenter/enlight_ucenter_client/proto/user"
	"github.com/gudongkun/single_ucenter/handler"
	"github.com/gudongkun/single_ucenter/subscriber"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	//初始化gorm
	custom_gorm.InitEngine("root:123456@(localhost:3306)/single?charset=utf8mb4")
	//初始化xorm
	//custom_xorm.InitEngine( "root:123456@(localhost:3306)/single?charset=utf8mb4")
	//初始化jaeger
	jaeger.NewJaegerTracer("single_u_center_dev", "127.0.0.1:6831")
	//初始化 用户服务
	enlight_ucenter_client.InitService("single_u_center_dev")
	enlight_ucenter_client.UCenterService.Init()
	// 注册服务处理程序
	user.RegisterUserHandler(enlight_ucenter_client.UCenterService.Server(), new(handler.User))
	// broker方式 注册消息处理
	pubSub := enlight_ucenter_client.UCenterService.Server().Options().Broker
	if err := pubSub.Connect(); err != nil {
		log.Fatal(err)
	}
	suber, _ := pubSub.Subscribe("sayHello", subscriber.SayHelloBroker)
	defer func() {
		fmt.Println("client close conn and Unsubscribe")
		pubSub.Disconnect() //关闭链接
		suber.Unsubscribe() //取消订阅
	}()


	// broker方式 注册消息处理结束
	go single_common.PrometheusBoot(8050)
	// Run service
	if err := enlight_ucenter_client.UCenterService.Run(); err != nil {
		log.Fatal(err)
	}
}
