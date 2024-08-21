package check

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"k8s.io/apimachinery/pkg/util/wait"
	"net"
	"processon/internal/app/server/dao"
	"processon/internal/app/server/model"
	"processon/internal/consts"
	"time"
)

type CheckJobManager struct {
	cfg *gcfg.Config
}

func NewCheckJobManager() (*CheckJobManager, error) {
	config, err := gcfg.New()
	if err != nil {
		return nil, err
	}

	return &CheckJobManager{cfg: config}, nil
}

func (cm *CheckJobManager) RunJobCheckManager(ctx context.Context) error {
	internal, err := cm.cfg.Get(ctx, "checkjobConf.CheckSubmitJobInternalSeconds")
	if err != nil {
		return err
	}
	go wait.UntilWithContext(ctx, cm.SpanCheckJob, time.Second*time.Duration(gconv.Int(internal)))
	<-ctx.Done()
	glog.Error(ctx, "RunJobCheckManager Context Done")
	return gerror.New("RunJobCheckManager Context Done")
}

func (cm *CheckJobManager) SpanCheckJob(ctx context.Context) {
	var pool *grpool.Pool
	batch, err := cm.cfg.Get(ctx, "checkjobConf.RunCheckJobBatch")
	if err != nil || gconv.Int(batch) == 0 {
		glog.Error(ctx, "checkjobConf.RunCheckJobBatch :", err)
		pool = grpool.New(10)
	} else {
		pool = grpool.New(gconv.Int(batch))
	}
	jobs := make([]*model.CheckJob, 0)
	err = dao.CheckJob.Ctx(ctx).Where("job_has_synced", 0).Scan(&jobs)
	if err != nil {
		glog.Error(ctx, "Get ALl Jobs: :", err)
	}
	if len(jobs) == 0 {
		return
	}
	for i := 0; i < len(jobs); i++ {
		job := jobs[i]
		pool.Add(ctx, func(ctx context.Context) {
			cm.SubmitJob(ctx, job)
		})
	}

}
func GetIp() string {
	conn, error := net.Dial("udp", "8.8.8.8:80")
	if error != nil {
		glog.Error(context.TODO(), error)
	}
	defer conn.Close()
	ipAddress := conn.LocalAddr().(*net.UDPAddr).IP.String()
	return ipAddress
}
func (cm *CheckJobManager) SubmitJob(ctx context.Context, job *model.CheckJob) {
	set := gset.NewStrSet()
	set.Add(job.IpJson...)
	hosts := gstr.Join(set.Slice(), ",")
	m := make(map[string]interface{})
	dir, err := cm.cfg.Get(ctx, "checkjobConf.NodeRunCheckDir")
	if err != nil {
		glog.Error(ctx, "get checkjobConf.NodeRunCheckDir failed")
		m["NodeRunCheckDir"] = "/root/env-check/"
	}
	m["NodeRunCheckDir"] = gconv.String(dir)
	if err != nil {
		glog.Error(ctx, "get checkjobConf.AgentBinDownloadDir failed")
	}
	m["binFilePath"] = consts.BinFilePath
	port, err := cm.cfg.Get(ctx, "server.address")
	addr := GetIp() + gconv.String(port)
	m["CheckServerAddr"] = addr
	m["ScriptName"] = job.ScriptName
	m["DesiredResultName"] = job.BaselineName
	scriptFilePath := fmt.Sprintf("%s%s", gconv.String(dir), job.ScriptName)
	m["scriptFilePath"] = scriptFilePath
	resultFilePath := fmt.Sprintf("%s%s.json", gconv.String(dir), job.BaselineName)
	m["resultFilePath"] = resultFilePath
	m["JobId"] = job.Id
	m["BaseLineId"] = job.BaselineId
	m["ScriptId"] = job.ScriptId
	m["reportUrl"] = addr + "/processon/failed/report"
	err = AnsibleRun(hosts, m, "resource/template/script.yaml")
	if err != nil {
		glog.Error(ctx, "%s run failed", hosts)
	}
	_, err = dao.CheckJob.Ctx(ctx).Where("id", job.Id).Data(g.Map{"job_has_synced": 1}).Update()
	if err != nil {
		glog.Error(ctx, "update job %d failed", job.Id)
	}
}
