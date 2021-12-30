package fofa

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	"fofax/internal/cli"
	"fofax/internal/printer"
	"fofax/internal/utils"
	"github.com/jweny/xhttp"
)

type FoFa struct {
	page    int64
	FetchFn fieldFn
	option  *cli.Options
	client  *xhttp.Client
}

type ApiResults struct {
	Mode    string     `json:"mode"`
	Error   bool       `json:"error"`
	ErrMsg  string     `json:"errmsg"`
	Query   string     `json:"query"`
	Page    int        `json:"page"`
	Size    int        `json:"size"`
	Results [][]string `json:"results"`
}

type fieldFn func(fields []string, allSize int32) bool

//type fixUrlFn func(hostInfo *utils.FixUrl, allSize int32) bool

func NewFoFa(option *cli.Options) *FoFa {
	hOpt := xhttp.NewDefaultClientOptions()
	hOpt.Headers = map[string]string{
		"user-agent": fmt.Sprintf("fofax-client-%s", cli.FoFaXVersion),
	}
	if len(option.Proxy) != 0 {
		printer.Successf("Use Proxy: %s", option.Proxy)
		hOpt.Proxy = option.Proxy
	}
	client, err := xhttp.NewClient(hOpt, nil)
	if err != nil {
		return nil
	}
	return &FoFa{
		option: option,
		client: client,
	}
}

func (f *FoFa) SetFetchCallback(fn func(fields []string, allSize int32) bool) {
	f.FetchFn = fn
}

func (f *FoFa) buildQueryUrl(queryStr string) string {
	return f.option.FoFaURL + queryStr
}

func (f *FoFa) fetchByFields(fields string, queryStr string) bool {
	f.page = 1
	maxSize := f.option.FetchSize
	if maxSize < 0 {
		// max window限制
		maxSize = 10000 * 50000
	}
	// 找到最小值
	perPage := int(math.Min(float64(maxSize), 10000))
	if f.option.Debug {
		printer.Debugf("FoFa Size : %d", perPage)
		printer.Debugf("FoFa input Query of: %s", queryStr)
	}

	for {
		uri := fmt.Sprintf(
			"/api/v1/search/all?email=%s&key=%s&qbase64=%s&size=%d&page=%d&fields=%s",
			f.option.FoFaEmail, f.option.FoFaKey,
			base64.StdEncoding.EncodeToString([]byte(queryStr)),
			perPage,
			f.page,
			fields,
		)

		fullURL := f.buildQueryUrl(uri)
		if f.option.Debug {
			if f.option.ShowPrivacy {
				printer.Debug(fullURL)
			} else {
				hiddenUri := fmt.Sprintf(
					"/api/v1/search/all?email=%s&key=%s&qbase64=%s&size=%d&page=%d&fields=%s",
					"*****@*******", utils.GetHidePasswd(f.option.FoFaKey),
					base64.StdEncoding.EncodeToString([]byte(queryStr)),
					perPage,
					f.page,
					fields,
				)
				printer.Debug(f.buildQueryUrl(hiddenUri))
			}
		}
		hr, _ := http.NewRequest("GET", fullURL, nil)
		req := &xhttp.Request{RawRequest: hr}
		// 发起请求
		ctx := context.Background()
		resp, err := f.client.Do(ctx, req)
		if err != nil {
			printer.Errorf(printer.HandlerLine("request failed: %s" + err.Error()))
			return false
		}
		latency, err := resp.GetLatency()
		if f.option.Debug {
			printer.Debugf("Resp Time: %d/millis", latency.Milliseconds())

		}
		if err != nil {
			printer.Errorf(printer.HandlerLine("request failed: %s" + err.Error()))
			return false
		}
		if resp.GetStatus() != 200 {
			printer.Errorf("Http Status Code : %d", resp.GetStatus())
			return false
		}
		var apiResult ApiResults
		if err := json.Unmarshal(resp.Body, &apiResult); err != nil {
			printer.Errorf("Json Unmarshal Failed: %s", string(resp.Body))
			return false
		}
		if len(apiResult.ErrMsg) != 0 {
			printer.Errorf("FoFa Response ErrMsg: %s", getApiErrInfo(apiResult.ErrMsg))
			return false
		}
		if f.option.Debug {
			printer.Debugf("Fofa Api Query: %s", apiResult.Query)
		}
		printer.Successf("Fetch Data From FoFa: [%d/%d]", len(apiResult.Results), apiResult.Size)
		for _, result := range apiResult.Results {
			//if len(result[0]) == 0 || result[0] == ":0" {
			//	printer.Debug("There is no HostInfo!")
			//	continue
			//	// https://fofa.so/api/v1/search/all?email=lubyruffy@gmail.com&key=xxx&qbase64=YXBwPSJIaWt2aXNpb24tQ2FtZXJhcy1hbmQtU3VydmVpbGxhbmNlIiAmJiBwcm90b2NvbCE9cnN0cA==&size=100&page=1&fields=host
			//}

			if !f.FetchFn(result, int32(apiResult.Size)) {
				return true
			}
			maxSize--
			if maxSize == 0 {
				return true
			}
		}

		// 没有数据，退出
		if len(apiResult.Results) == 0 || len(apiResult.Results) < perPage {
			return true
		}
		f.page++
	}
}

// FetchFullHostInfo 提取完整带协议的字段
func (f *FoFa) FetchFullHostInfo(queryStr string) bool {
	return f.fetchByFields("protocol,ip,port,host", queryStr)
}

// FetchOneField 提取指定的字段
func (f *FoFa) FetchOneField(field, queryStr string) bool {
	return f.fetchByFields(field, queryStr)
}

// FetchTitlesOfDomain 提取 title
func (f *FoFa) FetchTitlesOfDomain(queryStr string) bool {
	return f.fetchByFields("protocol,ip,port,host,title,lastupdatetime", queryStr)
}

func (f *FoFa) Fetch(queryStr string) bool {
	return f.fetchByFields("host,title,lastupdatetime", queryStr)
}

//func (f *FoFa) fetchFn(fields []string, allSize int32) bool {
//	hostInfo := utils.NewFixUrl(
//		fmt.Sprintf("%s://%s:%s",
//			fields[0], fields[1], fields[2]))
//	return f.FetchCallBack(hostInfo, allSize)
//}

func getApiErrInfo(code string) string {
	switch code {
	case "820000":
		return "查询语法错误"
	default:
		return code
	}
}
