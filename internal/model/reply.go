package model

import (
	"blog/internal/model/entity"
	"github.com/gogf/gf/v2/os/gtime"
)

const (
	AllStatus = iota
	CheckStatus
	SuccessStatus
	FailStatus
)

// ReplyInput 新增回复，多了Status
type ReplyInput struct {
	ReplyInputApp
	Status ReplyStatus `json:"status" v:"in:1,2,3" dc:"审核状态，1待审核 2审核通过 3审核失败，默认2"`
}

// ReplyInputApp App新增回复
type ReplyInputApp struct {
	ReplyBody
	Aid Id `json:"aid" v:"required|integer|between:1,4294967295"`
	Pid Id `json:"pid" v:"required|integer|between:0,4294967295" dc:"如果传入pid，aid会自动跟随父回复"`
}

// ReplyBody 回复内容主体
type ReplyBody struct {
	Name    string `json:"name" v:"required|length:1,20"`
	Email   string `json:"email" v:"required|email|length:1,100"`
	Site    string `json:"site" v:"url|length:1,50"`
	Content string `json:"content" v:"required|length:2, 100000"`
}

type ReplyQuery struct {
	Paging
	Aid    Id          `json:"aid" v:"integer|between:1,4294967295"`
	Status ReplyStatus `json:"status" v:"in:1,2,3" dc:"查询状态，1待审核 2审核通过 3审核失败	，不传则查询所有"`
	Search string      `v:"length: 1,30" json:"search" dc:"查询文本，会检索名称，邮箱，内容"`
}

type ReplyShow struct {
	ArticleTitle string       `json:"articleTitle" dc:"回复的文章标题"`
	ParentReply  entity.Reply `json:"parentReply" dc:"父级回复"`
	entity.Reply
}

// ReplyFloorApp 回复楼，是一个二维结构
type ReplyFloorApp struct {
	Id        Id              `json:"id"        description:""`
	Pid       Id              `json:"pid"       description:"回复父id"`
	PName     string          `json:"pName"     description:"回复父名称"`
	Email     string          `json:"email"     description:"回复人邮箱"`
	Name      string          `json:"name"      description:"回复人名称"`
	Site      string          `json:"site"      description:"回复人网站"`
	Content   string          `json:"content"   description:"回复内容"`
	CreatedAt *gtime.Time     `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time     `json:"updatedAt" description:"更新时间"`
	List      []ReplyFloorApp `json:"list"      description:"子回复列表"`
}
