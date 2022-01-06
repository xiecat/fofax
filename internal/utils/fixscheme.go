package utils

import (
	"fmt"
	"strings"

	"github.com/xiecat/fofax/internal/printer"
)

func FixFullHostInfoScheme(fields []string) string {
	if len(fields) < 4 {
		printer.Errorf("fileds len err: %v", fields)
	}
	//protocol,ip,port,host
	protocol := strings.TrimSpace(fields[0])
	protocol = strings.ReplaceAll(protocol, "_", "-")
	ip := strings.TrimSpace(fields[1])
	port := strings.TrimSpace(fields[2])
	host := strings.TrimSpace(fields[3])
	schemaType := strings.TrimSpace(fields[3])

	if strings.HasPrefix(schemaType, "https") {
		return fmt.Sprintf("https://%s:%s", ip, port)
	}
	if strings.HasPrefix(schemaType, "http") {
		return fmt.Sprintf("http://%s:%s", ip, port)
	}
	return fmt.Sprintf("%s://%s", protocol, host)
}
