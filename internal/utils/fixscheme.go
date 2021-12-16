package utils

import (
	"fmt"
	"fofax/internal/printer"
	"strings"
)

func FixFullHostInfoScheme(fields []string) string {
	if len(fields) < 4 {
		printer.Errorf("fileds len err: %v", fields)
	}
	host := strings.TrimSpace(fields[3])

	if strings.HasPrefix(host, "https://") {
		return host
	}
	return fmt.Sprintf("http://%s", host)
}
