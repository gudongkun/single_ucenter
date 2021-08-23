package global

import (

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/xorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"


)
// casbin 地址 https://casbin.org/docs/en/model-storage

var CasbinEnforcer * casbin.Enforcer

func InitCasbin( dsn string )  {
	adapter,_ := xormadapter.NewAdapter("mysql",dsn)
	// Initialize the model from Go code.
	m,_ := model.NewModelFromString(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act || r.sub == "root"
	`)

	CasbinEnforcer,_ = casbin.NewEnforcer(m,adapter)
	CasbinEnforcer.EnableLog(true)
	//何时冲击加载？：重新加载接口。
	CasbinEnforcer.LoadPolicy()
}
