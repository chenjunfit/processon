// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CheckScript is the golang structure of table check_script for DAO operations like Where/Data.
type CheckScript struct {
	g.Meta        `orm:"table:check_script, do:true"`
	Id            interface{} //
	ScriptName    interface{} // 脚本名称
	ScriptContent interface{} // 脚本内容
	Creator       interface{} // 创建者
	CreateAt      *gtime.Time //
	DeleteAt      *gtime.Time //
	UpdateAt      *gtime.Time //
}
