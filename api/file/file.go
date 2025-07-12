// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package file

import (
	v1 "blog/api/file/v1"
	"context"
)

type IFileV1 interface {
	Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error)
}
