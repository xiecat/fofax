package fx

import (
	"errors"
	"fmt"
	"fofax/internal/printer"
	"fofax/internal/table"
	"io/ioutil"
	"path/filepath"
	"strings"
)

var Info *FoFaxQuery

type FoFaxQuery struct {
	Plugins []Plugin // query
	Query   map[string]bool
	Tags    map[string]bool
	Id      map[string]bool
}

func NewFoFaxQuery(fxpath string) *FoFaxQuery {
	fq := &FoFaxQuery{
		Plugins: []Plugin{},
		Tags:    make(map[string]bool),
		Query:   make(map[string]bool),
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

func (fx FoFaxQuery) SearchSingle(query string) (Plugin, error) {
	if len(query) < 3 {
		return Plugin{}, errors.New("no found")
	}
	for _, q := range fx.Plugins {
		if StrEqual(query, q.Id) || StrEqual(query, q.Query) {
			return q, nil
		}
	}

	return Plugin{}, errors.New("not found")
}

func (fx FoFaxQuery) SearchExpTab(rawStrs string) {
	var id, query, ruleName, ruleEnglish, author, tag string
	strs := strings.Split(rawStrs, ";")

	if len(strs) == 1 && len(strings.Split(strs[0], "=")) == 1 {
		fx.SearchOrTable(rawStrs, rawStrs, rawStrs, rawStrs, rawStrs, rawStrs)
		return
	}

	for _, expr := range strs {
		exprSplit := strings.Split(expr, "=")
		if len(exprSplit) != 2 {
			printer.Fatal("expr err")
		}
		switch strings.TrimSpace(strings.ToLower(exprSplit[0])) {
		case "id":
			id = trimOther(exprSplit[1])
		case "q":
			query = trimOther(exprSplit[1])
		case "query":
			query = trimOther(exprSplit[1])
		case "r":
			ruleName = trimOther(exprSplit[1])
		case "rulename":
			ruleName = trimOther(exprSplit[1])
		case "re":
			ruleEnglish = trimOther(exprSplit[1])
		case "ruleenglish":
			ruleEnglish = trimOther(exprSplit[1])
		case "a":
			author = trimOther(exprSplit[1])
		case "author":
			author = trimOther(exprSplit[1])
		case "t":
			tag = trimOther(exprSplit[1])
		case "tag":
			tag = trimOther(exprSplit[1])
		}
	}
	printer.Debugf("id=%s,query=%s,ruleName=%s,ruleEnglish=%s,author=%s,tag=%s", id, query, ruleName, ruleEnglish, author, tag)
	fx.SearchTable(id, query, ruleName, ruleEnglish, author, tag)
}

func (fx FoFaxQuery) SearchOrTable(id, query, ruleName, ruleEnglish, Author, tag string) {
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
		if StrContain(id, q.Id) || StrContain(query, q.Query) || StrContain(ruleName, q.RuleName) ||
			StrContain(ruleEnglish, q.RuleEnglish) || StrContain(Author, q.Author) ||
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
	if len(results) == 0 {
		table.Output([]Tinfo{{"Info", "Not found"}})
	}
	table.Output(results)
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
		if StrEqual(id, q.Id) && StrEqual(query, q.Query) && StrContain(ruleName, q.RuleName) &&
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
	if len(results) == 0 {
		table.Output([]Tinfo{{"Info", "Not found"}})
	}
	table.Output(results)
}

func (fx FoFaxQuery) ListTags() {
	tlist := []string{}
	for k := range fx.Tags {
		tlist = append(tlist, k)
	}
	table.Output([]Tinfo{
		{"Tags", strings.Join(tlist, ",")},
	})
}
func (fx FoFaxQuery) SearchSingleTable(query string) {
	if query == "" {
		table.Output([]Tinfo{{"Error", "id or query not null"}})
	}
	for _, q := range fx.Plugins {
		fmt.Println(q.Id, query)
		if StrEqual(query, q.Id) || StrEqual(query, q.Query) {
			q.ShowInfoTable()
			return
		}
	}
	table.Output([]Tinfo{{"Info", "Not found"}})
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
	} else {
		Info.Id[p.Id] = true
	}

	if _, ok := Info.Query[p.Query]; ok {
		printer.Fatalf("Duplicate entry for  query: %s, other info id: %s, Author: %s", p.Query, p.Id, p.Author)
		return
	}

	for _, v := range p.Tag {
		Info.Tags[v] = true
	}
	Info.Query[p.Query] = true
	Info.Plugins = append(Info.Plugins, p)
}

func AddLists(plist []Plugin) {
	for _, p := range plist {
		Add(p)
	}
}

func trimOther(s string) string {
	s = strings.TrimSpace(s)
	if (strings.HasSuffix(s, "'") && strings.HasPrefix(s, "'")) || (strings.HasSuffix(s, "\"") && strings.HasPrefix(s, "\"")) {
		s = s[1 : len(s)-1]
	}
	return s
}

func getFxLists(path string) {
	fs, _ := ioutil.ReadDir(path)
	for _, file := range fs {
		if file.IsDir() {
			fmt.Println(path + file.Name())
			getFxLists(path + file.Name() + "/")
		} else {
			yp := filepath.Join(path, file.Name())
			if strings.HasSuffix(yp, "yaml") || strings.HasSuffix(yp, "yml") {
				p, err := LoadPlugin(yp)
				p.Type = TypeYaml
				p.FileDir = yp
				if err != nil {
					printer.Fatalf("load plugins error:%s", err.Error())
				}
				Add(*p)
			}
		}
	}
}
