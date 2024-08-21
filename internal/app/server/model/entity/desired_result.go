// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DesiredResult is the golang structure for table desired_result.
type DesiredResult struct {
	Id           int         `json:"id"           orm:"id"            description:""`
	BaselineName string      `json:"baselineName" orm:"baseline_name" description:"基线名称"`
	BaselineJson string      `json:"baselineJson" orm:"baseline_json" description:"基线内容"`
	Creator      string      `json:"creator"      orm:"creator"       description:"创建者"`
	CreateAt     *gtime.Time `json:"createAt"     orm:"create_at"     description:""`
	DeleteAt     *gtime.Time `json:"deleteAt"     orm:"delete_at"     description:""`
	UpdateAt     *gtime.Time `json:"updateAt"     orm:"update_at"     description:""`
}
