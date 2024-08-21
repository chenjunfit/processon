package check

//在mac上调用失败，ansible版本不一致,python版本不一致
import (
	"context"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
)

func AnsibleRun(remotehost string, extraVar map[string]interface{}, yamlPath string) error {
	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Connection: "smart",
		Inventory:  remotehost,
		User:       "root",
		ExtraVars:  extraVar,
	}
	flag := gfile.Exists(yamlPath)
	if !flag {
		glog.Error(context.TODO(), "yamlPath not exists")
		return gerror.New("yamlPath not exists")
	}
	err := playbook.NewAnsiblePlaybookExecute(yamlPath).
		WithPlaybookOptions(ansiblePlaybookOptions).
		Execute(context.TODO())

	if err != nil {
		glog.Error(context.TODO(), "ansible run playbook failed")
	}
	return err
}
