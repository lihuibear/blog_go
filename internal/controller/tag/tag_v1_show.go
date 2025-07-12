package tag

import (
	"blog/internal/logic/tag"
	"context"

	"blog/api/tag/v1"
)

func (c *ControllerV1) Show(ctx context.Context, req *v1.ShowReq) (res *v1.ShowRes, err error) {
	info, err := tag.Show(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.ShowRes{
		Tag: info,
	}, nil
}
