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

var CheckScript = cNew()

type sCheckScriptController struct {
}

func cNew() *sCheckScriptController {
	return &sCheckScriptController{}
}

func (s *sCheckScriptController) Add(ctx context.Context, req *processon.CheckScriptAddReq) (res *processon.CheckScriptOperateRes, err error) {
	res = new(processon.CheckScriptOperateRes)
	err = service.CheckScript().Add(ctx, req)
	return
}

func (s *sCheckScriptController) Del(ctx context.Context, req *processon.CheckScriptDeleteReq) (res *processon.CheckScriptOperateRes, err error) {
	res = new(processon.CheckScriptOperateRes)
	err = service.CheckScript().Del(ctx, req)
	return
}

func (s *sCheckScriptController) Update(ctx context.Context, req *processon.CheckScriptEditReq) (res *processon.CheckScriptOperateRes, err error) {
	res = new(processon.CheckScriptOperateRes)
	err = service.CheckScript().Update(ctx, req)
	return
}
func (s *sCheckScriptController) List(ctx context.Context, req *processon.CheckScriptSearchReq) (res *processon.CheckScriptSearchRes, err error) {
	var (
		total int
		list  []*entity.CheckScript
	)
	res = new(processon.CheckScriptSearchRes)
	total, list, err = service.CheckScript().List(ctx, req)
	res.CheckScriptList = list
	res.Total = total
	return
}
func (s *sCheckScriptController) Download(ctx context.Context, req *processon.CheckScriptDowloadReq) (res *processon.CheckScriptDowloadRes, err error) {
	content, err := service.CheckScript().Download(ctx, req)
	fileName := req.ScriptName
	r := g.RequestFromCtx(ctx).Response
	r.Header().Set("Content-Type", "application/octet-stream")
	r.Header().Set("Content-Transfer-Encoding", "binary")
	r.Header().Set("Content-Disposition", fmt.Sprintf(`attachment;filename=%s`, url.QueryEscape(fileName)))
	r.ClearBuffer()
	r.Write(content)
	liberr.ErrIsNil(ctx, err, "导出失败")
	return
}
