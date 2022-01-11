package utils

import (
	"net/url"
	"strings"
)

type FixUrl struct {
	HostInfo        string //ip:port 肯定带端口
	FixedHostInfo   string //scheme://host 不一定带端口，默认情况下会省略
	IP              string
	Port            string
	Path            string // 不包含参数？
	Dir             string // 解析后的
	Method          string
	ContentType     string
	Data            string
	PostDataType    string
	u               *url.URL
	mustIP          string
	ParseWithScheme bool //是否解析的时候自带scheme
}

// NewFixUrl 归一化url格式
func NewFixUrl(hostinfo string) (*FixUrl, error) {
	fullurl := hostinfo
	hasScheme := strings.Contains(hostinfo, "://")
	if !hasScheme {
		fullurl = "http://" + fullurl
	}
	fullurl = strings.Trim(fullurl, " \t\r\n")

	u, err := url.Parse(fullurl)
	if err != nil {
		return nil, err
	}
	ipArray := strings.Split(u.Host, ":")
	ip := ipArray[0]
	port := u.Port()

	if len(port) == 0 {
		if u.Scheme == "https" {
			port = "443"
		} else if u.Scheme == "http" {
			port = "80"
		}
	}

	// 存在目标是通过host:port重定向到https://host:port的，
	// 所以如果url是http://x.com就最好不要写成http://x.com:80
	// 否则会导致重定向到https://x.com:80，进行返回"server gave HTTP response to HTTPS client"的错误
	fixedHostInfo := u.Scheme + "://"
	if (port == "80" && u.Scheme == "http") || (port == "443" && u.Scheme == "https") {
		fixedHostInfo += ip
	} else {
		fixedHostInfo += u.Host
	}

	return &FixUrl{
		HostInfo:        ip + ":" + port,
		u:               u,
		IP:              ip,
		Port:            port,
		Path:            u.Path,
		FixedHostInfo:   fixedHostInfo,
		Method:          "GET",
		ParseWithScheme: hasScheme,
	}, nil
}

func (fu *FixUrl) Scheme() string {
	return fu.u.Scheme
}

func (fu *FixUrl) String() string {
	return fu.u.String()
}

func IsWebsite(url string) bool {

	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return true
	}
	return false
}

func IsHttp(url string) bool {
	if strings.HasPrefix(url, "http://") {
		return true
	}
	return false
}

func IsHttps(url string) bool {
	if strings.HasPrefix(url, "https://") {
		return true
	}
	return false
}
