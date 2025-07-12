package article

import (
	"blog/api/article/app"
	"blog/internal/logic/article"
	"blog/internal/model"
	"context"
)

func (c *ControllerApp) List(ctx context.Context, req *app.ListReq) (res *app.ListRes, err error) {
	query := &model.ArticleQuery{
		Paging: model.Paging{
			Page: req.Page,
			Size: req.Size,
		},
		GrpId:  req.GrpId,
		Search: req.Search,
		Onshow: true,
		IsDel:  false,
	}
	list, total, err := article.List(ctx, query)
	if err != nil {
		return nil, err
	}

	var listOut []app.List
	for _, v := range list {
		listOut = append(listOut, app.List{
			Id:          v.Id,
			GrpId:       v.GrpId,
			Title:       v.Title,
			Author:      v.Author,
			Thumb:       v.Thumb,
			Tags:        v.Tags,
			Description: v.Description,
			Hist:        v.Hist,
			Post:        v.Post,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			LastedAt:    v.LastedAt,
		})
	}
	return &app.ListRes{
		List:  listOut,
		Total: total,
	}, nil
}
