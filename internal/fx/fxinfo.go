package fx

var defalutPlugins = []Plugin{
	{
		Id:          "fx-2021-01",
		Query:       "jupyter Unauth",
		RuleName:    "jupyter 未授权",
		RuleEnglish: "jupyter unauthorized",
		Description: "jupyter 未设置密码时可以随意访问",
		Author:      "becivells",
		FofaQuery:   `body="ipython-main-app" && title="Home Page - Select or create a notebook"`,
		Tag:         []string{"unauthorized"},
		Type:        TypeInline,
		Source:      "",
	},
	{
		Id:          "fx-2021-02",
		Query:       "jupyter Unauth1",
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
