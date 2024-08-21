// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CheckScriptDao is the data access object for table check_script.
type CheckScriptDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns CheckScriptColumns // columns contains all the column names of Table for convenient usage.
}

// CheckScriptColumns defines and stores column names for table check_script.
type CheckScriptColumns struct {
	Id            string //
	ScriptName    string // 脚本名称
	ScriptContent string // 脚本内容
	Creator       string // 创建者
	CreateAt      string //
	DeleteAt      string //
	UpdateAt      string //
}

// checkScriptColumns holds the columns for table check_script.
var checkScriptColumns = CheckScriptColumns{
	Id:            "id",
	ScriptName:    "script_name",
	ScriptContent: "script_content",
	Creator:       "creator",
	CreateAt:      "create_at",
	DeleteAt:      "delete_at",
	UpdateAt:      "update_at",
}

// NewCheckScriptDao creates and returns a new DAO object for table data access.
func NewCheckScriptDao() *CheckScriptDao {
	return &CheckScriptDao{
		group:   "default",
		table:   "check_script",
		columns: checkScriptColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CheckScriptDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CheckScriptDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CheckScriptDao) Columns() CheckScriptColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CheckScriptDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CheckScriptDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CheckScriptDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
