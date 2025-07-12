package app

import (
	"blog/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type ReplyReq struct {
	g.Meta `path:"reply" method:"post" sm:"文章回复" tags:"app"`
	*model.ReplyInputApp
}

type ReplyRes struct {
}
