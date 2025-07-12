// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package reply

import (
	"blog/api/reply/app"
	v1 "blog/api/reply/v1"
	"context"
)

type IReplyApp interface {
	Reply(ctx context.Context, req *app.ReplyReq) (res *app.ReplyRes, err error)
}

type IReplyV1 interface {
	ReplyCre(ctx context.Context, req *v1.ReplyCreReq) (res *v1.ReplyCreRes, err error)
	ReplyUpd(ctx context.Context, req *v1.ReplyUpdReq) (res *v1.ReplyUpdRes, err error)
	ReplyDel(ctx context.Context, req *v1.ReplyDelReq) (res *v1.ReplyDelRes, err error)
	ReplyList(ctx context.Context, req *v1.ReplyListReq) (res *v1.ReplyListRes, err error)
	ReplyShow(ctx context.Context, req *v1.ReplyShowReq) (res *v1.ReplyShowRes, err error)
	ReplyCheck(ctx context.Context, req *v1.ReplyCheckReq) (res *v1.ReplyCheckRes, err error)
}
