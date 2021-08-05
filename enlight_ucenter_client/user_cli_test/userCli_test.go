package user_cli_test

import (
	"context"
	"github.com/gudongkun/single_ucenter/enlight_ucenter_client"
	"github.com/gudongkun/single_ucenter/enlight_ucenter_client/user_cli"
	"testing"
)

func TestGetUser(t *testing.T) {
	enlight_ucenter_client.InitService()
	res,err := user_cli.GetUser(context.Background(),1)
	if err != nil {
		t.Error(res,err)
		t.Fail()
	}
	t.Log(res)
}



