package enlight_ucenter_client

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

var UCenterService micro.Service

func InitService() {
	// New Service
	UCenterService = micro.NewService(
		micro.Name("enlight_ucenter_loc"),
		micro.Version("latest"),
		micro.Registry(etcd.NewRegistry(
			registry.Addrs("192.168.56.125:2379"),
		)),
	)
	// Initialise service
	//UCenterService.Init()
}
