package processon

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"processon/api/processon/v1/processon"
	"processon/internal/app/server/dao"
	"processon/internal/app/server/model/do"
	"processon/internal/app/server/model/entity"
	"processon/internal/app/server/service"
	"processon/internal/consts"
	liberr "processon/internal/util"
)

type sCheckScript struct {
}

func init() {
	service.RegisterCheckScript(sNew())
}
func sNew() *sCheckScript {
	return &sCheckScript{}
}

func (s *sCheckScript) Add(ctx context.Context, req *processon.CheckScriptAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.CheckScript.Ctx(ctx).Data(do.CheckScript{
			ScriptContent: req.ScriptContent,
			ScriptName:    req.ScriptName,
			Creator:       req.Creator,
		},
		).Save()
		liberr.ErrIsNil(ctx, err, "插入失败")
	})
	return
}
func (s *sCheckScript) Update(ctx context.Context, req *processon.CheckScriptEditReq) (err error) {
	var id interface{}
	err = g.Try(ctx, func(ctx context.Context) {
		id, err = dao.CheckScript.Ctx(ctx).Fields("id").Where("id", req.Id).Value()
		if id != nil && gconv.Int(id) != req.Id {
			liberr.ErrIsNil(ctx, err, "id不存在")
		}
		_, err = dao.CheckScript.Ctx(ctx).Data(do.CheckScript{
			Id:            req.Id,
			ScriptContent: req.ScriptContent,
			ScriptName:    req.ScriptName,
			Creator:       req.Creator,
		}).Where("id", req.Id).Update()
		liberr.ErrIsNil(ctx, err, "插入失败")
	})
	return
}
func (s *sCheckScript) Del(ctx context.Context, req *processon.CheckScriptDeleteReq) (err error) {
	var id interface{}
	err = g.Try(ctx, func(ctx context.Context) {
		id, err = dao.CheckScript.Ctx(ctx).Fields("id").Where("id", req.Id).Value()
		if id != nil && gconv.Int(id) != req.Id {
			liberr.ErrIsNil(ctx, err, "id不存在")
		}
		_, err = dao.CheckScript.Ctx(ctx).Where("id", req.Id).Delete()
		liberr.ErrIsNil(ctx, err, "插入失败")
	})
	return
}
func (s *sCheckScript) List(ctx context.Context, req *processon.CheckScriptSearchReq) (total int, list []*entity.CheckScript, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.CheckScript.Ctx(ctx)
		if req.KeyWords != "" {
			keyWords := "%" + req.KeyWords + "%"
			m = m.Where("scriptName like ? or  creator like ?", keyWords, keyWords)
		}

		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		total, err = m.Count()
		err = m.FieldsEx(dao.CheckScript.Columns().DeleteAt).Page(req.PageNum, req.PageSize).Order("id DESC").Scan(&list)
		liberr.ErrIsNil(ctx, err, "获取任务列表失败")
	})
	return
}

func (s *sCheckScript) Download(ctx context.Context, req *processon.CheckScriptDowloadReq) (string, error) {
	content, err := dao.CheckScript.Ctx(ctx).Where("id", req.Id).Where("script_name", req.ScriptName).Fields("script_content").Value()
	if err != nil {
		return "", err
	}
	return gconv.String(content), nil
}
