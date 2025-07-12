package file

import (
	v1 "blog/api/file/v1"
	"blog/internal/logic/file"
	"context"
)

func (c *ControllerV1) Upload(ctx context.Context, req *v1.UploadReq) (*v1.UploadRes, error) {
	info, err := file.Upload(ctx, req.File)
	if err != nil {
		return nil, err
	}
	return &v1.UploadRes{
		FileInfo: *info,
	}, nil
}
