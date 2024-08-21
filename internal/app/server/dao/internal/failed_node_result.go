// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FailedNodeResultDao is the data access object for table failed_node_result.
type FailedNodeResultDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns FailedNodeResultColumns // columns contains all the column names of Table for convenient usage.
}

// FailedNodeResultColumns defines and stores column names for table failed_node_result.
type FailedNodeResultColumns struct {
	Id         string //
	JobId      string // 任务id
	NodeIp     string // 节点ip
	ResultJson string // 执行结果json
	ErrMsg     string // 错误信息
	Creator    string // 创建者
	CreateAt   string //
	UpdateAt   string //
	DeleteAt   string //
}

// failedNodeResultColumns holds the columns for table failed_node_result.
var failedNodeResultColumns = FailedNodeResultColumns{
	Id:         "id",
	JobId:      "job_id",
	NodeIp:     "node_ip",
	ResultJson: "result_json",
	ErrMsg:     "err_msg",
	Creator:    "creator",
	CreateAt:   "create_at",
	UpdateAt:   "update_at",
	DeleteAt:   "delete_at",
}

// NewFailedNodeResultDao creates and returns a new DAO object for table data access.
func NewFailedNodeResultDao() *FailedNodeResultDao {
	return &FailedNodeResultDao{
		group:   "default",
		table:   "failed_node_result",
		columns: failedNodeResultColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FailedNodeResultDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FailedNodeResultDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FailedNodeResultDao) Columns() FailedNodeResultColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FailedNodeResultDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FailedNodeResultDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FailedNodeResultDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
