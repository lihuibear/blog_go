package link

import (
	v1 "blog/api/link/v1"
	"blog/internal/logic/link"
	"context"
)

func (c *ControllerV1) LinkUpd(ctx context.Context, req *v1.LinkUpdReq) (res *v1.LinkUpdRes, err error) {
	err = link.Upd(ctx, req.Id, req.LinkInput)
	return
}
