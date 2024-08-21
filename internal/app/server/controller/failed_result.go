package controller

import (
	"context"
	"processon/api/processon/v1/processon"
	"processon/internal/app/server/model/entity"
	"processon/internal/app/server/service"
)

var FailedResult = fNew()

type sFailedResultController struct {
}

func fNew() *sFailedResultController {
	return &sFailedResultController{}
}

func (s *sFailedResultController) List(ctx context.Context, req *processon.FailedResultSearchReq) (res *processon.FailedResultSearchRes, err error) {
	var (
		total int
		list  []*entity.FailedNodeResult
	)
	res = new(processon.FailedResultSearchRes)
	total, list, err = service.FailedResult().List(ctx, req)
	res.FailedNodeResultList = list
	res.Total = total
	return
}
func (s *sFailedResultController) Add(ctx context.Context, req *processon.FailedResultAddReq) (res *processon.FailedResultOperateRes, err error) {
	res = new(processon.FailedResultOperateRes)
	err = service.FailedResult().Add(ctx, req)
	return
}
