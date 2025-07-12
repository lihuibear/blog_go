// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sentence

import (
	v1 "blog/api/sentence/v1"
	"context"
)

type ISentenceV1 interface {
	Cre(ctx context.Context, req *v1.CreReq) (res *v1.CreRes, err error)
	Upd(ctx context.Context, req *v1.UpdReq) (res *v1.UpdRes, err error)
	Show(ctx context.Context, req *v1.ShowReq) (res *v1.ShowRes, err error)
	Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error)
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
}
