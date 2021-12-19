package fx

var Querys = []Plugin{
	{
		Num:         1,
		Grammer:     "jupyterunath",
		RuleName:    "jupyter 未授权",
		RuleEnglish: "jupyter unth",
		Description: "jupyter 未设置密码时可以随意访问",
		Author:      "becivells",
		FofaQuery:   `title="jupyter"`,
		Tag:         []string{"", ""},
		Type:        TypeInline,
		Source:      "",
	},
}
