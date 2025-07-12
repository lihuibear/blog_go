package link

import (
	v1 "blog/api/link/v1"
	"blog/internal/logic/link"
	"context"
)

func (c *ControllerV1) LinkCre(ctx context.Context, req *v1.LinkCreReq) (res *v1.LinkCreRes, err error) {
	err = link.Cre(ctx, req.LinkInput)
	return
}
