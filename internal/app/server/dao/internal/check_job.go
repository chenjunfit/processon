// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CheckJobDao is the data access object for table check_job.
type CheckJobDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns CheckJobColumns // columns contains all the column names of Table for convenient usage.
}

// CheckJobColumns defines and stores column names for table check_job.
type CheckJobColumns struct {
	Id              string //
	CheckJobName    string // 任务名称
	ScriptId        string // 脚本id
	ClusterName     string // 集群名称
	BaselineId      string // 基线id
	ServiceTreePath string // 公司对接的服务树接口
	IpJson          string // 机器列表
	JobHasSynced    string // 任务是否下发
	JobHasCompleted string // 任务是否完成
	AllNum          string // 总数
	SuccessNum      string // 成功总数
	FailNum         string // 失败总数
	Creator         string // 创建者
	DeleteAt        string //
	UpdateAt        string //
	CreateAt        string //
	ScriptName      string // 脚本名称
	BaselineName    string // 基线名称
}

// checkJobColumns holds the columns for table check_job.
var checkJobColumns = CheckJobColumns{
	Id:              "id",
	CheckJobName:    "check_job_name",
	ScriptId:        "script_id",
	ClusterName:     "cluster_name",
	BaselineId:      "baseline_id",
	ServiceTreePath: "service_tree_path",
	IpJson:          "ip_json",
	JobHasSynced:    "job_has_synced",
	JobHasCompleted: "job_has_completed",
	AllNum:          "all_num",
	SuccessNum:      "success_num",
	FailNum:         "fail_num",
	Creator:         "creator",
	DeleteAt:        "delete_at",
	UpdateAt:        "update_at",
	CreateAt:        "create_at",
	ScriptName:      "script_name",
	BaselineName:    "baseline_name",
}

// NewCheckJobDao creates and returns a new DAO object for table data access.
func NewCheckJobDao() *CheckJobDao {
	return &CheckJobDao{
		group:   "default",
		table:   "check_job",
		columns: checkJobColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CheckJobDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CheckJobDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CheckJobDao) Columns() CheckJobColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CheckJobDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CheckJobDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CheckJobDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
