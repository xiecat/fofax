package utils

import (
	"errors"
	"fmt"
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

// is_ipv6=true && port="8080"
// "https://2407:c080:17ef:ffff::7703:d86e"
func FixFofaHostsHttpIpv6(ipv6URL string) (*url.URL, error) {
	// 提取 IPv6 地址部分
	start := strings.Index(ipv6URL, "://") + 3
	if start == -1 {
		return nil, errors.New("invalid host： " + ipv6URL)
	}
	ipv6 := ipv6URL[start:]

	// 构建规范化的 IPv6 URL
	formattedURL := fmt.Sprintf("%s[%s]", ipv6URL[0:start], ipv6)
	return url.Parse(formattedURL)
}

// "[240e:468:810:ab69:2042:7ff:fe58:881a]:8080"
func FetchIpv6AndIpv4HostAndPort(uHost string) string {

	if strings.HasPrefix(uHost, "[") {
		strat := strings.Index(uHost, "[") + 1
		end := strings.Index(uHost, "]")
		return "[" + uHost[strat:end] + "]"
	}
	return strings.Split(uHost, ":")[0]
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
		// "https://2407:c080:17ef:ffff::7703:d86e"

		u, err = FixFofaHostsHttpIpv6(fullurl)
		if err != nil {
			return nil, err
		}
	}

	//"[240e:468:810:ab69:2042:7ff:fe58:881a]:8080"
	ip := FetchIpv6AndIpv4HostAndPort(u.Host)
	port := u.Port()
	//ipv6 [240e:468:810:ab69:2042:7ff:fe58:881a]:8080
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
