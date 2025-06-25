package admin

import (
	"blog/internal/logic/admin"
	"context"

	"blog/api/admin/v1"
)

// 获取登录信息

// Info 获取管理员信息
// 该函数调用 admin.Info 方法获取管理员信息，并将信息封装到 InfoRes 结构体中返回
// 参数:
//
//	ctx context.Context: 上下文对象，用于传递请求范围的请求头、超时设置等
//	req *v1.InfoReq: 请求参数对象，包含请求相关信息（本例中未使用）
//
// 返回值:
//
//	*v1.InfoRes: 包含管理员信息的响应对象
//	error: 错误对象，如果执行过程中遇到错误则返回
func (c *ControllerV1) Info(ctx context.Context, req *v1.InfoReq) (*v1.InfoRes, error) {
	// 调用 admin.Info 方法获取管理员信息
	adminInfo, err := admin.Info(ctx)
	// 如果获取信息过程中发生错误，直接返回错误
	if err != nil {
		return nil, err
	}

	// 将获取到的管理员信息封装到 InfoRes 结构体中，并作为函数返回值返回
	return &v1.InfoRes{
		Username:  adminInfo.Username,
		Nickname:  adminInfo.Nickname,
		Avatar:    adminInfo.Avatar,
		Register:  adminInfo.Register,
		LastLogin: adminInfo.LastLogin,
	}, nil
}
