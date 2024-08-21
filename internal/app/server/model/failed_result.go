package model

type FailedNodeResult struct {
	JobId      int    `json:"jobId"      orm:"job_id"      description:"任务id"`
	NodeIp     string `json:"nodeIp"     orm:"node_ip"     description:"节点ip"`
	ResultJson string `json:"resultJson" orm:"result_json" description:"执行结果json"`
	ErrMsg     string `json:"errMsg"     orm:"err_msg"     description:"错误信息"`
	IsSuccess  bool   `json:"isSuccess" orm:"_"`
}
