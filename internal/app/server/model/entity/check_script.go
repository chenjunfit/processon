// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CheckScript is the golang structure for table check_script.
type CheckScript struct {
	Id            int         `json:"id"            orm:"id"             description:""`
	ScriptName    string      `json:"scriptName"    orm:"script_name"    description:"脚本名称"`
	ScriptContent string      `json:"scriptContent" orm:"script_content" description:"脚本内容"`
	Creator       string      `json:"creator"       orm:"creator"        description:"创建者"`
	CreateAt      *gtime.Time `json:"createAt"      orm:"create_at"      description:""`
	DeleteAt      *gtime.Time `json:"deleteAt"      orm:"delete_at"      description:""`
	UpdateAt      *gtime.Time `json:"updateAt"      orm:"update_at"      description:""`
}
