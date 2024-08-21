package controller

import (
	"context"
	"processon/api/processon/v1/processon"
	"processon/internal/app/server/model/entity"
	"processon/internal/app/server/service"
)

var CheckJob = dNew()

type sCheckJobController struct {
}

func dNew() *sCheckJobController {
	return &sCheckJobController{}
}

func (s *sCheckJobController) Add(ctx context.Context, req *processon.CheckJobAddReq) (res *processon.CheckJobOperateRes, err error) {
	res = new(processon.CheckJobOperateRes)
	err = service.CheckJob().Add(ctx, req)
	return
}

func (s *sCheckJobController) Del(ctx context.Context, req *processon.CheckJobDeleteReq) (res *processon.CheckJobOperateRes, err error) {
	res = new(processon.CheckJobOperateRes)
	err = service.CheckJob().Del(ctx, req)
	return
}

func (s *sCheckJobController) Update(ctx context.Context, req *processon.CheckJobEditReq) (res *processon.CheckJobOperateRes, err error) {
	res = new(processon.CheckJobOperateRes)
	err = service.CheckJob().Update(ctx, req)
	return
}
func (s *sCheckJobController) List(ctx context.Context, req *processon.CheckJobSearchReq) (res *processon.CheckJobSearchRes, err error) {
	var (
		total int
		list  []*entity.CheckJob
	)
	res = new(processon.CheckJobSearchRes)
	total, list, err = service.CheckJob().List(ctx, req)
	res.CheckJobList = list
	res.Total = total
	return
}
