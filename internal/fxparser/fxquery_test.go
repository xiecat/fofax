package fxparser

import "testing"

var successCase = []string{
	`title="beijing"`,
	`header="elastic"`,
	`body="网络空间测绘"`,
	`fid="kIlUsGZ8pT6AtgKSKD63iw==`,
	`domain="qq.com"`,
	`icp="京ICP证030173号"`,
	`js_name="js/jquery.js"`,
	`js_md5="82ac3f14327a8b7ba49baa208d4eaa15"`,
	`cname="ap21.inst.siteforce.com"`,
	`cname_domain="siteforce.com"`,
	`icon_hash="-247388890"`,
	`host=".gov.cn"`,
	`port="6379"`,
	`ip="1.1.1.1"`,
	`ip="220.181.111.1/24"`,
	`status_code="402"`,
	`protocol="quic"`,
	`country="CN"`,
	`region="Xinjiang"`,
	`city="Ürümqi"`,
	`cert="baidu"`,
	`cert.subject="Oracle Corporation"`,
	`cert.issuer="DigiCert"`,
	`cert.is_valid=true`,
	`banner=users && protocol=ftp`,
	`type=service`,
	`os="centos"`,
	`server=="Microsoft-IIS/10"`,
	`app="Microsoft-Exchange"`,
	`after="2017" && before="2017-10-01"`,
	`asn="19551"`,
	`org="LLC Baxet"`,
	`base_protocol="udp"`,
	`is_fraud=false`,
	`is_honeypot=false`,
	`is_ipv6=true`,
	`is_domain=true`,
	`port_size="6"`,
	`port_size_gt="6"`,
	`port_size_lt="12"`,
	`ip_ports="80,161"`,
	`ip_country="CN"`,
	`ip_region="Zhejiang"`,
	`ip_city="Hangzhou"`,
	`ip_after="2021-03-18"`,
	`ip_before="2019-09-09"`,
}

var failCase = []string{
	``,
}

func TestParserTree(t *testing.T) {
	for _, v := range successCase {
		_, err := parseTree(v)
		if err != nil {
			t.Fatal(err)
		}
	}
}
