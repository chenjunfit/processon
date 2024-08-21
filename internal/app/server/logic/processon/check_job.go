package processon

import (
	"context"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
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

type sCheckJob struct {
}

func init() {
	service.RegisterCheckJob(jNew())
}
func jNew() *sCheckJob {
	return &sCheckJob{}
}

func (s *sCheckJob) Add(ctx context.Context, req *processon.CheckJobAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		id, err := dao.DesiredResult.Ctx(ctx).Fields("id").Where("id", req.BaselineId).Where("baseline_name", req.BaselineName).Value()
		if id == nil {
			err = gerror.New("baseline_id not exist")
			liberr.ErrIsNil(ctx, err)
		}
		id, err = dao.CheckScript.Ctx(ctx).Fields("id").Where("id", req.ScriptId).Where("script_name", req.ScriptName).Value()
		if id == nil {
			err = gerror.New("script_name not exist")
			liberr.ErrIsNil(ctx, err)
		}
		hosts := make([]string, 0)
		err = gjson.Unmarshal([]byte(req.IpJson), &hosts)
		if err != nil {
			err = gerror.New("unmarshal ipJson failed")
			liberr.ErrIsNil(ctx, err)
		}
		set := gset.NewStrSet()
		set.Add(hosts...)
		g.Dump(set)
		_, err = dao.CheckJob.Ctx(ctx).Data(do.CheckJob{
			CheckJobName:    req.CheckJobName,
			ScriptId:        req.ScriptId,
			ScriptName:      req.ScriptName,
			ClusterName:     req.ClusterName,
			BaselineId:      req.BaselineId,
			BaselineName:    req.BaselineName,
			ServiceTreePath: req.ServiceTreePath,
			IpJson:          set.String(),
			AllNum:          set.Size(),
			Creator:         req.Creator,
		},
		).Save()
		liberr.ErrIsNil(ctx, err, "插入失败")
	})
	return

}
func (s *sCheckJob) Update(ctx context.Context, req *processon.CheckJobEditReq) (err error) {
	var id interface{}
	err = g.Try(ctx, func(ctx context.Context) {
		id, err = dao.CheckJob.Ctx(ctx).Fields("id").Where("id", req.Id).Value()
		if id != nil && gconv.Int(id) != req.Id {
			liberr.ErrIsNil(ctx, err, "id不存在")
		}
		ips, err := gjson.LoadJson(req.IpJson)
		if err != nil {
			err = gerror.New("IpJson is not correct")
			liberr.ErrIsNil(ctx, err)
		}
		_, err = dao.CheckJob.Ctx(ctx).Data(do.CheckJob{
			CheckJobName:    req.CheckJobName,
			ScriptId:        req.ScriptId,
			ScriptName:      req.ScriptName,
			ClusterName:     req.ClusterName,
			BaselineId:      req.BaselineId,
			BaselineName:    req.BaselineName,
			ServiceTreePath: req.ServiceTreePath,
			IpJson:          req.IpJson,
			AllNum:          len(ips.Array()),
			Creator:         req.Creator,
		}).Where("id", req.Id).Update()
		liberr.ErrIsNil(ctx, err, "插入失败")
	})
	return
}
func (s *sCheckJob) Del(ctx context.Context, req *processon.CheckJobDeleteReq) (err error) {
	var id interface{}
	err = g.Try(ctx, func(ctx context.Context) {
		id, err = dao.CheckJob.Ctx(ctx).Fields("id").Where("id", req.Id).Value()
		if id != nil && gconv.Int(id) != req.Id {
			liberr.ErrIsNil(ctx, err, "id不存在")
		}
		_, err = dao.CheckJob.Ctx(ctx).Where("id", req.Id).Delete()
		liberr.ErrIsNil(ctx, err, "插入失败")
	})
	return
}
func (s *sCheckJob) List(ctx context.Context, req *processon.CheckJobSearchReq) (total int, list []*entity.CheckJob, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.CheckJob.Ctx(ctx)
		if req.KeyWords != "" {
			keyWords := "%" + req.KeyWords + "%"
			m = m.Where("check_job_name like ? or  creator like ? or cluster_name like ?", keyWords, keyWords, keyWords)
		}

		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		total, err = m.Count()
		err = m.FieldsEx(dao.CheckJob.Columns().DeleteAt).Page(req.PageNum, req.PageSize).Order("id DESC").Scan(&list)
		liberr.ErrIsNil(ctx, err, "获取任务列表失败")
	})
	return
}
