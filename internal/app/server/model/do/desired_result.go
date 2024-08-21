// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DesiredResult is the golang structure of table desired_result for DAO operations like Where/Data.
type DesiredResult struct {
	g.Meta       `orm:"table:desired_result, do:true"`
	Id           interface{} //
	BaselineName interface{} // 基线名称
	BaselineJson interface{} // 基线内容
	Creator      interface{} // 创建者
	CreateAt     *gtime.Time //
	DeleteAt     *gtime.Time //
	UpdateAt     *gtime.Time //
}
