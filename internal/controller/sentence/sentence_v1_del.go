package sentence

import (
	v1 "blog/api/sentence/v1"
	"blog/internal/logic/sentence"
	"context"
)

func (c *ControllerV1) Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error) {
	err = sentence.Del(ctx, req.Id)
	return
}
