package processon

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"processon/api/processon/v1/processon"
	"processon/internal/app/server/dao"
	"processon/internal/app/server/model/do"
	"processon/internal/app/server/model/entity"
	"processon/internal/app/server/service"
	"processon/internal/consts"
	liberr "processon/internal/util"
)

type sFailedResult struct {
}

func init() {
	service.RegisterFailedResult(dNew())
}
func dNew() *sFailedResult {
	return &sFailedResult{}
}

func (s *sFailedResult) List(ctx context.Context, req *processon.FailedResultSearchReq) (total int, list []*entity.FailedNodeResult, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.FailedNodeResult.Ctx(ctx)
		if req.KeyWords != "" {
			keyWords := "%" + req.KeyWords + "%"
			m = m.Where("node_ip like ? or err_msg like ?", keyWords, keyWords)
		}

		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		total, err = m.Count()
		err = m.FieldsEx(dao.CheckScript.Columns().DeleteAt).Page(req.PageNum, req.PageSize).Order("id DESC").Scan(&list)
		liberr.ErrIsNil(ctx, err, "获取失败节点列表失败")
	})
	return
}
func (s *sFailedResult) Add(ctx context.Context, req *processon.FailedResultAddReq) (err error) {
	if req.IsSuccess {
		return nil
	}
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.FailedNodeResult.Ctx(ctx).Data(do.FailedNodeResult{
			JobId:      req.JobId,
			NodeIp:     req.NodeIp,
			ResultJson: req.ResultJson,
			ErrMsg:     req.ErrMsg,
		}).Save()
		liberr.ErrIsNil(ctx, err, "插入失败")
	})
	return
}
