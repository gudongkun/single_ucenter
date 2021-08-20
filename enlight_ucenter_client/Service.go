package enlight_ucenter_client

import (
	"github.com/gudongkun/single_common/jaeger_log"
	"github.com/gudongkun/single_ucenter/enlight_ucenter_client/proto/user"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/broker/mqtt/v2"
	_ "github.com/micro/go-plugins/broker/mqtt/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	wrapperTrace "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
)

var (
	UCenterService micro.Service
	User           user.UserService //User user对象 只需要再客户端初始
)

//InitClient 初始化客户端
func InitClient(serviceName string) {
	InitService(serviceName)
	User = user.NewUserService(serviceName, UCenterService.Client())
}

//InitService 初始话 service
func InitService(serviceName string) {
	//jaegerTracer, closer, err := plugins.NewJaegerTracer(serviceName, "127.0.0.1:6831")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer closer.Close()

	// New Service etcd
	//UCenterService = micro.NewService(
	//	micro.Name("enlight_ucenter_loc"),
	//	micro.Version("latest"),
	//	micro.Registry(etcd.NewRegistry(
	//		registry.Addrs("192.168.56.125:2379"),
	//	)),
	//)
	// consul注册中心
	UCenterService = micro.NewService(
		micro.Name(serviceName),
		micro.Version("latest"),
		micro.Registry(consul.NewRegistry(func(op *registry.Options) {
			//if op.Context == nil {
			//	op.Context = context.Background()
			//}
			//op.Context = context.WithValue(op.Context, "consul_config", consulConf)
			op.Addrs = []string{"127.0.0.1:8500"}
		}),
		),
		micro.Broker(mqtt.NewBroker()), //实例化mqtt broker代理
		micro.WrapHandler(
			wrapperTrace.NewHandlerWrapper(opentracing.GlobalTracer()),
			prometheus.NewHandlerWrapper(),
			jaeger_log.NewLogWrapper,
		),
	)

	//默认注册中心
	//UCenterService = micro.NewService(
	//	micro.Name("go.micro.srv.pubsub"),
	//)
	// Initialise service
	//UCenterService.Init()
}
