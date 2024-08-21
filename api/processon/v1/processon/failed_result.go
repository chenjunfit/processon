package processon

import (
	"github.com/gogf/gf/v2/frame/g"
	"processon/api/processon/v1/common"
	"processon/internal/app/server/model/entity"
)

// FailedResultSearchReq 搜索,查询请求参数
type FailedResultSearchReq struct {
	g.Meta   `path:"/failed/list" tags:"错误信息" method:"get" summary:"错误列表"`
	KeyWords string `p:"keyWords"`
	common.PageReq
}

type FailedResultSearchRes struct {
	g.Meta               `mime:"application/json"`
	FailedNodeResultList []*entity.FailedNodeResult `json:"failedNodeResultList"`
	common.ListRes
}

// FailedResultAddReq
type FailedResultAddReq struct {
	g.Meta     `path:"/failed/report" tags:"错误信息" method:"post" summary:"添加执行错误结果"`
	JobId      int    `v:"required" json:"jobId"      dc:"任务id"`
	NodeIp     string `v:"required" json:"nodeIp"     dc:"节点ip"`
	ResultJson string `v:"required" json:"resultJson" orm:"result_json" description:"执行结果json"`
	ErrMsg     string `json:"errMsg"     orm:"err_msg"     description:"错误信息"`
	IsSuccess  bool   `v:"required" json:"isSuccess"`
}
type FailedResultOperateRes struct {
	g.Meta `mime:"application/json"`
}
