package handler

import (
	"context"
	"github.com/gudongkun/single_common/custom_gorm"
	"github.com/gudongkun/single_ucenter/enlight_ucenter_client/proto/user"
	"github.com/gudongkun/single_ucenter/models"
)

type User struct{}

// GetName Call is a single request handler called via client.Call or the generated client code
func (e *User) GetName(ctx context.Context, req *user.UserId, rsp *user.UserName) error {
	rsp.Name = "周泽楷"
	custom_gorm.Db(ctx).Create(&models.User{Name: "周泽楷"})
	//log.Info("GetName", req)
	return nil
}

// GetAge Stream is a server side stream handler called via client.Stream or the generated client code
func (e *User) GetAge(ctx context.Context, req *user.UserId, rsp *user.UserAge) error {
	rsp.Age = 23
	//log.Info("GetAge")
	return nil
}

// GetSex PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *User) GetSex(ctx context.Context, req *user.UserId, rsp *user.UserSex) error {
	rsp.Sex = "男"
	//log.Info("GetSex")
	return nil
}
