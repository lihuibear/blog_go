package tag

import (
	"blog/api/tag/v1"
	"blog/internal/logic/tag"
	"blog/internal/model"
	"context"
)

// List 获取列表信息
// 该方法用于根据给定的组ID获取相关的列表信息，并返回给调用者。
// 主要功能包括调用tag.List方法获取原始数据，并将这些数据转换为v1.List类型切片。
// 参数:
//
//	ctx - 上下文，用于传递请求范围的 deadline、取消信号、请求特定的值等。
//	req - 包含请求信息的结构体，此处仅使用了GrpId字段。
//
// 返回值:
//
//	res - 包含列表信息的响应结构体指针，如果获取数据失败，则为nil。
//	err - 错误信息，如果执行成功，则为nil。
func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	// 调用List方法获取原始数据，如果发生错误则直接返回。
	data, err := tag.List(ctx, req.GrpId)
	if err != nil {
		return nil, err
	}

	// 初始化列表切片，用于存储转换后的数据。
	var list []v1.List
	// 遍历原始数据，将其转换为v1.List类型，并添加到列表切片中。
	for _, v := range data {
		list = append(list, v1.List{
			Id:   model.Id(v.Id),
			Name: v.Name,
		})
	}
	// 返回包含转换后列表信息的响应结构体指针和错误信息。
	return &v1.ListRes{
		List: list,
	}, err
}
