package link

import (
	v1 "blog/api/link/v1"
	"blog/internal/logic/link"
	"context"
)

func (c *ControllerV1) LinkDel(ctx context.Context, req *v1.LinkDelReq) (res *v1.LinkDelRes, err error) {
	err = link.Del(ctx, req.Id)
	return
}
