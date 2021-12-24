package fxparser

import (
	"fmt"
	"testing"
)

type ps struct {
	Source string
	Parse  string
}

var successCase = []ps{
	{
		Source: `title="beijing"`,
		Parse:  `title="beijing"`,
	},

	{
		Source: `header="elastic"`,
		Parse:  `header="elastic"`,
	},

	{
		Source: `body="网络空间测绘"`,
		Parse:  `body="网络空间测绘"`,
	},

	{
		Source: `fid="kIlUsGZ8pT6AtgKSKD63iw=="`,
		Parse:  `fid="kIlUsGZ8pT6AtgKSKD63iw=="`,
	},

	{
		Source: `domain="qq.com"`,
		Parse:  `domain="qq.com"`,
	},

	{
		Source: `icp="京ICP证030173号"`,
		Parse:  `icp="京ICP证030173号"`,
	},

	{
		Source: `js_name="js/jquery.js"`,
		Parse:  `js_name="js/jquery.js"`,
	},

	{
		Source: `js_md5="82ac3f14327a8b7ba49baa208d4eaa15"`,
		Parse:  `js_md5="82ac3f14327a8b7ba49baa208d4eaa15"`,
	},

	{
		Source: `cname="ap21.inst.siteforce.com"`,
		Parse:  `cname="ap21.inst.siteforce.com"`,
	},

	{
		Source: `cname_domain="siteforce.com"`,
		Parse:  `cname_domain="siteforce.com"`,
	},

	{
		Source: `icon_hash="-247388890"`,
		Parse:  `icon_hash="-247388890"`,
	},

	{
		Source: `host=".gov.cn"`,
		Parse:  `host=".gov.cn"`,
	},

	{
		Source: `port="6379"`,
		Parse:  `port="6379"`,
	},

	{
		Source: `ip="1.1.1.1"`,
		Parse:  `ip="1.1.1.1"`,
	},

	{
		Source: `ip="220.181.111.1/24"`,
		Parse:  `ip="220.181.111.1/24"`,
	},

	{
		Source: `status_code="402"`,
		Parse:  `status_code="402"`,
	},

	{
		Source: `protocol="quic"`,
		Parse:  `protocol="quic"`,
	},

	{
		Source: `country="CN"`,
		Parse:  `country="CN"`,
	},

	{
		Source: `region="Xinjiang"`,
		Parse:  `region="Xinjiang"`,
	},

	{
		Source: `city="Ürümqi"`,
		Parse:  `city="Ürümqi"`,
	},

	{
		Source: `cert="baidu"`,
		Parse:  `cert="baidu"`,
	},

	{
		Source: `cert.subject="Oracle Corporation"`,
		Parse:  `cert.subject="Oracle Corporation"`,
	},

	{
		Source: `cert.issuer="DigiCert"`,
		Parse:  `cert.issuer="DigiCert"`,
	},

	{
		Source: `cert.is_valid=true`,
		Parse:  `cert.is_valid=true`,
	},

	{
		Source: `banner=users && protocol=ftp`,
		Parse:  `banner=users&&protocol=ftp`,
	},

	{
		Source: `type=service`,
		Parse:  `type=service`,
	},

	{
		Source: `os="centos"`,
		Parse:  `os="centos"`,
	},

	{
		Source: `server=="Microsoft-IIS/10"`,
		Parse:  `server=="Microsoft-IIS/10"`,
	},

	{
		Source: `app="Microsoft-Exchange"`,
		Parse:  `app="Microsoft-Exchange"`,
	},

	{
		Source: `after="2017" && before="2017-10-01"`,
		Parse:  `after="2017"&&before="2017-10-01"`,
	},

	{
		Source: `asn="19551"`,
		Parse:  `asn="19551"`,
	},

	{
		Source: `org="LLC Baxet"`,
		Parse:  `org="LLC Baxet"`,
	},

	{
		Source: `base_protocol="udp"`,
		Parse:  `base_protocol="udp"`,
	},

	{
		Source: `is_fraud=false`,
		Parse:  `is_fraud=false`,
	},

	{
		Source: `is_honeypot=false`,
		Parse:  `is_honeypot=false`,
	},

	{
		Source: `is_ipv6=true`,
		Parse:  `is_ipv6=true`,
	},

	{
		Source: `is_domain=true`,
		Parse:  `is_domain=true`,
	},

	{
		Source: `port_size="6"`,
		Parse:  `port_size="6"`,
	},

	{
		Source: `port_size_gt="6"`,
		Parse:  `port_size_gt="6"`,
	},

	{
		Source: `port_size_lt="12"`,
		Parse:  `port_size_lt="12"`,
	},

	{
		Source: `ip_ports="80,161"`,
		Parse:  `ip_ports="80,161"`,
	},

	{
		Source: `ip_country="CN"`,
		Parse:  `ip_country="CN"`,
	},

	{
		Source: `ip_region="Zhejiang"`,
		Parse:  `ip_region="Zhejiang"`,
	},

	{
		Source: `ip_city="Hangzhou"`,
		Parse:  `ip_city="Hangzhou"`,
	},

	{
		Source: `ip_after="2021-03-18"`,
		Parse:  `ip_after="2021-03-18"`,
	},

	{
		Source: `ip_before="2019-09-09"`,
		Parse:  `ip_before="2019-09-09"`,
	},
	// 多条件
	{
		Source: `ip_after="2021-03-18"||ip_before="2019-09-09"`,
		Parse:  `ip_after="2021-03-18"||ip_before="2019-09-09"`,
	},

	{
		Source: `title="beijing"&&header="elastic"`,
		Parse:  `title="beijing"&&header="elastic"`,
	},
}

var fofaTx = []ps{
	{
		Source: `banner`,
		Parse:  `"banner"`,
	},
	{
		Source: `banner=`,
		Parse:  `"banner"`,
	},
	{
		Source: `banner=aaa bbb`,
		Parse:  `"banner="aaa bbb"`,
	},
	{
		Source: `banner=aaa bbb&&ccc`,
		Parse:  `banner="aaa bbb"&&"ccc"`,
	},
	{
		Source: `banner=&&type="service"`,
		Parse:  `banner=""type="service"`,
	},
}

var failCase = []ps{}

func TestParserTreeSuccess(t *testing.T) {
	for _, v := range successCase {
		tree, err := parseTree(v.Source)
		if err != nil {
			t.Fatal(v, err)
		}
		if tree.GetText() != v.Parse {
			t.Fatal(fmt.Sprintf("source: %s,parse: %s", v.Source, tree.GetText()))
		}
	}
}

func TestParserTreeFail(t *testing.T) {
	for _, v := range failCase {
		_, err := parseTree(v.Source)
		if err == nil {
			t.Fatal(v, "不应该被解析成功")
		}
	}
}
