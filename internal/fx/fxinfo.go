package fx

var defalutPlugins = []Plugin{

	{
		Id:          "fx-2021-0001",
		Query:       "Google Reverse",
		RuleName:    "Google反代服务器",
		RuleEnglish: "Google Reverse proxy",
		Description: "不用挂代理就可以访问的Google搜索，但搜索记录可能会被记录。",
		Author:      "fofa",
		FofaQuery:   `body="var c = Array.prototype.slice.call(arguments, 1);return function() {var d=c.slice();"`,
		Tag:         []string{"google"},
		Type:        TypeInline,
		Source:      "https://tp.wjx.top/m/67114073.aspx",
	},
	{
		Id:          "fx-2021-0002",
		Query:       "Python SimpleHTTP",
		RuleName:    "Python SimpleHTTP服务器",
		RuleEnglish: "Python SimpleHTTP Server",
		Description: "Python SimpleHTTP临时服务器",
		Author:      "fofa",
		FofaQuery:   `(server="SimpleHTTP/0.6 Python/3.6.3" || server="SimpleHTTP/0.6 Python/3.6.8" || server="SimpleHTTP/0.6 Python/3.7.0" || server="SimpleHTTP/0.6 Python/3.7.1" || server="SimpleHTTP/0.6 Python/3.7.2" || server="SimpleHTTP/0.6 Python/3.7.3" || server="SimpleHTTP/0.6 Python/3.7.4" || server="SimpleHTTP/0.6 Python/3.7.5" || server="SimpleHTTP/0.6 Python/3.7.6" || server="SimpleHTTP/0.6 Python/3.7.7" || server="SimpleHTTP/0.6 Python/3.7.8" || server="SimpleHTTP/0.6 Python/3.7.9" || server="SimpleHTTP/0.6 Python/3.8.1" || server="SimpleHTTP/0.6 Python/3.8.2" || server="SimpleHTTP/0.6 Python/3.8.3" || server="SimpleHTTP/0.6 Python/3.8.4" || server="SimpleHTTP/0.6 Python/3.8.5" || server="SimpleHTTP/0.6 Python/3.8.6" || server="SimpleHTTP/0.6 Python/3.8.7" || server="SimpleHTTP/0.6 Python/3.8.8" || server="SimpleHTTP/0.6 Python/3.8.9" || server="SimpleHTTP/0.6 Python/3.9.1" || server="SimpleHTTP/0.6 Python/3.9.2" || server="SimpleHTTP/0.6 Python/3.9.3" || server="SimpleHTTP/0.6 Python/3.9.4" || server="SimpleHTTP/0.6 Python/3.9.5" || server="SimpleHTTP/0.6 Python/3.9.6" || server="SimpleHTTP/0.6 Python/3.9.7" || server="SimpleHTTP/0.6 Python/3.9.8" || server="SimpleHTTP/0.6 Python/3.9.9") && title="Directory listing for"`,
		Tag:         []string{"file"},
		Type:        TypeInline,
		Source:      "https://tp.wjx.top/m/67114073.aspx",
	},
	{
		Id:          "fx-2021-0003",
		Query:       "jupyter Unauthorized",
		RuleName:    "jupyter 未授权",
		RuleEnglish: "jupyter unauthorized",
		Description: "jupyter 未设置密码时可以随意访问",
		Author:      "becivells",
		FofaQuery:   `body="ipython-main-app" && title="Home Page - Select or create a notebook"`,
		Tag:         []string{"unauthorized"},
		Type:        TypeInline,
		Source:      "",
	},
}
