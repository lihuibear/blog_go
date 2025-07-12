package tag_grp

import (
	"blog/api/tag_grp/v1"
	"blog/internal/logic/tag_grp"
	"context"
)

func (c *ControllerV1) Cre(ctx context.Context, req *v1.CreReq) (res *v1.CreRes, err error) {
	err = tag_grp.Cre(ctx, req.Name)
	return nil, err
}
