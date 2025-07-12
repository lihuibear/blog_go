package article

import (
	"blog/api/article/app"
	"blog/internal/logic/article"
	"context"
)

func (c *ControllerApp) About(ctx context.Context, req *app.AboutReq) (*app.AboutRes, error) {
	info, err := article.Show(ctx, 1)
	if err != nil {
		return nil, err
	}
	return &app.AboutRes{
		One: &app.One{
			Id:          info.Id,
			GrpId:       info.GrpId,
			Title:       info.Title,
			Author:      info.Author,
			Thumb:       info.Thumb,
			Tags:        info.Tags,
			Description: info.Description,
			Content:     info.Content,
			Hist:        info.Hist,
			Post:        info.Post,
			CreatedAt:   info.CreatedAt,
			UpdatedAt:   info.UpdatedAt,
			LastedAt:    info.LastedAt,
		},
	}, nil
}
