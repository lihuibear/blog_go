package reply

import (
	v1 "blog/api/reply/v1"
	"blog/internal/logic/reply"
	"context"
)

func (c *ControllerV1) ReplyShow(ctx context.Context, req *v1.ReplyShowReq) (*v1.ReplyShowRes, error) {
	info, err := reply.Show(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.ReplyShowRes{
		ReplyShow: info,
	}, nil
}
