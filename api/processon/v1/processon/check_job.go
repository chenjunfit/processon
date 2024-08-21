package processon

import (
	"github.com/gogf/gf/v2/frame/g"
	"processon/api/processon/v1/common"
	"processon/internal/app/server/model/entity"
)

// CheckJobSearchReq 搜索,查询请求参数
type CheckJobSearchReq struct {
	g.Meta   `path:"/job/list" tags:"任务信息" method:"get" summary:"任务列表"`
	KeyWords string `p:"keyWords"`
	common.PageReq
}

type CheckJobSearchRes struct {
	g.Meta       `mime:"application/json"`
	CheckJobList []*entity.CheckJob `json:"checkJobList"`
	common.ListRes
}

// CheckJobOperateRes 添加,删除，修改任务接口返回体
type CheckJobOperateRes struct {
	g.Meta `mime:"application/json"`
}

// CheckJobAddReq 添加任务接口
type CheckJobAddReq struct {
	g.Meta          `path:"/job/add" tags:"任务信息" method:"post" summary:"添加任务"`
	CheckJobName    string `v:"required" json:"checkJobName"        dc:"任务名称"`
	ScriptId        int    `v:"required" json:"scriptId"            dc:"脚本名称"`
	ScriptName      string `v:"required" json:"scriptId"            dc:"脚本名称"`
	ClusterName     string `v:"required" json:"clusterName"           dc:"集群名称"`
	BaselineId      int    `v:"required" json:"baselineId"         dc:"基线名称"`
	BaselineName    string `v:"required" json:"baselineId"         dc:"基线名称"`
	ServiceTreePath string `json:"serviceTreePath" dc:"公司对接的服务树接口"`
	IpJson          string `v:"required" json:"ipJson"                  dc:"机器列表"`
	Creator         string `v:"required" json:"creator"                  dc:"创建者"`
}

// CheckJobEditReq 编辑任务接口
type CheckJobEditReq struct {
	g.Meta `path:"/job/edit" tags:"任务信息" method:"post" summary:"编辑任务"`
	Id     int `v:"required" json:"id"        dc:"任务Id"`
	CheckJobAddReq
}

// CheckJobDeleteReq 删除任务接口
type CheckJobDeleteReq struct {
	g.Meta `path:"/job/del" tags:"任务信息" method:"post" summary:"删除任务"`
	Id     int `v:"required" json:"id"        dc:"任务Id"`
}
