package main

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/oklog/run"
	"os"
	"os/signal"
	"processon/internal/app/server/logic/check"
	_ "processon/internal/boot"
	"syscall"

	"github.com/gogf/gf/v2/os/gctx"

	"processon/internal/cmd"
)

func main() {
	var (
		localG  run.Group
		signals = []os.Signal{
			os.Interrupt, os.Kill, syscall.SIGKILL, syscall.SIGSTOP,
			syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGILL, syscall.SIGTRAP,
			syscall.SIGABRT, syscall.SIGSYS, syscall.SIGTERM,
		}
	)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, signals...)
	ctxAll, cancelAll := context.WithCancel(context.TODO())

	//监听系统信号
	{
		localG.Add(func() error {
			for {
				select {
				case <-ch:
					{
						glog.Error(ctxAll, "recv signal exist", <-ch)
						return gerror.New("recv signal exist")

					}
				case <-ctxAll.Done():
					{
						glog.Error(ctxAll, "recv other signal exist")
						return gerror.New("recv other signal exist")

					}
				}
			}
		}, func(err error) {
			cancelAll()
		})
	}
	//http服务
	{
		localG.Add(func() error {
			for {
				err := cmd.Http.RunWithError(gctx.New())
				if err != nil {
					return err
				}
				select {
				case <-ctxAll.Done():
					{
						glog.Error(ctxAll, "context done ")
						return gerror.New("context done")
					}

				}

			}

		}, func(err error) {
			cancelAll()
		})
	}
	//检查任务
	{
		localG.Add(func() error {
			cm, err := check.NewCheckJobManager()
			if err != nil {
				glog.Error(ctxAll, "NewCheckJobManaer err:", err)
			}
			err = cm.RunJobCheckManager(ctxAll)
			return err
		}, func(err error) {
			cancelAll()
		})
	}
	//计算任务
	{
		localG.Add(func() error {
			cm, err := check.NewCoumputeJobManager()
			if err != nil {
				glog.Error(ctxAll, "NewComputeJobManaer err:", err)
			}
			err = cm.RunComputeJobManager(ctxAll)
			return err
		}, func(err error) {
			cancelAll()
		})
	}
	localG.Run()

}
