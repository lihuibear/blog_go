package sentence

import (
	v1 "blog/api/sentence/v1"
	"blog/internal/logic/sentence"
	"blog/internal/model"
	"context"
)

func (c *ControllerV1) Upd(ctx context.Context, req *v1.UpdReq) (res *v1.UpdRes, err error) {
	err = sentence.Upd(ctx, &model.SentenceInput{
		Id:       req.Id,
		BookId:   req.BookId,
		Sentence: req.Sentence,
		TagIds:   req.TagIds,
	})
	return
}
