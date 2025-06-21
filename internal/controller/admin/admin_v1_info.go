package admin

import (
	"blog/internal/logic/admin"
	"context"

	"blog/api/admin/v1"
)

// 获取登录信息

func (c *ControllerV1) Info(ctx context.Context, req *v1.InfoReq) (*v1.InfoRes, error) {
	adminInfo, err := admin.Info(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.InfoRes{
		Username:  adminInfo.Username,
		Nickname:  adminInfo.Nickname,
		Avatar:    adminInfo.Avatar,
		Register:  adminInfo.Register,
		LastLogin: adminInfo.LastLogin,
	}, nil
}
