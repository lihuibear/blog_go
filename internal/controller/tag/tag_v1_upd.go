package tag

import (
	"blog/internal/logic/tag"
	"context"

	"blog/api/tag/v1"
)

func (c *ControllerV1) Upd(ctx context.Context, req *v1.UpdReq) (res *v1.UpdRes, err error) {
	err = tag.Upd(ctx, req.Id, req.GrpId, req.Name)
	return nil, err
}
