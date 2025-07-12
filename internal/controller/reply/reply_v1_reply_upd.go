package reply

import (
	v1 "blog/api/reply/v1"
	"blog/internal/logic/reply"
	"context"
)

func (c *ControllerV1) ReplyUpd(ctx context.Context, req *v1.ReplyUpdReq) (res *v1.ReplyUpdRes, err error) {
	err = reply.Upd(ctx, req.Id, req.ReplyBody)
	return
}
