package article_grp

import (
	"blog/internal/logic/article_grp"
	"context"

	"blog/api/article_grp/v1"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (*v1.ListRes, error) {
	list, err := article_grp.List(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.ListRes{
		List:  list,
		Total: uint(len(list)),
	}, nil
}
