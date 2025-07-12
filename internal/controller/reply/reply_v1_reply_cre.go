package reply

import (
	v1 "blog/api/reply/v1"
	"blog/internal/logic/reply"
	"blog/internal/model"
	"blog/internal/model/entity"
	"context"
)

func (c *ControllerV1) ReplyCre(ctx context.Context, req *v1.ReplyCreReq) (res *v1.ReplyCreRes, err error) {
	if req.Status == 0 {
		req.Status = model.SuccessStatus
	}
	_, err = reply.Cre(ctx, &entity.Reply{
		Aid:     int(req.Aid),
		Pid:     int(req.Pid),
		Email:   req.Email,
		Site:    req.Site,
		Name:    req.Name,
		Content: req.Content,
		Status:  int(req.Status),
	})
	return
}
