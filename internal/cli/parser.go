package cli

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"fofax/internal/fx"
	"fofax/internal/fxparser"
	"fofax/internal/goflags"
	"fofax/internal/printer"
	"fofax/internal/utils"
)

const (
	NONE_Mode = iota
	Stdin_Mode
	Query_Mode
	File_Mode
	DefaultField    = ""
	DefaultFieldSep = ","
)

type Options struct {
	Mode int
	query
	queryOfFile
	filter
	config
	fxconfig
	FxQuery     *fx.FoFaxQuery
	Version     bool
	Use         bool
	Open        bool // 浏览器打开
	NolimitOpen bool
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
	CoinFile  string `key:"-fcf"`
	// 批量 URL，计算 icon hash 后进行查询
	UrlIconFile string `key:"-iuf"`
	// 通过文件中多个的多个 url 查询其证书
	PeerCertificatesFile string `key:"-ucf"`
}

type filter struct {
	// 填写需要获取的字段如，port
	FetchFields string
	// 指定显示的分隔符
	FetchFieldsSplit string
	// 提取指定根域名的所有 title
	FetchTitlesOfDomain bool
	// 提取指定根域名的所有 title
	FetchJarmOfDomain bool
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
	FoFaKeyFake   string
	Proxy         string
	Debug         bool
	ShowPrivacy   bool
	ConfigFile    string
	Update        bool
	DisableUpdate bool
}
type fxconfig struct {
	FxSearch       string
	FxSearchSingle string
	FxList         bool
	Fxtags         bool
	FxParseTree    bool
	GenFx          string
	FofaExt        bool
	FxDir          string
}

var (
	args  *Options
	flags *goflags.FlagSet
)

func initOptions() {
	rand.Seed(time.Now().UnixNano())
	args = new(Options)
	args.FoFaEmail = os.Getenv("FOFA_EMAIL")
	args.FoFaKey = os.Getenv("FOFA_KEY")
	args.FoFaURL = "https://fofa.so"
	args.FetchSize = 100
	args.FxDir = filepath.Join(filepath.Dir(utils.GetDefaultConf()), "fxrules")
	args.ConfigFile = utils.GetDefaultConf()
	args.Stdin = utils.HasStdin()
	if !args.Stdin {
		banner()
	} else {
		args.Mode = Stdin_Mode
	}
}

func init() {
	initOptions()
	flags = goflags.NewFlagSet()
	flags.SetDescription("fofaX is a command line fofa query tool, simple is the best!")
	flags.StringVar(&args.ConfigFile, "config", args.ConfigFile, "fofax configuration file.The file reading order("+strings.Join(utils.ConfDefaultPath, ",")+")")
	createGroup(
		flags, "config", "CONFIGS",
		flags.StringVarP(&args.FoFaEmail, "fofa-email", "email", args.FoFaEmail, "Fofa API Email"),
		flags.StringVarP(&args.FoFaKey, "fofakey", "key", args.FoFaKey, "Fofa API Key"),
		flags.StringVarP(&args.Proxy, "proxy", "p", "", "proxy for http like http://127.0.0.1:8080"),
		flags.StringVar(&args.FoFaURL, "fofa-url", args.FoFaURL, "Fofa url"),
		flags.BoolVar(&args.Debug, "debug", false, "Debug mode"),
		flags.BoolVarP(&args.ShowPrivacy, "show-privacy", "sp", false, "Debug mode Show Privacy"),
		flags.BoolVar(&args.Update, "update", false, "Update fofax"),
		flags.BoolVarP(&args.DisableUpdate, "disable-update", "du", false, "Close update alerts"),
	)
	createGroup(
		flags, "filters", "FILTERS",
		flags.IntVarP(&args.FetchSize, "fetch-size", "fs", args.FetchSize, "The maximum number of query"),
		flags.BoolVarP(&args.Exclude, "exclude", "e", args.Exclude, "Exclude the honeypot."),
		flags.BoolVarP(&args.ExcludeCountryCN, "exclude-country-cn", "ec", false, "Exclude CN."),
		// 好像没用
		//flags.BoolVarP(&args.UniqByIP, "unique-by-ip", "ubi", args.UniqByIP, "以IP的方式进行去重"),
		flags.BoolVarP(&args.FetchFullHostInfo, "fetch-fullHost-info", "ffi", false, "URL fetch, with scheme, hostname, port"),
		flags.BoolVarP(&args.FetchTitlesOfDomain, "fetch-titles-ofDomain", "fto", false, "Fetch website title"),
		flags.BoolVarP(&args.FetchJarmOfDomain, "fetch-jarm-ofDomain", "fjo", false, "Fetch website jarm"),
		flags.StringVarP(&args.FetchFields, "fetch-fields", "ff", DefaultField, "Fetch by fields.eg: (ip,port)"),
		flags.StringVar(&args.FetchFieldsSplit, "fetch-fields-split", "\t", "Specify characters to split data in different fields"),
	)
	createGroup(
		flags, "query", "Single query/cert/icon",
		flags.StringVarP(&args.Query, "query", "q", args.Query, "FoFa query statement"),
		flags.StringVarP(&args.PeerCertificates, "url-cert", "uc", args.PeerCertificates, "Enter the certificate of the https URL to query"),
		flags.StringVarP(&args.UrlIcon, "url-to-icon-hash", "iu", args.UrlIcon, "Enter the URL of an icon, calculate it and query it"),
		flags.StringVarP(&args.IconFilePath, "icon-file-path", "if", args.IconFilePath, "Calculate the hash of the local icon file, then query it"),
	)
	createGroup(
		flags, "queryFile", "Multiple query/cert/icon",
		flags.StringVarP(&args.CoinFile, "fofa-coin-file", "fcf", args.CoinFile, "Load files downloaded with fofa coins (only use -ffi -fs or not)"),
		flags.StringVarP(&args.QueryFile, "query-file", "qf", args.QueryFile, "Load files, query multiple statements"),
		flags.StringVarP(&args.PeerCertificatesFile, "url-cert-file", "ucf", args.UrlIconFile, "Read the URL from the file, calculate the cert and then query it"),
		flags.StringVarP(&args.UrlIconFile, "icon-hash-url-file", "iuf", args.UrlIconFile, "Retrieve the URL from the file, calculate the icon hash and query it"),
	)
	createGroup(
		flags, "fxgroup", "fx grammer",
		flags.StringVarP(&args.GenFx, "gen", "g", args.GenFx, "Generate fx statement files eg: default_fx.yaml"),
		flags.StringVarP(&args.FxDir, "fxdir", "fd", args.FxDir, "fxdir directory"),
		flags.BoolVarP(&args.FxList, "lists", "l", false, "List of fx statements"),
		flags.BoolVarP(&args.Fxtags, "list-tags", "lt", false, "List fx tags "),
		flags.StringVarP(&args.FxSearch, "search", "s", args.FxSearch, "Search for fx statements. Statements are separated by semicolons eg: id=fx-2021-01;query=\"jupyter Unauth\""),
		flags.BoolVar(&args.FxParseTree, "tree", false, "Print syntax tree"),
		flags.BoolVarP(&args.FofaExt, "fofa-ext", "fe", false, "Using extended syntax(fx)"),
		flags.StringVarP(&args.FxSearchSingle, "show-single", "ss", args.QueryFile, "Display a single fx message"),
	)
	flags.BoolVarP(&args.Version, "version", "v", false, "Show fofaX version")
	flags.BoolVar(&args.Use, "use", false, "Syntax queries")
	flags.BoolVar(&args.Open, "open", false, "Open with your browser only support pipline/-q/-uc/-iu/-if")
	flags.BoolVar(&args.NolimitOpen, "no-limit-open", false, "No limit to the number of openings in your browser")
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
		printer.Successf("Create fx statements File storage directory: %s", args.FxDir)
		err := os.Mkdir(args.FxDir, os.ModePerm)
		if err != nil {
			printer.Fatalf("Unable to create a directory: %s", err.Error())
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

	if args.Version {
		printer.Infof("Version: %s", FoFaXVersion)
		printer.Infof("Branch: %s", Branch)
		printer.Infof("Commit: %s", Commit)
		printer.Infof("Date: %s", Date)
		printer.Infof("CodeName: %s", getVname(FoFaXVersion))
		fmt.Print("fofaX is a command line fofa query tool, simple is the best!\n\n")
		os.Exit(0)
	}

	if args.Use {
		ShowUsage()
		os.Exit(0)
	}
	// 检查更新
	if !args.DisableUpdate {
		checkUpdateInfo()
	}
	// 检查基本信息
	checkFoFaInfo()

	// 检查参数是否互斥
	err := checkMutFlags()
	if err != nil {
		printer.Error(printer.HandlerLine(err.Error()))
		os.Exit(1)
	}
	//检查coinFile 是否存在
	if args.CoinFile != "" && !utils.FileExist(args.CoinFile) {
		printer.Fatalf("file: %s not exist", args.CoinFile)
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
	if flagNum == 0 && args.Mode != Stdin_Mode && !args.Update {
		// return errors.New("query are empty")
		fmt.Print("fofaX is a command line fofa query tool, simple is the best!\n\n")
		PrintSingleUsage()
		os.Exit(0)
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

func checkUpdateInfo() {
	lastFile := filepath.Join(filepath.Dir(utils.GetDefaultConf()), ".fofax-last")
	if !utils.FileExist(lastFile) {
		_ = os.MkdirAll(filepath.Dir(lastFile), os.ModePerm)
		ioutil.WriteFile(lastFile, []byte(strings.TrimSpace(Date)), os.ModePerm)
	}
	lastTime, err := ioutil.ReadFile(lastFile)
	if err != nil {
		printer.Error(err)
		return
	}
	lasTime, err := time.Parse("2006-01-02T15:04:05Z", strings.TrimSpace(string(lastTime)))
	if err != nil {
		printer.Error(err)
		return
	}
	if -time.Until(lasTime) > 24*time.Hour || args.Update {
		err := updateTips(FoFaXVersion)
		if err != nil {
			printer.Error(err.Error())
		}
		ioutil.WriteFile(lastFile, []byte(time.Now().Format("2006-01-02T15:04:05Z")), os.ModePerm)
	}

}
