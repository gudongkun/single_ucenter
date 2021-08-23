package main

import (
	"github.com/gudongkun/single_common/jaeger"
	"github.com/gudongkun/single_ucenter/enlight_ucenter_client"
	"github.com/gudongkun/single_ucenter/enlight_ucenter_client/client/user"
	wrapperTrace "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"

	"context"
	"testing"
)

func TestGetUser(t *testing.T) {
	// opentracing 处理逻辑开始
	_,cl,_:= jaeger.NewJaegerTracer("single.gateway", "127.0.0.1:6831")
	defer cl.Close() //必须close否则不上传
	// 初始化span
	ctx1, span1, _ := wrapperTrace.StartSpanFromContext(context.Background(), opentracing.GlobalTracer(), "/SomeDefaultApi")

	span1.Finish()
	//从context里面获取span
	ctx2, span2,_ := wrapperTrace.StartSpanFromContext(ctx1, opentracing.GlobalTracer(),"subService2")
	span2.Finish()
	//childof 创建span
	span3 :=opentracing.StartSpan("subService3",opentracing.ChildOf(span1.Context()))
	span3.Finish()
	// opentracing 处理逻辑结束
	enlight_ucenter_client.InitClient("single.enlight.ucenter","127.0.0.1:8500")
	res, err := user.GetName(ctx2, 1)
	if err != nil {
		t.Error(res, err)
		t.Fail()
	}
	user.GetAge(ctx2, 1)
	user.GetSex(ctx2, 1)
	user.GetSex(ctx2, 1)
	//fmt.Println("GetUser result: ", res)
	// PubSayHi 测试 broker 方式。发送消息队列
	//err2 := user.PubSayHi(context.Background())
	//if err2 != nil {
	//	t.Error(err2)
	//	t.Fail()
	//}

}
