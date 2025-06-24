package article

import (
	"blog/internal/logic/article"
	"blog/internal/logic/article_grp"
	"blog/internal/model"
	"blog/internal/utility"
	"context"

	"blog/api/article/v1"
)

func (c *ControllerV1) Cre(ctx context.Context, req *v1.CreReq) (*v1.CreRes, error) {
	// 判断该分类是否存在
	if ok := article_grp.IsExist(ctx, req.GrpId); !ok {
		return nil, utility.Err.Skip(10101)
	}

	LastId, err := article.Cre(ctx, &model.ArticleInput{
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
	if err != nil {
		return nil, err
	}

	return &v1.CreRes{LastId: LastId}, nil
}
