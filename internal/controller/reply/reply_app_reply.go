package reply

import (
	"blog/api/reply/app"
	"blog/internal/logic/reply"
	"blog/internal/model"
	"blog/internal/model/entity"
	"context"
)

func (c *ControllerApp) Reply(ctx context.Context, req *app.ReplyReq) (res *app.ReplyRes, err error) {
	_, err = reply.Cre(ctx, &entity.Reply{
		Aid:     int(req.Aid),
		Pid:     int(req.Pid),
		Email:   req.Email,
		Site:    req.Site,
		Name:    req.Name,
		Content: req.Content,
		// 直接审核通过
		Status: model.SuccessStatus,
	})
	return
}
