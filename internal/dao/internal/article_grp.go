// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ArticleGrpDao is the data access object for the table article_grp.
type ArticleGrpDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ArticleGrpColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ArticleGrpColumns defines and stores column names for the table article_grp.
type ArticleGrpColumns struct {
	Id          string //
	Name        string // 名称
	Tags        string // 标签，依英文逗号隔开
	Description string // 简介
	Onshow      string // 是否显示
	Order       string // 排序，越大越靠前
}

// articleGrpColumns holds the columns for the table article_grp.
var articleGrpColumns = ArticleGrpColumns{
	Id:          "id",
	Name:        "name",
	Tags:        "tags",
	Description: "description",
	Onshow:      "onshow",
	Order:       "order",
}

// NewArticleGrpDao creates and returns a new DAO object for table data access.
func NewArticleGrpDao(handlers ...gdb.ModelHandler) *ArticleGrpDao {
	return &ArticleGrpDao{
		group:    "default",
		table:    "article_grp",
		columns:  articleGrpColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ArticleGrpDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ArticleGrpDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ArticleGrpDao) Columns() ArticleGrpColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ArticleGrpDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ArticleGrpDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ArticleGrpDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
