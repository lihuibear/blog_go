package v1

import (
	"blog/internal/model"
	"blog/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type LinkCreReq struct {
	g.Meta `path:"link/create" method:"post" sm:"新增" tags:"友链"`
	*model.LinkInput
}

type LinkCreRes struct {
}

type LinkUpdReq struct {
	g.Meta `path:"link/update/{id}" method:"post" sm:"修改" tags:"友链"`
	*model.IdInput
	*model.LinkInput
}

type LinkUpdRes struct {
}

type LinkDelReq struct {
	g.Meta `path:"link/delete/{id}" method:"post" sm:"删除" tags:"友链"`
	*model.IdInput
}

type LinkDelRes struct {
}

type LinkListReq struct {
	g.Meta `path:"link/list" method:"get" sm:"查询列表" tags:"友链"`
}

type LinkListRes struct {
	List  []entity.Link `json:"list"`
	Total uint          `json:"total"`
}
