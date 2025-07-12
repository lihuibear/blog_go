package sentence

import (
	v1 "blog/api/sentence/v1"
	"blog/internal/logic/sentence"
	"blog/internal/model"
	"context"
)

func (c *ControllerV1) Cre(ctx context.Context, req *v1.CreReq) (res *v1.CreRes, err error) {
	err = sentence.Cre(ctx, &model.SentenceInput{
		BookId:   req.BookId,
		TagIds:   req.TagIds,
		Sentence: req.Sentence,
	})
	return
}
