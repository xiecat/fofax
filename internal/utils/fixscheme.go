package utils

import (
	"fmt"
	"strings"

	"fofax/internal/printer"
)

func FixFullHostInfoScheme(fields []string) string {
	if len(fields) < 4 {
		printer.Errorf("fileds len err: %v", fields)
	}
	//protocol,ip,port,host
	protocol := strings.TrimSpace(fields[0])
	ip := strings.TrimSpace(fields[1])
	port := strings.TrimSpace(fields[2])
	host := strings.TrimSpace(fields[3])

	if strings.HasPrefix(protocol, "https") {
		return fmt.Sprintf("https://%s:%s", ip, port)
	}
	if strings.HasPrefix(host, "https://") || strings.HasPrefix(host, "http://") {
		return host
	}
	return fmt.Sprintf("http://%s", host)
}
