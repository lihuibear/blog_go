package link

import (
	v1 "blog/api/link/v1"
	"blog/internal/logic/link"
	"context"
)

func (c *ControllerV1) LinkList(ctx context.Context, req *v1.LinkListReq) (res *v1.LinkListRes, err error) {
	list, err := link.List(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.LinkListRes{
		List:  list,
		Total: uint(len(list)),
	}, nil
}
