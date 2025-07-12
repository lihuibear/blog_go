package v1

import "blog/internal/model"

type List struct {
	Id       model.Id `json:"id"`
	BookId   model.Id `json:"book_id"`
	Sentence string   `json:"sentence"`
}
