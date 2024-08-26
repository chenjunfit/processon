package metric

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/apimachinery/pkg/util/wait"
	"processon/internal/app/server/dao"
	"processon/internal/app/server/model/entity"
	"time"
)

type MetricJobManager struct {
	cfg *gcfg.Config
}

var (
	jobMetrics = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "CheckJob",
	}, []string{
		"CheckJobName",
		"CreateAt",
		"Creator",
		"AllNum",
		"FailNum",
		"BaselineName",
		"ScriptName",
		"ProbabilityOfSuccess",
	},
	)
	failedMetrics = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "FailedResult",
	}, []string{
		"NodeIp",
		"ErrMsg",
		"DesiredKey",
		"DesiredValue",
		"ActualValue",
		"CheckJobName",
	},
	)
)

func Register() {
	prometheus.Register(jobMetrics)
	prometheus.Register(failedMetrics)

}
func NewMetricJobManager() (*MetricJobManager, error) {
	config, err := gcfg.New()
	if err != nil {
		return nil, err
	}

	return &MetricJobManager{cfg: config}, nil
}
func (mj *MetricJobManager) RunMetricManager(ctx context.Context) error {
	internal, err := mj.cfg.Get(ctx, "checkjobConf.MetricsJobInternalSeconds")
	if err != nil || internal == nil {
		return err
	}
	go wait.UntilWithContext(ctx, mj.SpanMetricJob, time.Second*time.Duration(gconv.Int(internal)))
	<-ctx.Done()
	glog.Error(ctx, "RunMetricManager Context Done")
	return gerror.New("RunMetricManager Context Done")
}
func (mj *MetricJobManager) SpanMetricJob(ctx context.Context) {
	//基线指标展示
	go mj.MetricBaseLine(ctx)
	//展示对比结果
	jobList := make([]*entity.CheckJob, 0)
	err := dao.CheckJob.Ctx(ctx).Where("job_has_completed = ?", 1).Scan(&jobList)
	if err != nil {
		glog.Error(ctx, "Get CheckJob list failed", err)
	}
	if len(jobList) == 0 {
		return
	}
	pool := grpool.New(len(jobList))
	for _, job := range jobList {
		_ = pool.Add(ctx, func(ctx context.Context) {
			mj.SubmitJob(ctx, job)
		})
	}
	//
}
func (mj *MetricJobManager) SubmitJob(ctx context.Context, job *entity.CheckJob) {
	mapLabel := make(map[string]string)
	mapLabel["CheckJobName"] = job.CheckJobName
	mapLabel["CreateAt"] = job.CreateAt.String()
	mapLabel["Creator"] = job.Creator
	mapLabel["AllNum"] = gconv.String(job.AllNum)
	mapLabel["FailNum"] = gconv.String(job.FailNum)
	mapLabel["BaselineName"] = gconv.String(job.BaselineName)
	mapLabel["ScriptName"] = job.ScriptName
	mapLabel["ProbabilityOfSuccess"] = gconv.String(float64(job.SuccessNum) / float64(job.AllNum) * 100)
	jobMetrics.With(mapLabel).Set(1)

	//失败详情
	faildList := make([]*entity.FailedNodeResult, 0)
	err := dao.FailedNodeResult.Ctx(ctx).Where("job_id", job.Id).Scan(&faildList)
	if err != nil {
		glog.Error(ctx, "Get failedNodeResult Failed", err)
	}
	if len(faildList) == 0 {
		return
	}
	baseline := entity.DesiredResult{}
	err = dao.DesiredResult.Ctx(ctx).Where("id", job.BaselineId).Scan(&baseline)
	if err != nil {
		glog.Error(ctx, "Get DesiredResult Failed", err)
	}
	json, err := gjson.LoadJson(baseline.BaselineJson)

	if err != nil {
		glog.Error(ctx, "Load BaseLine Failed", err)
	}
	for _, node := range faildList {
		actulaResult, _ := gjson.LoadJson(node.ResultJson)
		for key, value := range json.Map() {
			if value != actulaResult.Map()[key] {
				if actulaResult.Map()[key] == "" {
					actulaResult.Map()[key] = "empty_string"
				}
				failedLabel := make(map[string]string)
				failedLabel["NodeIp"] = node.NodeIp
				if node.ErrMsg == "" {
					node.ErrMsg = "empty_string"
				}
				failedLabel["ErrMsg"] = node.ErrMsg
				failedLabel["CheckJobName"] = job.CheckJobName
				failedLabel["DesiredKey"] = key
				failedLabel["DesiredValue"] = gconv.String(value)
				failedLabel["ActualValue"] = gconv.String(actulaResult.Map()[key])
				failedMetrics.With(failedLabel).Set(1)
			}
		}
	}

}
func (mj *MetricJobManager) MetricBaseLine(ctx context.Context) {
	BaseLineList := make([]*entity.DesiredResult, 0)
	err := dao.DesiredResult.Ctx(ctx).Where("id >", 0).Scan(&BaseLineList)
	if err != nil {
		glog.Error(ctx, "Get All BaseLine Failed")
	}
	for _, baseline := range BaseLineList {
		Json, err := gjson.LoadJson(baseline.BaselineJson)
		if err != nil {
			glog.Error(ctx, "LoadContent failed", err)
		}
		keys := make([]string, 0)
		mapLabel := make(map[string]string)
		keys = append(keys, "BaseLineName", "Creator")
		mapLabel["BaseLineName"] = baseline.BaselineName
		mapLabel["Creator"] = baseline.Creator
		for key, value := range Json.Map() {
			keys = append(keys, key)
			if value == "" {
				mapLabel[key] = "empty_string"
			} else {
				mapLabel[key] = gconv.String(value)
			}

		}
		baselinelabel := prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "BaseLine_" + baseline.BaselineName,
		}, keys,
		)
		prometheus.Register(baselinelabel)
		baselinelabel.With(mapLabel).Set(1)
	}
}
