package tag_grp

import (
	"blog/api/tag_grp/v1"
	"blog/internal/logic/tag_grp"
	"blog/internal/model"
	"context"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	data, err := tag_grp.List(ctx)
	if err != nil {
		return nil, err
	}

	var list []v1.List
	for _, v := range data {
		list = append(list, v1.List{
			Id:   model.Id(v.Id),
			Name: v.Name,
		})
	}

	return &v1.ListRes{
		List: list,
	}, err
}
