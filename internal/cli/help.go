package cli

import (
	"fmt"
	"math/rand"
	"runtime"
	"strings"
)

type fgx struct {
	Usage   string
	Comment string
}

func (f *fgx) String() string {
	return fmt.Sprintf("Comment: %s\nUsage: %s\n", f.Comment, f.Usage)
}

var UsageLists = []fgx{
	{
		Usage:   `fofax -q 'fx="google-reverse"' -fe  -open`,
		Comment: "搜索fx 中 google-reverse 并用浏览器打开, 查询时使用扩展功能必须加 -fe 参数",
	},
	{
		Usage:   `fofax -q 'fx="google-reverse"' -fe`,
		Comment: "搜索 fx 中 google-reverse, 查询时使用扩展功能必须加 -fe 参数",
	},
	{
		Usage:   `fofax  app="Grafana" -open`,
		Comment: `搜索 fofa 语法 app="Grafana" 并用浏览器打开`,
	},
	{
		Usage:   `fofax -s google`,
		Comment: "搜索含有 google 的 fx 语句",
	},
	{
		Usage:   `fofax -ss fx-2021-1001`,
		Comment: "列出 fx 语句的详细内容",
	},
	{
		Usage:   `fofax -use`,
		Comment: "打印 fofa 语法规则",
	},
	{
		Usage:   `fofax -uc https://www.baidu.com -open`,
		Comment: "搜索百度网站相同证书的站点, 并用浏览器打开",
	},
}

func PrintSingleUsage() {
	ulen := len(UsageLists)
	sug := UsageLists[rand.Intn(ulen)].String()
	fmt.Println("Tips:")
	if runtime.GOOS == "windows" {
		fmt.Println(`windows 命令行下下需要特别注意双引号转义eg: ( 'fx=\"google-reverse\"' )`)
		sug := strings.Replace(sug, "fofax", "fofax.exe", 1)
		fmt.Println(strings.ReplaceAll(sug, "\"", "\\\""))
		return
	}
	fmt.Println(sug)
}
