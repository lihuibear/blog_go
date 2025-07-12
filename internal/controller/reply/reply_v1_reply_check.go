package reply

import (
	v1 "blog/api/reply/v1"
	"blog/internal/logic/reply"
	"context"
)

func (c *ControllerV1) ReplyCheck(ctx context.Context, req *v1.ReplyCheckReq) (res *v1.ReplyCheckRes, err error) {
	err = reply.Check(ctx, req.Id, req.Result, req.Reason)
	return
}
