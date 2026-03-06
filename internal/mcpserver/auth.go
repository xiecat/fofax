package mcpserver

import (
	"net/http"
	"strings"
)

func authorized(header http.Header, token string) bool {
	if strings.TrimSpace(token) == "" {
		return false
	}
	value := strings.TrimSpace(header.Get("Authorization"))
	if value == "" {
		return false
	}
	expect := "Bearer " + strings.TrimSpace(token)
	return value == expect
}
