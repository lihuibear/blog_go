package tag_grp

import (
	"blog/api/tag_grp/v1"
	"blog/internal/logic/tag_grp"
	"context"
)

func (c *ControllerV1) Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error) {
	err = tag_grp.Del(ctx, req.Id)
	return nil, err
}
