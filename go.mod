module github.com/gudongkun/single_ucenter

go 1.14

require (
	github.com/casbin/casbin/v2 v2.25.5
	github.com/casbin/xorm-adapter/v2 v2.3.2
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/gudongkun/single_common v0.0.0-20210820114953-e63fdb066e3e
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/mqtt/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/opentracing/opentracing-go v1.2.0
	google.golang.org/protobuf v1.27.1
)
