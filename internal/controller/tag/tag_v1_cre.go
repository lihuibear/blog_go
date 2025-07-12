package tag

import (
	"blog/internal/logic/tag"
	"context"

	"blog/api/tag/v1"
)

func (c *ControllerV1) Cre(ctx context.Context, req *v1.CreReq) (res *v1.CreRes, err error) {
	err = tag.Cre(ctx, req.GrpId, req.Name)
	return nil, err
}
