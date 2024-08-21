// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CheckJob is the golang structure of table check_job for DAO operations like Where/Data.
type CheckJob struct {
	g.Meta          `orm:"table:check_job, do:true"`
	Id              interface{} //
	CheckJobName    interface{} // 任务名称
	ScriptId        interface{} // 脚本id
	ClusterName     interface{} // 集群名称
	BaselineId      interface{} // 基线id
	ServiceTreePath interface{} // 公司对接的服务树接口
	IpJson          interface{} // 机器列表
	JobHasSynced    interface{} // 任务是否下发
	JobHasCompleted interface{} // 任务是否完成
	AllNum          interface{} // 总数
	SuccessNum      interface{} // 成功总数
	FailNum         interface{} // 失败总数
	Creator         interface{} // 创建者
	DeleteAt        *gtime.Time //
	UpdateAt        *gtime.Time //
	CreateAt        *gtime.Time //
	ScriptName      interface{} // 脚本名称
	BaselineName    interface{} // 基线名称
}
