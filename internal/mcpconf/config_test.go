package mcpconf

import "testing"

func TestResolveTransport(t *testing.T) {
	cases := []struct {
		transport string
		sse       bool
		want      string
	}{
		{"", false, ""},
		{"", true, TransportSSE},
		{"StdIO", false, TransportStdio},
		{" http ", false, TransportHTTP},
		{"sse", true, TransportSSE},
	}
	for _, c := range cases {
		if got := ResolveTransport(c.transport, c.sse); got != c.want {
			t.Errorf("ResolveTransport(%q, %v) = %q, want %q", c.transport, c.sse, got, c.want)
		}
	}
}

func TestValidTransport(t *testing.T) {
	for _, transport := range []string{TransportStdio, TransportHTTP, TransportBoth, TransportSSE} {
		if !ValidTransport(transport) {
			t.Errorf("ValidTransport(%q) = false, want true", transport)
		}
	}
	for _, transport := range []string{"", "SSE", "tcp", "stdio "} {
		if ValidTransport(transport) {
			t.Errorf("ValidTransport(%q) = true, want false", transport)
		}
	}
}

func TestNetworkTransport(t *testing.T) {
	cases := map[string]bool{
		TransportStdio: false,
		TransportHTTP:  true,
		TransportBoth:  true,
		TransportSSE:   true,
		"":             false,
	}
	for transport, want := range cases {
		if got := NetworkTransport(transport); got != want {
			t.Errorf("NetworkTransport(%q) = %v, want %v", transport, got, want)
		}
	}
}

func TestListenAddrFromEnv(t *testing.T) {
	cases := []struct {
		host string
		port string
		want string
	}{
		{"", "", DefaultListenAddr},
		{"", "8080", ":8080"},
		{"0.0.0.0", "", "0.0.0.0:18080"},
		{"127.0.0.1", "8080", "127.0.0.1:8080"},
	}
	for _, c := range cases {
		t.Setenv("MCP_HOST", c.host)
		t.Setenv("MCP_PORT", c.port)
		if got := ListenAddrFromEnv(); got != c.want {
			t.Errorf("ListenAddrFromEnv() host=%q port=%q = %q, want %q", c.host, c.port, got, c.want)
		}
	}
}

func TestNormalizeBasePath(t *testing.T) {
	cases := map[string]string{
		"":       DefaultBasePath,
		"/":      DefaultBasePath,
		"mcp":    "/mcp",
		"/mcp":   "/mcp",
		"/mcp/":  "/mcp",
		" /x/y/": "/x/y",
	}
	for input, want := range cases {
		if got := NormalizeBasePath(input); got != want {
			t.Errorf("NormalizeBasePath(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestNormalizeFofaBaseURL(t *testing.T) {
	cases := map[string]string{
		"":                              DefaultFofaBaseURL,
		"https://fofa.info":             "https://fofa.info",
		"https://fofa.info/":            "https://fofa.info",
		"https://fofa.info/api/v1":      "https://fofa.info",
		"https://fofa.info/api/v1/":     "https://fofa.info",
		" http://127.0.0.1:8080/api/v1": "http://127.0.0.1:8080",
		"http://fofa-proxy.local:9000":  "http://fofa-proxy.local:9000",
		"/api/v1":                       DefaultFofaBaseURL,
	}
	for input, want := range cases {
		if got := NormalizeFofaBaseURL(input); got != want {
			t.Errorf("NormalizeFofaBaseURL(%q) = %q, want %q", input, got, want)
		}
	}
}
