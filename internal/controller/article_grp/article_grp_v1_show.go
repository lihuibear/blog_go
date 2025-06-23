package article_grp

import (
	"blog/internal/logic/article_grp"
	"context"

	"blog/api/article_grp/v1"
)

func (c *ControllerV1) Show(ctx context.Context, req *v1.ShowReq) (*v1.ShowRes, error) {
	info, err := article_grp.Show(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.ShowRes{
		ArticleGrp: info,
	}, nil
}
