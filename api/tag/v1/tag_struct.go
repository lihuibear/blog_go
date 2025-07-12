package v1

import "blog/internal/model"

type List struct {
	Id   model.Id `json:"id"`
	Name string   `json:"name"`
}
