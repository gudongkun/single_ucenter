package casbin

import (
	"fmt"
	"github.com/gudongkun/single_ucenter/global"
	"testing"
)



func TestGetUser(t *testing.T) {
	global.InitCasbin("root:123456@(localhost:3306)/")
	global.CasbinEnforcer.AddPolicy("alice", "data1", "read")
	global.CasbinEnforcer.AddPolicy("bob", "data2", "write")
	global.CasbinEnforcer.AddPolicy("data2_admin", "data2", "read")
	global.CasbinEnforcer.AddPolicy("data2_admin", "data2", "write")

	global.CasbinEnforcer.AddRoleForUser("alice", "data2_admin")
	global.CasbinEnforcer.SavePolicy()


	ok,err := global.CasbinEnforcer.Enforce("alice","data2","read")
	fmt.Println("result",ok,err)

}