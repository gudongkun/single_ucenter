package subscriber

import (
	"github.com/micro/go-micro/v2/broker"
	log "github.com/micro/go-micro/v2/logger"
)

func SayHelloBroker(event broker.Event) error {

	//获取消息
	msg := event.Message()
	//获取主题
	topic := event.Topic()
	switch msg.Header["data_type"] {
	case "input":
		log.Infof("Function Process message:%+v, topic:%+v", msg, topic)
	case "struct":
		log.Infof("Function Process message::%+v, topic:%+v ", msg, topic)
	}
	return nil
}
