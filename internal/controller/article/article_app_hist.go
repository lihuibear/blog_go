package article

import (
	"blog/api/article/app"
	"blog/internal/logic/article"
	"context"
)

func (c *ControllerApp) Hist(ctx context.Context, req *app.HistReq) (res *app.HistRes, err error) {
	err = article.Hist(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return
}
