package article

import (
	"blog/api/article/app"
	"blog/internal/logic/reply"
	"context"
)

func (c *ControllerApp) Reply(ctx context.Context, req *app.ReplyReq) (res *app.ReplyRes, err error) {
	total, list, err := reply.ListForAid(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &app.ReplyRes{
		List:  list,
		Total: total,
	}, nil
}
