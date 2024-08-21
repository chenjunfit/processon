package main

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"net"
	"net/http"
	"os/exec"
	"processon/internal/app/server/model"
	"processon/internal/consts"
	"time"
)

/*
-jobId=4 -report-addr=172.23.11.253:8002/report -result_path=/root/env-check/k8s.json -script_path=/root/env-check/k8s-node.sh"
*/

type Agent struct {
	g.Meta `name:"check-env-agent" brief:"check env"`
}

type CommandInput struct {
	JobId      int    `v:"required" brief:"job id"`
	ReportUrl  string `v:"required" brief:"report data url"`
	ResultPath string `v:"required" brief:"baseline path"`
	ScriptPath string `v:"required" brief:"script path"`
}
type CommandOutput struct {
}

func (agent *Agent) Check(ctx context.Context, in CommandInput) (out *CommandOutput, err error) {
	var same bool = true
	outputStr, err := CommandWithTimeOut(ctx, in.ScriptPath)
	mActual := gconv.Map(gjson.New(outputStr))
	report := model.FailedNodeResult{
		JobId:      in.JobId,
		NodeIp:     GetIp(),
		ResultJson: gconv.String(mActual),
	}
	if err != nil {
		report.ResultJson = `{"msg":"shell exec failed"}`
		report.ErrMsg = err.Error()
	}
	desiredJson, err := gjson.Load(in.ResultPath)
	mDesired := gconv.Map(desiredJson)

	if len(mDesired) != len(mActual) {
		same = false
	}
	for key, value := range mDesired {
		if value != mActual[key] {
			same = false
			break
		}
	}
	report.IsSuccess = same
	actualResultPath := fmt.Sprintf("%s/%s/actual.json", gfile.Dir(gfile.Abs(in.ResultPath)), "actual")
	err = gfile.PutContents(actualResultPath, gconv.String(mActual))
	if err != nil {
		glog.Error(ctx, "write actualResult failed", err)
	}
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	for i := 0; i < 10; i++ {
		res, err := g.Client().Post(ctx, in.ReportUrl, report)
		if err != nil {
			glog.Error(ctx, "err: retrying...", err)
			time.Sleep(time.Second * 5)
		}
		if res.StatusCode == http.StatusOK {
			break
		}
	}
	return
}
func CommandWithTimeOut(ctx context.Context, scriptName string, args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(consts.ExecTimeOut))
	defer cancel()
	cmd := exec.CommandContext(ctx, scriptName, args...)
	if out, err := cmd.Output(); err != nil {
		if ctx.Err() != nil && ctx.Err() == context.DeadlineExceeded {
			return "", gerror.New(context.DeadlineExceeded.Error())
		}
		return gconv.String(out), err
	} else {
		return gconv.String(out), nil
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
func main() {
	cmd, err := gcmd.NewFromObject(Agent{})
	if err != nil {
		panic(err)
	}
	cmd.Run(gctx.New())

}
