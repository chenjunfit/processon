package processon

import (
	"github.com/gogf/gf/v2/frame/g"
	"processon/api/processon/v1/common"
	"processon/internal/app/server/model/entity"
)

// BaseLineSearchReq 搜索,查询请求参数
type BaseLineSearchReq struct {
	g.Meta   `path:"/baseline/list" tags:"基线信息" method:"get" summary:"基线列表"`
	KeyWords string `p:"keyWords"`
	common.PageReq
}

type BaseLineSearchRes struct {
	g.Meta       `mime:"application/json"`
	BaseLineList []*entity.DesiredResult `json:"baseLineList"`
	common.ListRes
}

// BaseLineOperateRes 添加,删除，修改基线接口返回体
type BaseLineOperateRes struct {
	g.Meta `mime:"application/json"`
}

// BaseLineAddReq 添加基线接口
type BaseLineAddReq struct {
	g.Meta       `path:"/baseline/add" tags:"基线信息" method:"post" summary:"添加基线"`
	BaselineName string `v:"required" json:"baselineName" dc:"基线名称"`
	BaselineJson string `v:"required" json:"baselineJson"  dc:"基线内容"`
	Creator      string `v:"required" json:"creator"       dc:"创建者"`
}

// BaseLineEditReq 编辑基线接口
type BaseLineEditReq struct {
	g.Meta       `path:"/baseline/edit" tags:"基线信息" method:"post" summary:"编辑基线"`
	Id           int    `v:"required" json:"id"        dc:"基线Id"`
	BaselineName string `v:"required" json:"baselineName" dc:"基线名称"`
	BaselineJson string `v:"required" json:"baselineJson"  dc:"基线内容"`
	Creator      string `v:"required" json:"creator"       dc:"创建者"`
}

// BaseLineDeleteReq 删除基线接口
type BaseLineDeleteReq struct {
	g.Meta `path:"/baseline/del" tags:"基线信息" method:"post" summary:"删除基线"`
	Id     int `v:"required" json:"id"        dc:"基线Id"`
}

type BaseLineDowloadReq struct {
	g.Meta       `path:"/baseline/dowload" tags:"基线信息" method:"get" summary:"下载基线"`
	Id           int    `v:"required" json:"id"        dc:"基线Id"`
	BaseLineName string `v:"required" json:"baselineName"       dc:"基线名称"`
}
type BaseLineDowloadRes struct {
}
