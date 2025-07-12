// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package reading

import (
	v1 "blog/api/reading/v1"
	"context"
)

type IReadingV1 interface {
	Reading(ctx context.Context, req *v1.ReadingReq) (res *v1.ReadingRes, err error)
}
