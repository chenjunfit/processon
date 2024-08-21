// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FailedNodeResult is the golang structure of table failed_node_result for DAO operations like Where/Data.
type FailedNodeResult struct {
	g.Meta     `orm:"table:failed_node_result, do:true"`
	Id         interface{} //
	JobId      interface{} // 任务id
	NodeIp     interface{} // 节点ip
	ResultJson interface{} // 执行结果json
	ErrMsg     interface{} // 错误信息
	Creator    interface{} // 创建者
	CreateAt   *gtime.Time //
	UpdateAt   *gtime.Time //
	DeleteAt   *gtime.Time //
}
