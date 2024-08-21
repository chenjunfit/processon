// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DesiredResultDao is the data access object for table desired_result.
type DesiredResultDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns DesiredResultColumns // columns contains all the column names of Table for convenient usage.
}

// DesiredResultColumns defines and stores column names for table desired_result.
type DesiredResultColumns struct {
	Id           string //
	BaselineName string // 基线名称
	BaselineJson string // 基线内容
	Creator      string // 创建者
	CreateAt     string //
	DeleteAt     string //
	UpdateAt     string //
}

// desiredResultColumns holds the columns for table desired_result.
var desiredResultColumns = DesiredResultColumns{
	Id:           "id",
	BaselineName: "baseline_name",
	BaselineJson: "baseline_json",
	Creator:      "creator",
	CreateAt:     "create_at",
	DeleteAt:     "delete_at",
	UpdateAt:     "update_at",
}

// NewDesiredResultDao creates and returns a new DAO object for table data access.
func NewDesiredResultDao() *DesiredResultDao {
	return &DesiredResultDao{
		group:   "default",
		table:   "desired_result",
		columns: desiredResultColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DesiredResultDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DesiredResultDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DesiredResultDao) Columns() DesiredResultColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DesiredResultDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DesiredResultDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DesiredResultDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
