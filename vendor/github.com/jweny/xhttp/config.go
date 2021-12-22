package xhttp

import (
	"github.com/jweny/xhttp/xtls"
	"golang.org/x/time/rate"
)

const (
	// MethodGet HTTP method
	MethodGet = "GET"

	// MethodPost HTTP method
	MethodPost = "POST"

	// MethodPut HTTP method
	MethodPut = "PUT"

	// MethodDelete HTTP method
	MethodDelete = "DELETE"

	// MethodPatch HTTP method
	MethodPatch = "PATCH"

	// MethodHead HTTP method
	MethodHead = "HEAD"

	// MethodOptions HTTP method
	MethodOptions = "OPTIONS"

	// MethodConnect HTTP method
	MethodConnect = "CONNECT"

	// MethodTrace HTTP method
	MethodTrace = "TRACE"

	// MethodMove HTTP method
	MethodMove = "MOVE"
)

// HTTPClientOptions http client 全局配置
type HTTPClientOptions struct {
	Proxy 			 string  `json:"proxy" yaml:"proxy" #:"使用代理"`
	DialTimeout         int  `json:"dial_timeout" yaml:"dial_timeout" #:"建立 tcp 连接的超时时间"`
	ReadTimeout         int  `json:"read_timeout" yaml:"read_timeout" #:"读取 http 响应的超时时间"`
	MaxConnsPerHost     int  `json:"max_conns_per_host" yaml:"max_conns_per_host" #:"同一 host 最大允许的连接数"`
	EnableHTTP2         bool `json:"enable_http2" yaml:"enable_http2" #:"是否启用 http2"`
	IdleConnTimeout     int  `json:"-" yaml:"-"`
	MaxIdleConns        int  `json:"-" yaml:"-"`
	TLSHandshakeTimeout int  `json:"-" yaml:"-"`

	FailRetries       int               `json:"fail_retries" yaml:"fail_retries" #:"请求失败的重试次数"`
	MaxRedirect       int               `json:"max_redirect" yaml:"max_redirect" #:"单个请求最大跳转数"`
	MaxRespBodySize   int64             `json:"max_resp_body_size" yaml:"max_resp_body_size" #:"最大允许的响应大小, 默认 2M"`
	MaxQPS            rate.Limit        `json:"max_qps" yaml:"max_qps" #:"每秒最大请求数"`
	AllowMethods      []string          `json:"allow_methods" yaml:"allow_methods" #:"允许的请求方法"`
	Headers           map[string]string `json:"-" yaml:"-"`
	Cookies           map[string]string `json:"cookies" yaml:"-"`
	TlsOptions        *xtls.ClientOptions
	Debug             bool
	DisableKeepAlives bool
	// global limiter
	Limiter 		  *rate.Limiter
}

func NewDefaultClientOptions() *HTTPClientOptions {
	return &HTTPClientOptions{
		// Proxy:               "http://127.0.0.1:8080",
		DialTimeout:     5,
		ReadTimeout:     10,
		IdleConnTimeout: 60,
		// todo 这里改成0
		FailRetries:         2,
		MaxConnsPerHost:     50,
		MaxIdleConns:        50,
		TLSHandshakeTimeout: 5,
		MaxRedirect:         10,
		MaxRespBodySize:     2 << 20, // 4M
		MaxQPS:              500,
		Headers:             make(map[string]string),
		AllowMethods: []string{
			MethodHead,
			MethodGet,
			MethodPost,
			MethodPut,
			MethodPatch,
			MethodDelete,
			MethodOptions,
			MethodConnect,
			MethodTrace,
			MethodMove,
		},
		Cookies:     make(map[string]string),
		EnableHTTP2: false,
		TlsOptions:  xtls.NewDefaultClientOptions(),
		Debug: 		 false,
		DisableKeepAlives: false,
		Limiter: 	 rate.NewLimiter(500, 1),
	}
}

func NewHTTPLimiter(options *HTTPClientOptions) *rate.Limiter {
	return rate.NewLimiter(options.MaxQPS, 1)
}


