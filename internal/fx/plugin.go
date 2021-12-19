package fx

import (
	"errors"
	"fmt"
	"fofax/internal/printer"
	"io/ioutil"
	"os"

	"fofax/internal/utils"
	"gopkg.in/yaml.v2"
)

type FxType int

const (
	TypeInline FxType = iota
	TypeYaml
)

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
	Num         int
	Grammer     string   // 查询语法
	RuleName    string   // 标题名
	RuleEnglish string   // 规则英文名
	Description string   // 描述
	Author      string   // 作者
	FofaQuery   string   // fofa查询
	Tag         []string // 标签
	Type        FxType   // 类别
	Source      string   // 来源
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

func GenDefaultPlugin(pluginFile string) error {
	return nil
}

func LoadPlugin(pathFile string) (*Plugin, error) {
	configFile, err := os.Open(pathFile)
	if err != nil {
		printer.Errorf("readPlugin(%s) os.Open failed: %v", pathFile, err)
	}
	defer configFile.Close()

	content, err := ioutil.ReadAll(configFile)
	if err != nil {
		printer.Errorf("readPlugin(%s) ioutil.ReadAll failed: %v", pathFile, err)
		return nil, err
	}
	plugin := &Plugin{}

	err = yaml.Unmarshal(content, plugin)
	if err != nil {
		return nil, fmt.Errorf("[%s] yaml format err:%s ", pathFile, err.Error())
	}
	return plugin, nil
}
