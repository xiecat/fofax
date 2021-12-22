package fx

import (
	"fmt"
	"fofax/internal/printer"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Tinfo struct {
	Name  string `table:"Name"`
	Value string `table:"Value"`
}

func GenDefaultPlugin(pluginFile string) error {
	base := Plugin{
		Id:          "fx-2021-01",
		Query:       `查询的字符串用于fx="jupyter Unauth" eg:(jupyter Unauth)`,
		RuleName:    "规则名称 eg:(jupyter 未授权)",
		RuleEnglish: "jupyter unauthorized",
		Description: "规则描述",
		Author:      "作者<邮箱>eg:(xiecat)",
		FofaQuery:   `fofa语句 eg:(body="ipython-main-app" && title="Home Page - Select or create a notebook")"`,
		Tag:         []string{"标签1 eg(unauthorized)", "标签2"},
		Source:      "语句来源",
	}
	base.GenPlugin(pluginFile)
	return nil
}

func LoadPlugin(pathFile string) (*Plugin, error) {
	configFile, err := os.Open(pathFile)
	if err != nil {
		printer.Errorf("readPlugin(%s) os.Open failed: %v", pathFile, err)
		return nil, err
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

// StrContain string in strings
func StrContain(query, rawStr2 string) bool {
	if query == "" {
		return true
	}
	if strings.Contains(rawStr2, query) {
		return true
	}
	return false
}

// StrEqual string in strings
func StrEqual(query, rawStr2 string) bool {
	if query == "" {
		return true
	}
	if strings.EqualFold(rawStr2, query) {
		return true
	}
	return false
}

// StrEqualInList string in strings
func StrEqualInList(query string, tags []string) bool {
	if query == "" {
		return true
	}
	for _, checkStr := range tags {
		if strings.EqualFold(query, checkStr) {
			return true
		}
	}
	return false
}
