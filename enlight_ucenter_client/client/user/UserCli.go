package user

import (
	"context"
	"github.com/gudongkun/single_ucenter/enlight_ucenter_client"
	"github.com/gudongkun/single_ucenter/enlight_ucenter_client/proto/user"
	"github.com/micro/go-micro/v2/broker"
)

//GetName 获取用户信息， enlight_ucenter_client是一个整体，子调父的方式也不可避免。
func GetName(ctx context.Context, uid uint64) (*user.UserName, error) {
	return enlight_ucenter_client.User.GetName(ctx, &user.UserId{Id: uid})
}
//GetAge 获取用户信息， enlight_ucenter_client是一个整体，子调父的方式也不可避免。
func GetAge(ctx context.Context, uid uint64) (*user.UserAge, error) {
	return enlight_ucenter_client.User.GetAge(ctx, &user.UserId{Id: uid})
}

//GetSex 获取用户信息， enlight_ucenter_client是一个整体，子调父的方式也不可避免。
func GetSex(ctx context.Context, uid uint64) (*user.UserSex, error) {
	return enlight_ucenter_client.User.GetSex(ctx, &user.UserId{Id: uid})
}

//PubSayHi 原生 broker 方式发送消息
func PubSayHi(ctx context.Context) error {
	pubSub := enlight_ucenter_client.UCenterService.Server().Options().Broker
	err := pubSub.Connect()
	if err != nil {
		panic(err)
	}

	//data,_ := proto.Marshal(student)
	data_msg := broker.Message{
		Header: map[string]string{
			"header-name":    "publish",
			"header-version": "v1.0.0",
			"data_type":      "input",
		},
		Body: []byte("important msg"),
		//Body:data,
	}
	err = pubSub.Publish("sayHello", &data_msg)
	return err

}
