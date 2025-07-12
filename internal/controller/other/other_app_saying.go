package other

import (
	"blog/api/other/app"
	"blog/internal/logic/sentence"
	"context"
)

func (c *ControllerApp) Saying(ctx context.Context, req *app.SayingReq) (res *app.SayingRes, err error) {
	saying, err := sentence.Saying(ctx)
	if err != nil {
		return nil, err
	}
	return &app.SayingRes{
		Saying: saying.Sentence,
	}, nil
}
