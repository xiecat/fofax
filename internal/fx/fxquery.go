package fx

import (
	"errors"
	"fmt"
	"fofax/internal/printer"
	"fofax/internal/table"
	"io/ioutil"
	"strings"
)

var Info *FoFaxQuery

type FoFaxQuery struct {
	Plugins map[string]Plugin // query
	Tags    map[string]bool
	Id      map[string]bool
}

func NewFoFaxQuery(fxpath string) *FoFaxQuery {
	fq := &FoFaxQuery{
		Plugins: make(map[string]Plugin),
		Tags:    make(map[string]bool),
		Id:      make(map[string]bool),
	}
	Info = fq
	// 内置
	AddLists(defalutPlugins)
	// 从文件读取
	getFxLists(fxpath)
	return fq
}

func (fx FoFaxQuery) Search(id, query, ruleName, ruleEnglish, Author, tag string) (plugins []Plugin) {
	for _, q := range fx.Plugins {
		if StrEqual(id, q.Id) && StrContain(query, q.Query) && StrContain(ruleName, q.RuleName) &&
			StrContain(ruleEnglish, q.RuleEnglish) && StrContain(Author, q.Author) &&
			StrEqualInList(tag, q.Tag) {
			plugins = append(plugins, q)
		}
	}
	return
}

func (fx FoFaxQuery) SearchSingle(id, query string) (Plugin, error) {
	if id == "" && query == "" {
		table.Output([]Tinfo{{"Error", "id or query not null"}})
	}
	for _, q := range fx.Plugins {
		if StrEqual(id, q.Id) && StrEqual(query, q.Query) {
			return q, nil
		}
	}
	return Plugin{}, errors.New("not found")
}

func (fx FoFaxQuery) SearchTable(id, query, ruleName, ruleEnglish, Author, tag string) {
	type qTable struct {
		Id          string `table:"Id" yaml:"Id"`
		Query       string `table:"Query" yaml:"Query"`             // 查询语法
		RuleName    string `table:"RuleName" yaml:"RuleName"`       // 标题名
		RuleEnglish string `table:"RuleEnglish" yaml:"RuleEnglish"` // 规则英文名
		Author      string `table:"Author" yaml:"Author"`           // 作者
		Tag         string `table:"Tag" yaml:"Tag"`                 // 标签
		Type        FxType `table:"Type" yaml:"-"`                  // 类别
	}

	var results []qTable

	for _, q := range fx.Plugins {
		if StrEqual(id, q.Id) && StrContain(query, q.Query) && StrContain(ruleName, q.RuleName) &&
			StrContain(ruleEnglish, q.RuleEnglish) && StrContain(Author, q.Author) &&
			StrEqualInList(tag, q.Tag) {
			results = append(results, qTable{
				Id:          q.Id,
				Query:       q.Query,
				RuleEnglish: q.RuleEnglish,
				RuleName:    q.RuleName,
				Author:      q.Author,
				Tag:         strings.Join(q.Tag, ","),
				Type:        q.Type,
			})
		}
	}
	table.Output(results)
}

func (fx FoFaxQuery) SearchSingleTable(query string) {
	if query == "" {
		table.Output([]Tinfo{{"Error", "id or query not null"}})
	}
	for _, q := range fx.Plugins {
		if StrEqual(query, q.Id) || StrEqual(query, q.Query) {
			q.ShowInfoTable()
			return
		}
	}
	table.Output([]Tinfo{{"Error", "Not found"}})
}

//type FoFaxQuery struct {
//	Plugins map[string]Plugin // query
//	Tags    map[string]bool
//	Id      map[string]bool
//}

func Add(p Plugin) {
	if _, ok := Info.Id[p.Id]; ok {
		printer.Fatalf("Duplicate entry for  id: %s, other info Query: %s, Author: %s", p.Id, p.Query, p.Author)
		return
	}

	if _, ok := Info.Plugins[p.Query]; ok {
		printer.Fatalf("Duplicate entry for  query: %s, other info id: %s, Author: %s", p.Query, p.Id, p.Author)
		return
	}

	Info.Plugins[p.Query] = p
}

func AddLists(plist []Plugin) {
	for _, p := range plist {
		Add(p)
	}
}

func getFxLists(path string) {
	fs, _ := ioutil.ReadDir(path)
	for _, file := range fs {
		if file.IsDir() {
			fmt.Println(path + file.Name())
			getFxLists(path + file.Name() + "/")
		} else {

			p, err := LoadPlugin(path + file.Name())
			if err != nil {
				printer.Fatalf("load plugins error:%s", err.Error())
			}
			Add(*p)
		}
	}
}
