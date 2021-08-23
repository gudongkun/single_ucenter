package global

import (
	"github.com/micro/go-micro/v2/config"
)

type Config struct {
	Service     string  `json:"service"`
	ServiceAddr     string  `json:"serviceAddr"`
	Consul  string  `json:"consul"`
	Dsn     string  `json:"dsn"`
	CasbinDsn     string  `json:"casbinDsn"`
	Jaeger  string  `json:"jaeger"`
	JwtSecret  string  `json:"jwtSecret"`
}

var Conf Config

func init() {
	config.LoadFile("./config.json")
	config.Scan(&Conf)
}

// 加载 json 配置文件
