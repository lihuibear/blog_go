package admin

import (
	"blog/internal/logic/admin"
	"context"

	"blog/api/admin/v1"
)

func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	err = admin.Logout(ctx)
	return
}
