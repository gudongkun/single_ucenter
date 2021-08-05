package handler

import (
	"context"
	"github.com/gudongkun/single_ucenter/enlight_ucenter_client/user_proto"
)

type User struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *User) GetName(ctx context.Context, req *user_proto.UserId, rsp *user_proto.UserName) error {
	rsp.Name = "周泽楷"
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *User) GetAge(ctx context.Context, req *user_proto.UserId, rsp *user_proto.UserAge) error {
	rsp.Age = 23
	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *User) GetSex(ctx context.Context, req *user_proto.UserId, rsp *user_proto.UserSex) error {
	rsp.Sex = "男"
	return nil
}
