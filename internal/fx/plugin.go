package fx

import (
	"errors"
	"fofax/internal/printer"
	"fofax/internal/table"
	"io/ioutil"
	"strings"

	"fofax/internal/utils"
	"gopkg.in/yaml.v2"
)

const (
	TypeInline FxType = iota
	TypeYaml
)

type FxType int

func (f FxType) String() string {
	switch f {
	case TypeYaml:
		return "文件规则"
	case TypeInline:
		return "内置规则"
	default:
		return "未知规则类型"
	}
}

type Plugin struct {
	Id          string   `table:"Id" yaml:"Id"`
	Query       string   `table:"Query" yaml:"Query"`             // 查询语法
	RuleName    string   `table:"RuleName" yaml:"RuleName"`       // 标题名
	RuleEnglish string   `table:"RuleEnglish" yaml:"RuleEnglish"` // 规则英文名
	Description string   `table:"Description" yaml:"Description"` // 描述
	Author      string   `table:"Author" yaml:"Author"`           // 作者
	FofaQuery   string   `table:"-" yaml:"FofaQuery"`             // fofa查询
	Tag         []string `table:"Tag" yaml:"Tag"`                 // 标签
	Type        FxType   `table:"Type" yaml:"-"`                  // 类别
	Source      string   `table:"Source" yaml:"Source"`           // 来源
}

func (q *Plugin) ShowInfoTable() {
	results := []Tinfo{
		{"ID", q.Id},
		{"Query", q.Query},
		{"RuleName", q.RuleName},
		{"RuleEnglish", q.RuleEnglish},
		{"Author", q.Author},
		{"FofaQuery", q.FofaQuery},
		{"Tag", strings.Join(q.Tag, ",")},
		{"Type", q.Type.String()},
		{"Description", q.Description},
	}
	table.Output(results)
}

func (f Plugin) QueryString() string {
	return f.FofaQuery
}

func (base *Plugin) GenPlugin(pluginFile string) error {
	if utils.FileExist(pluginFile) {
		return errors.New("File exist please check: " + pluginFile)
	}
	data, _ := yaml.Marshal(base)
	printer.Infof("Will Write Plugin file: %s\n", pluginFile)
	err := ioutil.WriteFile(pluginFile, data, 0644) // 写入
	if err != nil {
		printer.Errorf("%s can't  write Plugin file\n", pluginFile)
		return err
	}
	return nil

}
