package article_grp

import (
	"blog/internal/dao"
	"blog/internal/logic/article"
	"blog/internal/model"
	"blog/internal/model/do"
	"blog/internal/model/entity"
	"blog/internal/utility"
	"context"
)

// Cre 创建文章分类
func Cre(ctx context.Context, in *model.ArticleGrpInput) (err error) {
	_, err = dao.ArticleGrp.Ctx(ctx).Data(do.ArticleGrp{
		Name:        in.Name,
		Tags:        in.Tags,
		Description: in.Description,
		Onshow:      in.Onshow,
		Order:       in.Order,
	}).Insert()
	if err != nil {
		return utility.Err.Sys(err)
	}
	return
}

// Upd 更新文章分类
func Upd(ctx context.Context, id model.Id, in *model.ArticleGrpInput) (err error) {
	_, err = dao.ArticleGrp.Ctx(ctx).Data(do.ArticleGrp{
		Name:        in.Name,
		Tags:        in.Tags,
		Description: in.Description,
		Onshow:      in.Onshow,
		Order:       in.Order,
	}).Where("id", id).Update()
	if err != nil {
		return utility.Err.Sys(err)
	}
	return
}

//// Del 删除文章分类（不考虑分类下的文章）
//func Del(ctx context.Context, id model.Id) (err error) {
//	_, err = dao.ArticleGrp.Ctx(ctx).Where("id", id).Delete()
//
//	return
//}

// Del 删除文章分类
// Del 删除文章分类及其下的所有文章
// 该函数首先删除指定ID的分类，然后软删除该分类下的所有文章
// 参数:
//
//	ctx context.Context: 上下文对象，用于传递请求范围的信息
//	id model.Id: 要删除的分类ID
//
// 返回值:
//
//	err error: 错误对象，如果执行过程中发生错误，则返回错误信息
func Del(ctx context.Context, id model.Id) (err error) {
	// 删除指定ID的分类
	_, err = dao.ArticleGrp.Ctx(ctx).Where("id", id).Delete()
	// 软删除掉该分类下的文章
	data, err := dao.Article.Ctx(ctx).Fields("id").Where("grp_id", id).All()
	if err != nil {
		return utility.Err.Sys(err)
	}

	// 遍历分类下的所有文章，并软删除它们
	for _, v := range data {
		_ = article.Del(ctx, model.Id(v["id"].Uint()), false)
	}
	return
}

// List 读取文章分类列表
// List 查询并返回所有文章组列表。
// 该函数使用上下文 context.Context 来取消查询或处理超时情况。
// 返回值包括：[]entity.ArticleGrp 类型的列表，包含所有查询到的文章组；
// 和 error 类型的 err，用于返回查询过程中可能发生的错误。
func List(ctx context.Context) (list []entity.ArticleGrp, err error) {
	// 使用 dao.ArticleGrp.Ctx(ctx) 来执行查询，并按 'order' 字段降序排列结果。
	res, err := dao.ArticleGrp.Ctx(ctx).Order("order desc").All()
	if err != nil {
		// 如果查询过程中发生错误，直接返回空列表和错误。
		return nil, err
	}

	// 将查询结果转换为 []entity.ArticleGrp 类型的切片。
	// 这里使用 _ = res.Structs(&list) 来忽略可能的错误，因为转换失败的情况在本函数中不作处理。
	_ = res.Structs(&list)

	// 返回查询到的文章组列表和可能的错误。
	return list, err
}

// ListArticleCount 获取已经发布文章的文章分类列表统计
func ListArticleCount(ctx context.Context) (map[uint]uint, error) {
	listCount, err := dao.Article.Ctx(ctx).
		Fields("count(*) count,grp_id").
		Where("onshow", 1).
		Group("grp_id").All()
	if err != nil {
		return nil, utility.Err.Sys(err)
	}
	idCountMap := make(map[uint]uint, len(listCount))
	for _, v := range listCount {
		idCountMap[v["grp_id"].Uint()] = v["count"].Uint()
	}

	return idCountMap, nil
}

// Show 读取文章分类详情
func Show(ctx context.Context, id model.Id) (info *entity.ArticleGrp, err error) {
	err = dao.ArticleGrp.Ctx(ctx).Where("id", id).Scan(&info)
	if err != nil {
		err = utility.Err.Skip(10100)
	}
	return
}

// IsExist 根据id判断一个文章分类是否存在
func IsExist(ctx context.Context, id model.Id) bool {
	num, _ := dao.ArticleGrp.Ctx(ctx).Where("id", id).Count()
	return num == 1
}
