package user_cli

import (
	"context"
	"github.com/gudongkun/single_ucenter/enlight_ucenter_client"
	"github.com/gudongkun/single_ucenter/enlight_ucenter_client/user_proto"
)



func GetUser(ctx context.Context, uid uint64) (*user_proto.UserName, error) {
	User := user_proto.NewUserService("enlight_ucenter_loc", enlight_ucenter_client.UCenterService.Client())
	return User.GetName(ctx,&user_proto.UserId{Id: uid})
}

