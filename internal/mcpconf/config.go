// Package mcpconf holds the pure configuration helpers shared by the CLI
// parser and the MCP server: transport selection, listen address resolution
// and FOFA API base URL normalization.
package mcpconf

import (
	"os"
	"strings"
)

const (
	// TransportStdio serves MCP over stdin/stdout.
	TransportStdio = "stdio"
	// TransportHTTP serves MCP over the Streamable HTTP transport.
	TransportHTTP = "http"
	// TransportBoth runs stdio and Streamable HTTP in one process.
	TransportBoth = "both"
	// TransportSSE serves MCP over the legacy SSE transport.
	TransportSSE = "sse"

	// DefaultListenAddr is used when neither -mcp-listen nor MCP_HOST/MCP_PORT is set.
	DefaultListenAddr = ":18080"
	// DefaultBasePath is the HTTP mount point of the MCP endpoint.
	DefaultBasePath = "/mcp"
	// DefaultFofaBaseURL is the public FOFA API host.
	DefaultFofaBaseURL = "https://fofa.info"
)

// ResolveTransport normalizes the requested transport. An empty value means
// MCP is not enabled; the legacy -mcp-sse flag maps to the sse transport.
func ResolveTransport(transport string, legacySSE bool) string {
	t := strings.ToLower(strings.TrimSpace(transport))
	if t != "" {
		return t
	}
	if legacySSE {
		return TransportSSE
	}
	return ""
}

// ValidTransport reports whether transport is a supported value.
func ValidTransport(transport string) bool {
	switch transport {
	case TransportStdio, TransportHTTP, TransportBoth, TransportSSE:
		return true
	}
	return false
}

// NetworkTransport reports whether the transport needs a listener and
// therefore Bearer token authentication.
func NetworkTransport(transport string) bool {
	switch transport {
	case TransportHTTP, TransportBoth, TransportSSE:
		return true
	}
	return false
}

// ListenAddrFromEnv resolves the listen address from MCP_HOST/MCP_PORT,
// falling back to DefaultListenAddr.
func ListenAddrFromEnv() string {
	host := strings.TrimSpace(os.Getenv("MCP_HOST"))
	port := strings.TrimSpace(os.Getenv("MCP_PORT"))
	if host == "" && port == "" {
		return DefaultListenAddr
	}
	if port == "" {
		port = "18080"
	}
	if host == "" {
		return ":" + port
	}
	return host + ":" + port
}

// NormalizeBasePath normalizes the MCP HTTP mount point.
func NormalizeBasePath(basePath string) string {
	b := strings.TrimSpace(basePath)
	if b == "" || b == "/" {
		return DefaultBasePath
	}
	return "/" + strings.Trim(b, "/")
}

// NormalizeFofaBaseURL accepts both "https://fofa.info" and
// "https://fofa.info/api/v1" and returns the host part, because the API
// paths are appended by the client.
func NormalizeFofaBaseURL(raw string) string {
	base := strings.TrimRight(strings.TrimSpace(raw), "/")
	if base == "" {
		return DefaultFofaBaseURL
	}
	base = strings.TrimSuffix(base, "/api/v1")
	base = strings.TrimRight(base, "/")
	if base == "" {
		return DefaultFofaBaseURL
	}
	return base
}
