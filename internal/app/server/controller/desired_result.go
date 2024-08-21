package controller

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"net/url"
	"processon/api/processon/v1/processon"
	"processon/internal/app/server/model/entity"
	"processon/internal/app/server/service"
	liberr "processon/internal/util"
)

var BaseLine = New()

type sBaseLineController struct {
}

func New() *sBaseLineController {
	return &sBaseLineController{}
}

func (s *sBaseLineController) Add(ctx context.Context, req *processon.BaseLineAddReq) (res *processon.BaseLineOperateRes, err error) {
	res = new(processon.BaseLineOperateRes)
	err = service.DesiredResult().Add(ctx, req)
	return
}

func (s *sBaseLineController) Del(ctx context.Context, req *processon.BaseLineDeleteReq) (res *processon.BaseLineOperateRes, err error) {
	res = new(processon.BaseLineOperateRes)
	err = service.DesiredResult().Del(ctx, req)
	return
}

func (s *sBaseLineController) Update(ctx context.Context, req *processon.BaseLineEditReq) (res *processon.BaseLineOperateRes, err error) {
	res = new(processon.BaseLineOperateRes)
	err = service.DesiredResult().Update(ctx, req)
	return
}
func (s *sBaseLineController) List(ctx context.Context, req *processon.BaseLineSearchReq) (res *processon.BaseLineSearchRes, err error) {
	var (
		total int
		list  []*entity.DesiredResult
	)
	res = new(processon.BaseLineSearchRes)
	total, list, err = service.DesiredResult().List(ctx, req)
	res.BaseLineList = list
	res.Total = total
	return
}
func (s *sBaseLineController) Download(ctx context.Context, req *processon.BaseLineDowloadReq) (res *processon.BaseLineDowloadRes, err error) {
	content, err := service.DesiredResult().Download(ctx, req)
	if err != nil {
		return
	}
	fileName := req.BaseLineName + ".json"
	r := g.RequestFromCtx(ctx).Response
	r.Header().Set("Content-Type", "application/octet-stream")
	r.Header().Set("Content-Transfer-Encoding", "binary")
	r.Header().Set("Content-Disposition", fmt.Sprintf(`attachment;filename=%s`, url.QueryEscape(fileName)))
	r.ClearBuffer()
	r.Write(content)
	liberr.ErrIsNil(ctx, err, "导出失败")
	return
}
