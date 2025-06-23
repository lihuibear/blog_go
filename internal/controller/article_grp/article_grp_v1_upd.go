package article_grp

import (
	"blog/internal/logic/article_grp"
	"blog/internal/model"
	"context"

	"blog/api/article_grp/v1"
)

func (c *ControllerV1) Upd(ctx context.Context, req *v1.UpdReq) (res *v1.UpdRes, err error) {
	err = article_grp.Upd(ctx, req.Id, &model.ArticleGrpInput{
		Name:        req.Name,
		Tags:        req.Tags,
		Description: req.Description,
		Onshow:      req.Onshow,
		Order:       req.Order,
	})
	return
}
