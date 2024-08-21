package check

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/util/gconv"
	"k8s.io/apimachinery/pkg/util/wait"
	"processon/internal/app/server/dao"
	"processon/internal/app/server/model/do"
	"processon/internal/app/server/model/entity"
	"time"
)

type CoumputeJobManager struct {
	cfg *gcfg.Config
}

func NewCoumputeJobManager() (*CoumputeJobManager, error) {
	config, err := gcfg.New()
	if err != nil {
		return nil, err
	}

	return &CoumputeJobManager{cfg: config}, nil
}
func (cj *CoumputeJobManager) RunComputeJobManager(ctx context.Context) error {
	internal, err := cj.cfg.Get(ctx, "checkjobConf.ComputeJobInternalSeconds")
	if err != nil {
		return err
	}
	go wait.UntilWithContext(ctx, cj.SpanComputeJob, time.Second*time.Duration(gconv.Int(internal)))
	<-ctx.Done()
	glog.Error(ctx, "CoumputeJobManager Context Done")
	return gerror.New("CoumputeJobManager Context Done")
}
func (cj *CoumputeJobManager) SpanComputeJob(ctx context.Context) {
	var pool *grpool.Pool
	batch, err := cj.cfg.Get(ctx, "checkjobConf.RunComputeJobBatch")
	if err != nil || gconv.Int(batch) == 0 {
		pool = grpool.New(10)
		glog.Error(ctx, "checkjobConf.RunCheckJobBatch :", err)
	} else {
		pool = grpool.New(gconv.Int(batch))
	}
	mins, err := cj.cfg.Get(ctx, "checkjobConf.JobCompleteMinutes")
	var minutes float64
	if mins == nil {
		minutes = 5
	} else {
		minutes = mins.Float64()
	}
	jobs := make([]*entity.CheckJob, 0)
	err = dao.CheckJob.Ctx(ctx).Where("job_has_completed", 0).Scan(&jobs)
	if err != nil {
		glog.Error(ctx, "Get ALl Jobs: :", err)
	}
	if len(jobs) == 0 {
		return
	}
	for i := 0; i < len(jobs); i++ {
		job := jobs[i]
		if time.Now().Sub(job.UpdateAt.Time).Minutes() < minutes {
			continue
		}
		pool.Add(ctx, func(ctx context.Context) {
			cj.CoumputeJob(ctx, job)
		})
	}

}
func (cj *CoumputeJobManager) CoumputeJob(ctx context.Context, job *entity.CheckJob) {
	allNum := job.AllNum
	failNum, err := dao.FailedNodeResult.Ctx(ctx).Where("job_id", job.Id).Count()
	if err != nil {
		glog.Error(ctx, "query job_id %d failed", job.Id)
	}
	succesNum := allNum - failNum
	_, err = dao.CheckJob.Ctx(ctx).Data(do.CheckJob{
		JobHasCompleted: 1,
		SuccessNum:      succesNum,
		FailNum:         failNum,
	}).Where("id", job.Id).Update()
	if err != nil {
		glog.Error(ctx, "update job_id %d failed", job.Id)
	}
}
