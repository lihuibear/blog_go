package article

import (
	"blog/internal/logic/article"
	"context"

	"blog/api/article/v1"
)

func (c *ControllerV1) Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error) {
	err = article.Del(ctx, req.Id, req.IsReal)
	return
}
