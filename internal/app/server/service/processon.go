// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"processon/api/processon/v1/processon"
	"processon/internal/app/server/model/entity"
)

type (
	ICheckJob interface {
		Add(ctx context.Context, req *processon.CheckJobAddReq) (err error)
		Update(ctx context.Context, req *processon.CheckJobEditReq) (err error)
		Del(ctx context.Context, req *processon.CheckJobDeleteReq) (err error)
		List(ctx context.Context, req *processon.CheckJobSearchReq) (total int, list []*entity.CheckJob, err error)
	}
	ICheckScript interface {
		Add(ctx context.Context, req *processon.CheckScriptAddReq) (err error)
		Update(ctx context.Context, req *processon.CheckScriptEditReq) (err error)
		Del(ctx context.Context, req *processon.CheckScriptDeleteReq) (err error)
		List(ctx context.Context, req *processon.CheckScriptSearchReq) (total int, list []*entity.CheckScript, err error)
		Download(ctx context.Context, req *processon.CheckScriptDowloadReq) (string, error)
	}
	IDesiredResult interface {
		Add(ctx context.Context, req *processon.BaseLineAddReq) (err error)
		Update(ctx context.Context, req *processon.BaseLineEditReq) (err error)
		Del(ctx context.Context, req *processon.BaseLineDeleteReq) (err error)
		List(ctx context.Context, req *processon.BaseLineSearchReq) (total int, list []*entity.DesiredResult, err error)
		Download(ctx context.Context, req *processon.BaseLineDowloadReq) (string, error)
	}
	IFailedResult interface {
		List(ctx context.Context, req *processon.FailedResultSearchReq) (total int, list []*entity.FailedNodeResult, err error)
		Add(ctx context.Context, req *processon.FailedResultAddReq) (err error)
	}
)

var (
	localCheckJob      ICheckJob
	localCheckScript   ICheckScript
	localDesiredResult IDesiredResult
	localFailedResult  IFailedResult
)

func CheckJob() ICheckJob {
	if localCheckJob == nil {
		panic("implement not found for interface ICheckJob, forgot register?")
	}
	return localCheckJob
}

func RegisterCheckJob(i ICheckJob) {
	localCheckJob = i
}

func CheckScript() ICheckScript {
	if localCheckScript == nil {
		panic("implement not found for interface ICheckScript, forgot register?")
	}
	return localCheckScript
}

func RegisterCheckScript(i ICheckScript) {
	localCheckScript = i
}

func DesiredResult() IDesiredResult {
	if localDesiredResult == nil {
		panic("implement not found for interface IDesiredResult, forgot register?")
	}
	return localDesiredResult
}

func RegisterDesiredResult(i IDesiredResult) {
	localDesiredResult = i
}

func FailedResult() IFailedResult {
	if localFailedResult == nil {
		panic("implement not found for interface IFailedResult, forgot register?")
	}
	return localFailedResult
}

func RegisterFailedResult(i IFailedResult) {
	localFailedResult = i
}
