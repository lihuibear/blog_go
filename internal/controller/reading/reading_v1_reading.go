package reading

import (
	v1 "blog/api/reading/v1"
	"blog/internal/logic/reading"
	"context"
)

func (c *ControllerV1) Reading(ctx context.Context, req *v1.ReadingReq) (res *v1.ReadingRes, err error) {
	list, err := reading.List(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.ReadingRes{
		List: list,
	}, nil
}
