package article_grp

import (
	"blog/internal/logic/article_grp"
	"context"

	"blog/api/article_grp/v1"
)

// 删除

func (c *ControllerV1) Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error) {
	err = article_grp.Del(ctx, req.Id)
	return
}
