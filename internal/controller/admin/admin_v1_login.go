package admin

import (
	"blog/api/admin/v1"
	"blog/internal/logic/account"
	"blog/internal/model"
	"context"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginRes, error) {
	token, err := account.Login(ctx, &model.LoginInput{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &v1.LoginRes{
		Token: token,
	}, nil
}
