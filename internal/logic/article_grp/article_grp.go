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
func Del(ctx context.Context, id model.Id) (err error) {
	_, err = dao.ArticleGrp.Ctx(ctx).Where("id", id).Delete()
	// 软删除掉该分类下的文章
	data, err := dao.Article.Ctx(ctx).Fields("id").Where("grp_id", id).All()
	if err != nil {
		return utility.Err.Sys(err)
	}

	for _, v := range data {
		_ = article.Del(ctx, model.Id(v["id"].Uint()), false)
	}
	return
}

// List 读取文章分类列表
func List(ctx context.Context) (list []entity.ArticleGrp, err error) {
	res, err := dao.ArticleGrp.Ctx(ctx).Order("order desc").All()
	_ = res.Structs(&list)
	return
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
