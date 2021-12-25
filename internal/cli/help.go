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
		Usage:   `fofax -q 'app="Grafana"'`,
		Comment: `搜索 fofa 语法 app="Grafana" `,
	},
	{
		Usage:   `fofax -q 'app="Grafana"' -open`,
		Comment: `搜索 fofa 语法 app="Grafana" 并用浏览器打开`,
	},
	{
		Usage:   `fofax -q 'app="Grafana"' -ffi`,
		Comment: `搜索 fofa 语法 app="Grafana" 并且输出带 schema 的数据(url)`,
	},
	{
		Usage:   `fofax -q 'app="Grafana"' -fto`,
		Comment: `搜索 fofa 语法 app="Grafana" 并且输出 url 和 tilte`,
	},
	{
		Usage:   `echo 'app="APACHE-Solr"' | fofax`,
		Comment: `搜索 fofa 语法 app="APACHE-Solr" `,
	},
	{
		Usage:   `echo 'app="APACHE-Solr"' | fofax -ec`,
		Comment: `搜索 fofa 语法 app="APACHE-Solr" 中排除中国的数据`,
	},
	{
		Usage:   `echo 'app="APACHE-Solr"' | fofax -fs 5`,
		Comment: `搜索  app="APACHE-Solr" 中的前五条数据`,
	},
	{
		Usage:   ` echo 'app="APACHE-Solr"' | fofax -fs 10 -e `,
		Comment: `搜索  app="APACHE-Solr" 中排除蜜罐的前十条数据`,
	},
	{
		Usage:   `fofax -use`,
		Comment: "打印 fofa 语法规则",
	},
	{
		Usage:   `fofax -uc https://www.baidu.com`,
		Comment: "搜索百度网站相同证书的站点",
	},
	{
		Usage:   `fofax -uc https://www.baidu.com -open`,
		Comment: "搜索百度网站相同证书的站点, 并用浏览器打开",
	},
	{
		Usage:   `fofax -iu https://www.baidu.com/favicon.ico -fs 5`,
		Comment: "搜索和百度网站相同ico的数据",
	},
	//fx 语法
	{
		Usage:   `fofax -q 'fx="google-reverse"' -fe  -open`,
		Comment: "搜索fx 中 google-reverse 并用浏览器打开, 查询时使用扩展功能必须加 -fe 参数",
	},
	{
		Usage:   `fofax -q 'fx="google-reverse"' -fe`,
		Comment: "搜索 fx 中 google-reverse, 查询时使用扩展功能必须加 -fe 参数",
	},
	{
		Usage:   `fofax -s google`,
		Comment: "搜索含有 google 的 fx 语句",
	},
	{
		Usage:   `fofax -s 'tag=fofa'`,
		Comment: "搜索 fx 中 tag 为 fofa 的语句",
	},
	{
		Usage:   `fofax -s 'tag=fofa;rulename=ism'`,
		Comment: "搜索 fx 中 tag 为 fofa, rulename 中包含 ism 的语句",
	},
}

func PrintSingleUsage() {
	ulen := len(UsageLists)
	sug := UsageLists[rand.Intn(ulen)].String()
	fmt.Println("Tips:")
	fmt.Println("If you want to view the trial in your browser, try -open")
	fmt.Println("")
	if runtime.GOOS == "windows" {
		fmt.Println(`windows 命令行下下需要特别注意双引号转义eg: ( 'fx=\"google-reverse\"' )`)
		sug := "fofax.exe" + strings.TrimPrefix(sug, "fofax")
		fmt.Println(strings.ReplaceAll(sug, "\"", "\\\""))
		return
	}
	magenta(sug)
}
