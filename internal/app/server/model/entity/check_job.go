// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CheckJob is the golang structure for table check_job.
type CheckJob struct {
	Id              int         `json:"id"              orm:"id"                description:""`
	CheckJobName    string      `json:"checkJobName"    orm:"check_job_name"    description:"任务名称"`
	ScriptId        int         `json:"scriptId"        orm:"script_id"         description:"脚本id"`
	ClusterName     string      `json:"clusterName"     orm:"cluster_name"      description:"集群名称"`
	BaselineId      int         `json:"baselineId"      orm:"baseline_id"       description:"基线id"`
	ServiceTreePath string      `json:"serviceTreePath" orm:"service_tree_path" description:"公司对接的服务树接口"`
	IpJson          string      `json:"ipJson"          orm:"ip_json"           description:"机器列表"`
	JobHasSynced    int         `json:"jobHasSynced"    orm:"job_has_synced"    description:"任务是否下发"`
	JobHasCompleted int         `json:"jobHasCompleted" orm:"job_has_completed" description:"任务是否完成"`
	AllNum          int         `json:"allNum"          orm:"all_num"           description:"总数"`
	SuccessNum      int         `json:"successNum"      orm:"success_num"       description:"成功总数"`
	FailNum         int         `json:"failNum"         orm:"fail_num"          description:"失败总数"`
	Creator         string      `json:"creator"         orm:"creator"           description:"创建者"`
	DeleteAt        *gtime.Time `json:"deleteAt"        orm:"delete_at"         description:""`
	UpdateAt        *gtime.Time `json:"updateAt"        orm:"update_at"         description:""`
	CreateAt        *gtime.Time `json:"createAt"        orm:"create_at"         description:""`
	ScriptName      string      `json:"scriptName"      orm:"script_name"       description:"脚本名称"`
	BaselineName    string      `json:"baselineName"    orm:"baseline_name"     description:"基线名称"`
}
