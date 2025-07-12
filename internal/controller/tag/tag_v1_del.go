package tag

import (
	"blog/internal/logic/tag"
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"blog/api/tag/v1"
)

func (c *ControllerV1) Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error) {
	err = tag.Del(ctx, req.Id)
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

/*func (c *ControllerV1) Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error) {
	err = tag.Del(ctx, req.Id)
	return nil, err
}
*/
