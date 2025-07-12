package reply

import (
	v1 "blog/api/reply/v1"
	"blog/internal/logic/reply"
	"context"
)

func (c *ControllerV1) ReplyDel(ctx context.Context, req *v1.ReplyDelReq) (res *v1.ReplyDelRes, err error) {
	err = reply.Del(ctx, req.Id)
	return
}
