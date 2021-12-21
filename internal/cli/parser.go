package cli

import (
	"errors"
	"fmt"
	"fofax/internal/fx"
	"fofax/internal/fxparser"
	"fofax/internal/printer"
	"fofax/internal/utils"
	"os"
	"path/filepath"

	"github.com/projectdiscovery/goflags"
)

const (
	NONE_Mode = iota
	Stdin_Mode
	Query_Mode
	File_Mode
)

type Options struct {
	Mode int
	query
	queryOfFile
	filter
	config
	fxconfig
	FxQuery *fx.FoFaxQuery
	Version bool
	Use     bool
	// 标准输入
	Stdin bool
}

type query struct {
	// 查询的语句
	Query string `key:"-q"`
	// 计算 ico 文件
	IconFilePath string `key:"-if"`
	// 单个 URL，计算 icon hash 后进行查询
	UrlIcon string `key:"-ui"`
	// 输入 url(https) 查询其证书
	PeerCertificates string `key:"-uc"`
}

type queryOfFile struct {
	// 从文件中进行查询
	QueryFile string `key:"-qf"`
	// 批量 URL，计算 icon hash 后进行查询
	UrlIconFile string `key:"-iuf"`
	// 通过文件中多个的多个 url 查询其证书
	PeerCertificatesFile string `key:"-ucf"`
}

type filter struct {
	// 填写需要的另一个字段如，port
	FetchOneField string
	// 提取指定根域名的所有 title
	FetchTitlesOfDomain bool
	// 提取完整的 hostInfo，带有 protocol
	FetchFullHostInfo bool
	// 排除干扰
	Exclude bool
	// 排除国家
	ExcludeCountryCN bool
	// 去重 ,好像没用？
	//UniqByIP bool
	// 搜索数
	FetchSize int
}
type config struct {
	// fofa 地址
	FoFaURL   string
	FoFaEmail string
	FoFaKey   string
	// 脱敏密码
	FoFaKeyFake string
	Proxy       string
	Debug       bool
	ConfigFile  string
	FxDir       string
}
type fxconfig struct {
	FxSearch       string
	FxSearchSingle string
	FxList         bool
	Fxtags         bool
	FxParseTree    bool
	GenFx          string
	FofaExt        bool
}

var (
	args  *Options
	flags *goflags.FlagSet
)

func initOptions() {
	args = new(Options)
	args.FoFaEmail = os.Getenv("FOFA_EMAIL")
	args.FoFaKey = os.Getenv("FOFA_KEY")
	args.FoFaURL = "https://fofa.so"
	args.FetchSize = 100
	args.FxDir = getFxConf()
}

func init() {
	initOptions()
	flags = goflags.NewFlagSet()
	flags.SetDescription("fofaX is a command line fofa query tool, simple is the best!")
	//flags.StringVar(&args.ConfigFile, "config", args.ConfigFile, "fofadump configuration file.The file reading order")
	createGroup(
		flags, "config", "CONFIGS",
		flags.StringVarP(&args.FoFaEmail, "fofa-email", "email", args.FoFaEmail, "Fofa API Email"),
		flags.StringVarP(&args.FoFaKey, "fofakey", "key", args.FoFaKey, "Fofa API Key"),
		flags.StringVarP(&args.Proxy, "proxy", "p", "", "proxy for http like http://127.0.0.1:8080"),
		flags.StringVar(&args.FoFaURL, "fofa-url", args.FoFaURL, "Fofa url"),
		flags.BoolVar(&args.Debug, "debug", false, "Debug mode"),
	)
	createGroup(
		flags, "filters", "FILTERS:",
		flags.IntVarP(&args.FetchSize, "fetch-size", "fs", args.FetchSize, "The maximum number of query"),
		flags.BoolVarP(&args.Exclude, "exclude", "e", args.Exclude, "Exclude the honeypot."),
		flags.BoolVarP(&args.ExcludeCountryCN, "exclude-country-cn", "ec", false, "Exclude CN."),
		// 好像没用
		//flags.BoolVarP(&args.UniqByIP, "unique-by-ip", "ubi", args.UniqByIP, "以IP的方式进行去重"),
		flags.BoolVarP(&args.FetchFullHostInfo, "fetch-fullHost-info", "ffi", args.FetchFullHostInfo, "URL fetch, with scheme, hostname, port"),
		flags.BoolVarP(&args.FetchTitlesOfDomain, "fetch-titles-ofDomain", "fto", args.FetchTitlesOfDomain, "Fetch website title"),
		// flags.StringVarP(&args.FetchOneField, "fetch-one-field", "fof", args.FetchOneField, "填写需要的另一个字段如，port"),
	)
	createGroup(
		flags, "query", "Single query/ert/icon",
		flags.StringVarP(&args.Query, "query", "q", args.Query, "FoFa query statement"),
		flags.StringVarP(&args.PeerCertificates, "url-cert", "uc", args.PeerCertificates, "Enter the certificate of the https URL to query"),
		flags.StringVarP(&args.UrlIcon, "url-to-icon-hash", "ui", args.UrlIcon, "Enter the URL of an icon, calculate it and query it"),
		flags.StringVarP(&args.IconFilePath, "icon-file-path", "if", args.IconFilePath, "Calculate the hash of the local icon file, then query it"),
	)
	createGroup(
		flags, "queryFile", "Multiple query/cert/icon",
		flags.StringVarP(&args.QueryFile, "query-file", "qf", args.QueryFile, "Load files, query multiple statements"),
		flags.StringVarP(&args.PeerCertificatesFile, "url-cert-file", "ucf", args.UrlIconFile, "Read the URL from the file, calculate the cert and then query it"),
		flags.StringVarP(&args.UrlIconFile, "icon-hash-url-file", "iuf", args.UrlIconFile, "Retrieve the URL from the file, calculate the icon hash and query it"),
	)
	createGroup(
		flags, "fxgroup", "fx grammer",
		flags.StringVarP(&args.GenFx, "gen", "g", args.GenFx, "生成 fx 语法文件 eg: default_fx.yaml"),
		flags.BoolVarP(&args.FxList, "lists", "l", false, "列出fx语法列表"),
		flags.BoolVarP(&args.Fxtags, "list-tags", "lt", false, "列出fx语法列表"),
		flags.StringVarP(&args.FxSearch, "search", "s", args.FxSearch, "搜索 语句用分号分开 eg: id=fx-2021-01;query=\"jupyter Unauth\""),
		flags.BoolVar(&args.FxParseTree, "tree", false, "打印语法树"),
		flags.BoolVar(&args.FofaExt, "fofa-ext", true, "使用扩展语法(fx)"),
		flags.StringVarP(&args.FxSearchSingle, "show-single", "ss", args.QueryFile, "显示单个 fx 信息"),
	)
	flags.BoolVarP(&args.Version, "version", "v", false, "Show fofaX version")
	flags.BoolVar(&args.Use, "use", false, "Syntax queries")
	err := flags.Parse()
	if err != nil {
		printer.Error(printer.HandlerLine("Parse err :" + err.Error()))
		os.Exit(1)
	}
}

func createGroup(flagSet *goflags.FlagSet, name, desc string, flags ...*goflags.FlagData) {
	flagSet.SetGroup(name, desc)
	for _, currentFlag := range flags {
		currentFlag.Group(name)
	}
}
func ParseFxOptions() {
	if !utils.FileExist(args.FxDir) {
		printer.Infof("create  dir fxdir: %s", args.FxDir)
		err := os.Mkdir(args.FxDir, os.ModePerm)
		if err != nil {
			printer.Fatalf("无法创建目录: %s", err.Error())
		}
	}

	args.FxQuery = fx.NewFoFaxQuery(args.FxDir)
	if args.GenFx != "" {
		fx.GenDefaultPlugin(args.GenFx)
		os.Exit(0)
	}
	if args.FxList {
		args.FxQuery.SearchExpTab("")
		os.Exit(0)
	}
	if args.Fxtags {
		args.FxQuery.ListTags()
		os.Exit(0)
	}
	if args.FxSearchSingle != "" {
		args.FxQuery.SearchSingleTable(args.FxSearchSingle)
		os.Exit(0)
	}
	if args.FxSearch != "" {
		args.FxQuery.SearchExpTab(args.FxSearch)
		os.Exit(0)
	}
	if args.FxParseTree {
		fxparser.PrintParserTree(args.Query)
		os.Exit(1)
	}

}

func ParseOptions() *Options {
	ParseFxOptions()

	args.Stdin = utils.HasStdin()
	if !args.Stdin {
		banner()
	} else {
		args.Mode = Stdin_Mode
	}

	if args.Version {
		printer.Infof("Version: %s", FoFaXVersion)
		printer.Infof("Branch: %s", Branch)
		printer.Infof("Commit: %s", Commit)
		printer.Infof("Date: %s", Date)
		fmt.Println("fofaX is a command line fofa query tool, simple is the best!")
		os.Exit(0)
	}

	if args.Use {
		ShowUsage()
		os.Exit(0)
	}

	// 检查基本信息
	checkFoFaInfo()

	// 检查参数是否互斥
	err := checkMutFlags()
	if err != nil {
		printer.Error(printer.HandlerLine(err.Error()))
		os.Exit(1)
	}

	return args
}

// 用于设置互斥参数
func checkMutFlags() error {
	var flagNum int
	var flagStr string
	// 选定 `key:"xx"`
	queryMap, err := utils.StructToMap(args.query, "key")
	if err != nil {
		printer.Error(printer.HandlerLine("Struct To Map err :" + err.Error()))
		return nil
	}
	for k, v := range queryMap {
		if len(v.(string)) != 0 {

			flagNum += 1
		}
		flagStr += k + "/"
	}
	if flagNum == 1 {
		args.Mode = Query_Mode
	}
	queryFileMap, err := utils.StructToMap(args.queryOfFile, "key")
	if err != nil {
		printer.Error(printer.HandlerLine("Struct To Map err :" + err.Error()))
		return nil
	}
	for k, v := range queryFileMap {
		if len(v.(string)) != 0 {
			flagNum += 1
		}
		flagStr += k + "/"
	}
	// 要么单一扫描，要么从文件中批量扫描
	// 单一模式和批量模式互斥
	// 单一模式、批量模式中的各个参数也互斥
	if flagNum > 1 {
		return errors.New("these " + flagStr + " are mutually exclusive")
	}
	// 不输入 query 也应当提醒
	if flagNum == 0 && args.Mode != Stdin_Mode {
		return errors.New("query are empty")
	}
	if args.Mode != Query_Mode {
		args.Mode = File_Mode
	}
	return nil
}

// 检查 email，key
func checkFoFaInfo() {
	if args.FoFaKey == "" || args.FoFaEmail == "" {
		printer.Error("FoFaKey or FoFaEmail is empty")
		os.Exit(1)
	}
}

func getFxConf() (home string) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "fofa.yaml"
	}
	fxdir := filepath.Join(home, ".config", "fofax", "fxdir")
	if !utils.FileExist(fxdir) {
		printer.Infof("create  dir fxdir: %s", fxdir)
		err := os.MkdirAll(fxdir, os.ModePerm)
		if err != nil {
			printer.Fatalf("无法创建目录: %s", err.Error())
		}
	}

	return fxdir
}
