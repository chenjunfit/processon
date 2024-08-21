package processon

import (
	"github.com/gogf/gf/v2/frame/g"
	"processon/api/processon/v1/common"
	"processon/internal/app/server/model/entity"
)

// CheckScriptSearchReq 搜索,查询请求参数
type CheckScriptSearchReq struct {
	g.Meta   `path:"/script/list" tags:"脚本信息" method:"get" summary:"脚本列表"`
	KeyWords string `p:"keyWords"`
	common.PageReq
}

type CheckScriptSearchRes struct {
	g.Meta          `mime:"application/json"`
	CheckScriptList []*entity.CheckScript `json:"checkScriptList"`
	common.ListRes
}

// CheckScriptOperateRes 添加,删除，修改脚本接口返回体
type CheckScriptOperateRes struct {
	g.Meta `mime:"application/json"`
}

// CheckScriptAddReq 添加脚本接口
type CheckScriptAddReq struct {
	g.Meta        `path:"/script/add" tags:"脚本信息" method:"post" summary:"添加脚本"`
	ScriptName    string `v:"required" json:"scriptName"       dc:"脚本名称"`
	ScriptContent string `v:"required" json:"scriptContent"  dc:"脚本内容"`
	Creator       string `v:"required" json:"creator"         dc:"创建者"`
}

// CheckScriptEditReq 编辑脚本接口
type CheckScriptEditReq struct {
	g.Meta        `path:"/script/edit" tags:"脚本信息" method:"post" summary:"编辑脚本"`
	Id            int    `v:"required" json:"id"        dc:"脚本Id"`
	ScriptName    string `v:"required" json:"scriptName"       dc:"脚本名称"`
	ScriptContent string `v:"required" json:"scriptContent"  dc:"脚本内容"`
	Creator       string `v:"required" json:"creator"         dc:"创建者"`
}

// CheckScriptDeleteReq 删除脚本接口
type CheckScriptDeleteReq struct {
	g.Meta `path:"/script/del" tags:"脚本信息" method:"post" summary:"删除脚本"`
	Id     int `v:"required" json:"id"        dc:"脚本Id"`
}

type CheckScriptDowloadReq struct {
	g.Meta     `path:"/script/download" tags:"脚本信息" method:"get" summary:"下载脚本"`
	Id         int    `v:"required" json:"id"        dc:"脚本Id"`
	ScriptName string `v:"required" json:"scriptName"       dc:"脚本名称"`
}
type CheckScriptDowloadRes struct {
}
