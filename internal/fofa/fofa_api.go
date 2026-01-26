package fofa

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/xiecat/fofax/internal/cli"
	"github.com/xiecat/fofax/internal/printer"
	"github.com/xiecat/fofax/internal/utils"
)

type FoFa struct {
	page    int64
	next    string // 用于存储 next API 返回的 next 值
	FetchFn fieldFn
	option  *cli.Options
	client  *http.Client
}

type ApiResults struct {
	Mode    string     `json:"mode"`
	Error   bool       `json:"error"`
	ErrMsg  string     `json:"errmsg"`
	Query   string     `json:"query"`
	Page    int        `json:"page"`
	Size    int        `json:"size"`
	Next    string     `json:"next"` // next API 返回的 next 值，用于下一次请求
	Results [][]string `json:"results"`
}

type fieldFn func(fields []string, allSize int32) bool

//type fixUrlFn func(hostInfo *utils.FixUrl, allSize int32) bool

func NewFoFa(option *cli.Options) *FoFa {

	return &FoFa{
		option: option,
		client: option.Xclient,
	}
}

func (f *FoFa) SetFetchCallback(fn func(fields []string, allSize int32) bool) {
	f.FetchFn = fn
}

func (f *FoFa) buildQueryUrl(queryStr string) string {
	baseURL := f.option.FoFaURL
	if f.option.Next {
		baseURL = "https://fofa.info"
	}
	return baseURL + queryStr
}

func (f *FoFa) fetchByFields(fields string, queryStr string) bool {
	f.page = 1
	f.next = "" // 重置 next 值
	maxSize := f.option.FetchSize
	if maxSize < 0 {
		// max window限制
		maxSize = 10000 * 50000
	}

	for {
		// 找到最小值
		perPage := int(math.Min(float64(maxSize), 10000))
		if f.option.Debug {
			printer.Debugf("FoFa Size : %d", perPage)
			printer.Debugf("FoFa input Query of: %s", queryStr)
		}

		var isOptionsArgs string
		if f.option.Include {
			isOptionsArgs = "&fraud=true"
		}
		if f.option.OldData {
			isOptionsArgs += "&full=true"
		}
		auth := fmt.Sprintf("key=%s", f.option.FoFaKey)

		if f.option.FoFaEmail != "" {
			auth = fmt.Sprintf("email=%s&key=%s", f.option.FoFaEmail, f.option.FoFaKey)
		}

		// 根据 Next 参数选择 API 端点
		apiPath := "/api/v1/search/all"
		if f.option.Next {
			apiPath = "/api/v1/search/next"
		}

		var uri string
		if f.option.Next && f.next != "" {
			// 使用 next API 且已有 next 值，使用 next 参数而不是 page
			uri = fmt.Sprintf(
				"%s?%s%s&qbase64=%s&size=%d&next=%s&fields=%s",
				apiPath, auth, isOptionsArgs,
				base64.StdEncoding.EncodeToString([]byte(queryStr)),
				perPage,
				f.next,
				fields,
			)
		} else {
			// 第一次请求或使用传统 API，使用 page 参数
			uri = fmt.Sprintf(
				"%s?%s%s&qbase64=%s&size=%d&page=%d&fields=%s",
				apiPath, auth, isOptionsArgs,
				base64.StdEncoding.EncodeToString([]byte(queryStr)),
				perPage,
				f.page,
				fields,
			)
		}

		fullURL := f.buildQueryUrl(uri)
		if f.option.Debug {
			printer.Debug(utils.HiddenUrlKey(f.option.ShowPrivacy, fullURL))
		}
		req, err := http.NewRequest("GET", fullURL, nil)

		if err != nil {
			printer.Errorf(printer.HandlerLine("request failed: " + utils.HiddenUrlKey(f.option.ShowPrivacy, err.Error())))
			return false
		}
		if f.option.FetchFields != cli.DefaultField {
			printer.Debugf("Fields : %s", strings.Join(strings.Split(f.option.FetchFields, ","), f.option.FetchFieldsSplit))
		}
		req.Header.Set("fofax-client-%s", cli.FoFaXVersion)
		// 计算时长
		start := time.Now().UnixMilli()
		// 请求
		resp, err := f.client.Do(req)
		if err != nil {
			printer.Errorf(printer.HandlerLine("request failed: " + utils.HiddenUrlKey(f.option.ShowPrivacy, err.Error())))
			return false
		}
		if resp.StatusCode != 200 {
			printer.Errorf("Http Status Code : %d", resp.StatusCode)
			return false
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			printer.Errorf(printer.HandlerLine("body read failed: " + err.Error()))
		}
		if f.option.Debug {
			printer.Debugf("Resp Time: %.2f/millis", float64(time.Now().UnixMilli()-start))
		}

		var apiResult ApiResults
		if err := json.Unmarshal(body, &apiResult); err != nil {
			printer.Errorf("Json Unmarshal Failed: %s", string(body))
			return false
		}
		if len(apiResult.ErrMsg) != 0 && apiResult.Error == true {
			printer.Errorf("FoFa Response ErrMsg: %s", apiResult.ErrMsg)
			return false
		}
		if f.option.Debug {
			printer.Debugf("Fofa Api Query: %s", apiResult.Query)
		}
		if f.option.FetchFFIWithQueryAndSize {
			printer.Successf("Fetch Data From [%s]: [%d/%d]", queryStr, len(apiResult.Results), apiResult.Size)

		} else {
			printer.Successf("Fetch Data From FoFa: [%d/%d]", len(apiResult.Results), apiResult.Size)

		}

		for _, result := range apiResult.Results {

			if !f.FetchFn(result, int32(apiResult.Size)) {
				return true
			}
			//maxSize--
			//if maxSize == 0 {
			//	return true
			//}
		}

		// 如果使用 next API，保存返回的 next 值
		if f.option.Next {
			f.next = apiResult.Next
		}

		// 没有数据，退出
		if len(apiResult.Results) == 0 || maxSize < perPage {
			return true
		}

		// 如果使用 next API 且没有 next 值，说明已经获取完所有数据
		if f.option.Next && f.next == "" {
			return true
		}

		maxSize -= perPage
		if maxSize <= 0 {
			return true
		}

		// 根据使用的 API 类型决定如何分页
		if f.option.Next {
			// 使用 next API，next 值已经在上面保存，继续循环即可
			if !f.option.Coin {
				printer.Infof("Use fofa coins to get more than 10,000 data please use -coin to confirm")
				return true
			}
			printer.Infof("The fofa coin will be deducted !!!")
		} else {
			// 使用传统 API，递增 page
			f.page++
			if !f.option.Coin {
				printer.Infof("Use fofa coins to get more than 10,000 data please use -coin to confirm")
				return true
			}
			printer.Infof("The fofa coin will be deducted !!!")
		}
		time.Sleep(time.Duration(f.option.ReqIntervalTime) * time.Millisecond)
	}
}

// FetchFullHostInfo 提取完整带协议的字段
func (f *FoFa) FetchFullHostInfo(queryStr string) bool {
	return f.fetchByFields("protocol,ip,port,host,city", queryStr)
}

// FetchOneField 提取指定的字段
func (f *FoFa) FetchOneField(field, queryStr string) bool {
	return f.fetchByFields(field, queryStr)
}

// FetchField 提取指定的字段
func (f *FoFa) FetchField(field, queryStr string) bool {
	return f.fetchByFields(field, queryStr)
}

// FetchTitlesOfDomain 提取 title
func (f *FoFa) FetchTitlesOfDomain(queryStr string) bool {
	return f.fetchByFields("protocol,ip,port,host,city,title,country", queryStr)
}

// FetchJarmOfDomain 提取 title
func (f *FoFa) FetchJarmOfDomain(queryStr string) bool {
	return f.fetchByFields("protocol,ip,port,host,city,jarm,country", queryStr)
}

func (f *FoFa) Fetch(queryStr string) bool {
	return f.fetchByFields("host,port,ip,country", queryStr)
}
