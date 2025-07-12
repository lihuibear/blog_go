package sentence

import (
	"blog/internal/dao"
	"blog/internal/model"
	"blog/internal/model/do"
	"blog/internal/model/entity"
	"blog/internal/utility"
	"blog/utility/uinit"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// Cre 创建句子
func Cre(ctx context.Context, in *model.SentenceInput) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		res, err := dao.Sentence.Ctx(ctx).Data(do.Sentence{
			BookId:   in.BookId,
			Sentence: in.Sentence,
		}).Insert()
		if err != nil {
			return nil
		}
		lastId, err := res.LastInsertId()
		if err != nil {
			return nil
		}

		// 插入标签关联表数据
		var sentenceTags []do.SentenceTag
		for _, tagId := range in.TagIds {
			if tagId == 0 {
				continue
			}
			sentenceTags = append(sentenceTags, do.SentenceTag{
				SId: lastId,
				TId: tagId,
			})
		}

		if len(sentenceTags) == 0 {
			return nil
		}

		_, err = dao.SentenceTag.Ctx(ctx).Data(sentenceTags).Insert()
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return utility.Err.Sys(err)
	}
	return err
}

// Upd 更新句子
func Upd(ctx context.Context, in *model.SentenceInput) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.Sentence.Ctx(ctx).Data(do.Sentence{
			BookId:   in.BookId,
			Sentence: in.Sentence,
		}).Where("id", in.Id).Update()

		// 删除标签关联表数据
		_, err = dao.SentenceTag.Ctx(ctx).Where("s_id", in.Id).Delete()
		if err != nil {
			return nil
		}

		// 插入标签关联表数据
		var sentenceTags []do.SentenceTag
		for _, tagId := range in.TagIds {
			if tagId == 0 {
				continue
			}
			sentenceTags = append(sentenceTags, do.SentenceTag{
				SId: in.Id,
				TId: tagId,
			})
		}

		if len(sentenceTags) == 0 {
			return nil
		}

		_, err = dao.SentenceTag.Ctx(ctx).Data(sentenceTags).Insert()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return utility.Err.Sys(err)
	}

	return
}

// List 读取句子列表
func List(ctx context.Context, query *model.SentenceQuery) (list []entity.Sentence, total uint, err error) {
	if query == nil {
		query = &model.SentenceQuery{}
	}
	// 对于查询初始值的处理
	if query.Page == 0 {
		query.Page = 1
	}
	if query.Size == 0 {
		query.Size = 15
	}

	db := dao.Sentence.Ctx(ctx)
	if query.BookId > 0 {
		db = db.Where("book_id", query.BookId)
	}
	if len(query.Ids) > 0 {
		db = db.WhereIn("id", query.Ids)
	}
	if query.Search != "" {
		db = db.Where("sentence like ?", "%"+query.Search+"%")
	}
	data, totalInt, err := db.Page(query.Page, query.Size).AllAndCount(true)

	if err != nil {
		err = utility.Err.Sys(err)
		return
	}
	list = []entity.Sentence{}
	_ = data.Structs(&list)
	total = uint(totalInt)

	return
}

// Show 读取句子详情
func Show(ctx context.Context, id model.Id) (info *entity.Sentence, err error) {
	info = &entity.Sentence{}
	err = dao.Sentence.Ctx(ctx).Where("id", id).Scan(&info)
	if err != nil {
		err = utility.Err.Skip(10100)
	}
	return
}

// ShowTag 读取句子标签
func ShowTag(ctx context.Context, id model.Id) (list []entity.Tag, err error) {
	var st []entity.SentenceTag
	err = dao.SentenceTag.Ctx(ctx).Fields("t_id").Where("s_id", id).Scan(&st)
	if err != nil {
		return nil, utility.Err.Sys(err)
	}

	var tagIds []model.Id
	for _, v := range st {
		tagIds = append(tagIds, model.Id(v.TId))
	}

	list = []entity.Tag{}
	err = dao.Tag.Ctx(ctx).WhereIn("id", tagIds).Scan(&list)
	if err != nil {
		return nil, utility.Err.Sys(err)
	}

	return
}

// Del 删除句子
func Del(ctx context.Context, id model.Id) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.Sentence.Ctx(ctx).Where("id", id).Delete()
		if err != nil {
			return err
		}

		_, err = dao.SentenceTag.Ctx(ctx).Where("s_id", id).Delete()
		if err != nil {
			return nil
		}

		return nil
	})
	if err != nil {
		err = utility.Err.Sys(err)
	}
	return
}

// GetIdsByTagIds 根据标签id列表获取句子id列表
// num 获取的数量
func GetIdsByTagIds(ctx context.Context, tagIds []model.Id, p model.Paging) (ids []model.Id, total uint, err error) {
	// 默认赋值0，查询不到数据
	ids = []model.Id{0}
	data, totalInt, err := dao.SentenceTag.Ctx(ctx).
		Fields("s_id").
		WhereIn("t_id", tagIds).
		Having("count(DISTINCT t_id) = ?", len(tagIds)).
		Group("s_id").
		Page(p.Page, p.Size).AllAndCount(true)

	if err != nil {
		err = utility.Err.Sys(err)
		return
	}

	total = uint(totalInt)
	for _, v := range data {
		ids = append(ids, model.Id(v["s_id"].Int()))
	}

	return
}

// Saying 随机读取一句话，用于首页展示
func Saying(ctx context.Context) (info *entity.Sentence, err error) {
	db := dao.SentenceTag.Ctx(ctx).
		Fields("s_id").
		OrderRandom()
	if uinit.SayingTagId > 0 {
		db = db.Where("t_id", uinit.SayingTagId)
	}

	idData, err := db.Limit(1).One()
	if err != nil {
		err = utility.Err.Sys(err)
	}
	id := model.Id(idData["s_id"].Int())
	info, err = Show(ctx, id)

	return
}
