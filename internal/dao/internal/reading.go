// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ReadingDao is the data access object for the table reading.
type ReadingDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ReadingColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ReadingColumns defines and stores column names for the table reading.
type ReadingColumns struct {
	Id         string //
	Name       string // 书名
	Author     string // 作者
	Status     string // 状态: 1弃读 2完结 9在读
	FinishedAt string // 读完时间
}

// readingColumns holds the columns for the table reading.
var readingColumns = ReadingColumns{
	Id:         "id",
	Name:       "name",
	Author:     "author",
	Status:     "status",
	FinishedAt: "finished_at",
}

// NewReadingDao creates and returns a new DAO object for table data access.
func NewReadingDao(handlers ...gdb.ModelHandler) *ReadingDao {
	return &ReadingDao{
		group:    "default",
		table:    "reading",
		columns:  readingColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ReadingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ReadingDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ReadingDao) Columns() ReadingColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ReadingDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ReadingDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ReadingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
