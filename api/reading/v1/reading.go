package v1

import (
	"blog/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	ReadingReq struct {
		g.Meta `path:"reading" method:"get" sm:"查询阅读列表" tags:"reading"`
	}

	ReadingRes struct {
		List []entity.Reading `json:"list"`
	}
)
