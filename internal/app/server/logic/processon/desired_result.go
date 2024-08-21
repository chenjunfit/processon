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

type sDesiredResult struct {
}

func init() {
	service.RegisterDesiredResult(New())
}
func New() *sDesiredResult {
	return &sDesiredResult{}
}

func (s *sDesiredResult) Add(ctx context.Context, req *processon.BaseLineAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.DesiredResult.Ctx(ctx).Data(do.DesiredResult{
			BaselineName: req.BaselineName,
			BaselineJson: req.BaselineJson,
			Creator:      req.Creator,
		},
		).Unscoped().Save()
		g.Dump(err)
		liberr.ErrIsNil(ctx, err, "插入失败")
	})
	return
}
func (s *sDesiredResult) Update(ctx context.Context, req *processon.BaseLineEditReq) (err error) {
	var id interface{}
	err = g.Try(ctx, func(ctx context.Context) {
		id, err = dao.DesiredResult.Ctx(ctx).Fields("id").Where("id", req.Id).Value()
		if id != nil && gconv.Int(id) != req.Id {
			liberr.ErrIsNil(ctx, err, "id不存在")
		}
		_, err = dao.DesiredResult.Ctx(ctx).Data(do.DesiredResult{
			Id:           req.Id,
			BaselineName: req.BaselineName,
			BaselineJson: req.BaselineJson,
			Creator:      req.Creator,
		}).Where("id", req.Id).Update()
		liberr.ErrIsNil(ctx, err, "插入失败")
	})
	return
}
func (s *sDesiredResult) Del(ctx context.Context, req *processon.BaseLineDeleteReq) (err error) {
	var id interface{}
	err = g.Try(ctx, func(ctx context.Context) {
		id, err = dao.DesiredResult.Ctx(ctx).Fields("id").Value()
		if id != nil && gconv.Int(id) != req.Id {
			liberr.ErrIsNil(ctx, err, "id不存在")
		}
		_, err = dao.DesiredResult.Ctx(ctx).Where("id", req.Id).Delete()
		liberr.ErrIsNil(ctx, err, "插入失败")
	})
	return
}
func (s *sDesiredResult) List(ctx context.Context, req *processon.BaseLineSearchReq) (total int, list []*entity.DesiredResult, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.DesiredResult.Ctx(ctx)
		if req.KeyWords != "" {
			keyWords := "%" + req.KeyWords + "%"
			m = m.Where("baseline_name like ? or  creator like ?", keyWords, keyWords)
		}

		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		total, err = m.Count()
		err = m.FieldsEx(dao.DesiredResult.Columns().DeleteAt).Page(req.PageNum, req.PageSize).Order("id DESC").Scan(&list)
		liberr.ErrIsNil(ctx, err, "获取任务列表失败")
	})
	return
}
func (s *sDesiredResult) Download(ctx context.Context, req *processon.BaseLineDowloadReq) (string, error) {
	content, err := dao.DesiredResult.Ctx(ctx).Where("id", req.Id).Where("baseline_name", req.BaseLineName).Fields("baseline_json").Value()
	if err != nil {
		return "", err
	}
	return gconv.String(content), nil
}
