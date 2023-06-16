package utils

import (
	"fmt"
	"strings"

	"github.com/xiecat/fofax/internal/printer"
)

func FixFullHostIpv6(fullUrl string) string {
	if strings.Count(fullUrl, ":") > 2 && !strings.Contains(fullUrl, "[") && !strings.Contains(fullUrl, "]") {
		start := strings.Index(fullUrl, "://") + 3
		ipv6 := fullUrl[start:]
		return fmt.Sprintf("%s[%s]", fullUrl[0:start], ipv6)
	}
	return fullUrl
}

func FixFullHostInfoScheme(fields []string) string {
	if len(fields) < 4 {
		printer.Errorf("fileds len err: %v", fields)
	}
	//protocol,ip,port,host
	protocol := strings.TrimSpace(fields[0])
	protocol = strings.ReplaceAll(protocol, "_", "-")
	//ip := strings.TrimSpace(fields[1])
	//port := strings.TrimSpace(fields[2])
	host := strings.TrimSpace(fields[3])
	schemaType := strings.TrimSpace(fields[3])

	if strings.HasPrefix(schemaType, "https://") {
		return FixFullHostIpv6(host)
	}
	if strings.HasPrefix(schemaType, "http://") {
		return FixFullHostIpv6(host)
	}
	return FixFullHostIpv6(fmt.Sprintf("%s://%s", protocol, host))
}
