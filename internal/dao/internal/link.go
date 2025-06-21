// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LinkDao is the data access object for the table link.
type LinkDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  LinkColumns        // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// LinkColumns defines and stores column names for the table link.
type LinkColumns struct {
	Id          string //
	Name        string // 名称
	Description string // 描述
	Link        string // 链接
}

// linkColumns holds the columns for the table link.
var linkColumns = LinkColumns{
	Id:          "id",
	Name:        "name",
	Description: "description",
	Link:        "link",
}

// NewLinkDao creates and returns a new DAO object for table data access.
func NewLinkDao(handlers ...gdb.ModelHandler) *LinkDao {
	return &LinkDao{
		group:    "default",
		table:    "link",
		columns:  linkColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *LinkDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *LinkDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *LinkDao) Columns() LinkColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *LinkDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *LinkDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *LinkDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
