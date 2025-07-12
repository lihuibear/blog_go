package article_grp

import (
	"blog/api/article_grp/app"
	"blog/internal/logic/article_grp"
	"context"
)

func (c *ControllerApp) List(ctx context.Context, req *app.ListReq) (res *app.ListRes, err error) {
	list, err := article_grp.List(ctx)
	if err != nil {
		return nil, err
	}

	var listOut []app.List
	countList, _ := article_grp.ListArticleCount(ctx)

	for _, v := range list {
		listOut = append(listOut, app.List{
			Id:           v.Id,
			Name:         v.Name,
			Tags:         v.Tags,
			Description:  v.Description,
			ArticleCount: countList[v.Id],
		})
	}
	res = &app.ListRes{
		List: listOut,
	}
	return
}
