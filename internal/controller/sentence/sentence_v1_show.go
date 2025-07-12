package sentence

import (
	v1 "blog/api/sentence/v1"
	"blog/internal/logic/sentence"
	"context"
)

func (c *ControllerV1) Show(ctx context.Context, req *v1.ShowReq) (res *v1.ShowRes, err error) {
	data, err := sentence.Show(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	tagList, err := sentence.ShowTag(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.ShowRes{
		Sentence: data,
		TagList:  tagList,
	}, nil
}
