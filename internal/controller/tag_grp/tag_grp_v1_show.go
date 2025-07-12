package tag_grp

import (
	"blog/api/tag_grp/v1"
	"blog/internal/logic/tag_grp"
	"context"
)

func (c *ControllerV1) Show(ctx context.Context, req *v1.ShowReq) (res *v1.ShowRes, err error) {
	info, err := tag_grp.Show(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.ShowRes{
		TagGrp: info,
	}, nil
}
