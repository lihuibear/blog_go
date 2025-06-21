// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ReplyDao is the data access object for the table reply.
type ReplyDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ReplyColumns       // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ReplyColumns defines and stores column names for the table reply.
type ReplyColumns struct {
	Id        string //
	Aid       string // 文章id
	Rid       string // 根回复id，一个根可视为一个楼层
	Pid       string // 回复的id
	PName     string // 回复的名称
	Email     string // 回复人邮箱
	Name      string // 回复人名称
	Site      string // 回复人网站
	Content   string // 回复内容
	Status    string // 状态： 1待审核 2审核通过 3审核失败
	Reason    string // 审核失败原因
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// replyColumns holds the columns for the table reply.
var replyColumns = ReplyColumns{
	Id:        "id",
	Aid:       "aid",
	Rid:       "rid",
	Pid:       "pid",
	PName:     "p_name",
	Email:     "email",
	Name:      "name",
	Site:      "site",
	Content:   "content",
	Status:    "status",
	Reason:    "reason",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewReplyDao creates and returns a new DAO object for table data access.
func NewReplyDao(handlers ...gdb.ModelHandler) *ReplyDao {
	return &ReplyDao{
		group:    "default",
		table:    "reply",
		columns:  replyColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ReplyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ReplyDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ReplyDao) Columns() ReplyColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ReplyDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ReplyDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ReplyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
