// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FailedNodeResult is the golang structure for table failed_node_result.
type FailedNodeResult struct {
	Id         int         `json:"id"         orm:"id"          description:""`
	JobId      int         `json:"jobId"      orm:"job_id"      description:"任务id"`
	NodeIp     string      `json:"nodeIp"     orm:"node_ip"     description:"节点ip"`
	ResultJson string      `json:"resultJson" orm:"result_json" description:"执行结果json"`
	ErrMsg     string      `json:"errMsg"     orm:"err_msg"     description:"错误信息"`
	Creator    string      `json:"creator"    orm:"creator"     description:"创建者"`
	CreateAt   *gtime.Time `json:"createAt"   orm:"create_at"   description:""`
	UpdateAt   *gtime.Time `json:"updateAt"   orm:"update_at"   description:""`
	DeleteAt   *gtime.Time `json:"deleteAt"   orm:"delete_at"   description:""`
}
