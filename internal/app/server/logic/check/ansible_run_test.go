package check

import (
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestAnsibleRun(t *testing.T) {
	m := make(map[string]interface{})
	AnsibleRun("121.199.44.128", m, "./site.yaml")
}

func TestSet(t *testing.T) {
	set := gset.NewStrSet()
	set.Add([]string{"1", "1", "2"}...)
	g.Dump(set.Slice())

	str := `["123","34"]`
	hosts := make([]string, 0)
	err := gjson.Unmarshal([]byte(str), &hosts)
	g.Dump(err)
	g.Dump(hosts)
}
