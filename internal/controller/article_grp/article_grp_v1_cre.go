package article_grp

import (
	v1 "blog/api/article_grp/v1"
	"blog/internal/logic/article_grp"
	"blog/internal/model"
	"context"
)

//func (c *ControllerV1) Cre(ctx context.Context, req *v1.CreReq) (res *v1.CreRes, err error) {
//	err = article_grp.Cre(ctx, &model.ArticleGrpInput{
//
//	}
//	return nil, gerror.NewCode(gcode.CodeNotImplemented)
//}

// Cre 新增

func (c *ControllerV1) Cre(ctx context.Context, req *v1.CreReq) (res *v1.CreRes, err error) {
	err = article_grp.Cre(ctx, &model.ArticleGrpInput{
		Name:        req.Name,
		Tags:        req.Tags,
		Description: req.Description,
		Onshow:      req.Onshow,
		Order:       req.Order,
	})
	return
}
