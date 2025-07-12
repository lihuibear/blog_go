package tag_grp

import (
	"blog/api/tag_grp/v1"
	"blog/internal/logic/tag_grp"
	"context"
)

func (c *ControllerV1) Upd(ctx context.Context, req *v1.UpdReq) (res *v1.UpdRes, err error) {
	err = tag_grp.Upd(ctx, req.Id, req.Name)
	return nil, err
}
