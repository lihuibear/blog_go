package article

import (
	"blog/internal/logic/article"
	"context"

	"blog/api/article/v1"
)

func (c *ControllerV1) ReCre(ctx context.Context, req *v1.ReCreReq) (res *v1.ReCreRes, err error) {
	err = article.ReCre(ctx, req.Id)
	return
}
