package article

import (
	"blog/internal/logic/article"
	"blog/internal/logic/article_grp"
	"blog/internal/model"
	"blog/internal/utility"
	"context"

	"blog/api/article/v1"
)

func (c *ControllerV1) Upd(ctx context.Context, req *v1.UpdReq) (res *v1.UpdRes, err error) {
	// 判断该分类是否存在
	if ok := article_grp.IsExist(ctx, req.GrpId); !ok {
		err = utility.Err.Skip(10101)
		return
	}
	err = article.Upd(ctx, req.Id, &model.ArticleInput{
		GrpId:       req.GrpId,
		Title:       req.Title,
		Author:      req.Author,
		Thumb:       req.Thumb,
		Tags:        req.Tags,
		Description: req.Description,
		Content:     req.Content,
		Order:       req.Order,
		Ontop:       req.Ontop,
		Onshow:      req.Onshow,
		Hist:        req.Hist,
		Post:        req.Post,
	})
	return
}
