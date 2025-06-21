// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SentenceDao is the data access object for the table sentence.
type SentenceDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SentenceColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SentenceColumns defines and stores column names for the table sentence.
type SentenceColumns struct {
	Id       string //
	BookId   string //
	Sentence string //
}

// sentenceColumns holds the columns for the table sentence.
var sentenceColumns = SentenceColumns{
	Id:       "id",
	BookId:   "book_id",
	Sentence: "sentence",
}

// NewSentenceDao creates and returns a new DAO object for table data access.
func NewSentenceDao(handlers ...gdb.ModelHandler) *SentenceDao {
	return &SentenceDao{
		group:    "default",
		table:    "sentence",
		columns:  sentenceColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SentenceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SentenceDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SentenceDao) Columns() SentenceColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SentenceDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SentenceDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SentenceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
