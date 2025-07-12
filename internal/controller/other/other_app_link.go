package other

import (
	"blog/api/other/app"
	"blog/internal/logic/link"
	"context"
)

func (c *ControllerApp) Link(ctx context.Context, req *app.LinkReq) (*app.LinkRes, error) {
	list, err := link.List(ctx)
	if err != nil {
		return nil, err
	}
	return &app.LinkRes{
		List:  list,
		Total: uint(len(list)),
	}, nil
}
