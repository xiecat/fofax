package mcpserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/mark3labs/mcp-go/server"
	"github.com/xiecat/fofax/internal/cli"
	"github.com/xiecat/fofax/internal/mcpconf"
	"github.com/xiecat/fofax/internal/printer"
)

type Config struct {
	Transport  string
	ListenAddr string
	BasePath   string
	AuthToken  string
	Version    string
}

type Service struct {
	baseOptions *cli.Options
	cfg         Config
}

// Start runs the MCP service with the transport selected by -mcp-transport
// or the MCP_TRANSPORT environment variable.
func Start(options *cli.Options) error {
	cfg := Config{
		Transport:  strings.ToLower(strings.TrimSpace(options.McpTransport)),
		ListenAddr: strings.TrimSpace(options.McpListenAddr),
		BasePath:   mcpconf.NormalizeBasePath(options.McpBasePath),
		AuthToken:  strings.TrimSpace(options.McpAuthToken),
		Version:    cli.FoFaXVersion,
	}
	if cfg.ListenAddr == "" {
		cfg.ListenAddr = mcpconf.DefaultListenAddr
	}
	svc := &Service{
		baseOptions: options,
		cfg:         cfg,
	}
	return svc.Start()
}

func (s *Service) Start() error {
	printer.Silent = true

	switch s.cfg.Transport {
	case mcpconf.TransportStdio:
		return s.serveStdio()
	case mcpconf.TransportHTTP:
		return s.serveHTTP()
	case mcpconf.TransportBoth:
		return s.serveBoth()
	case mcpconf.TransportSSE:
		return s.serveSSE()
	default:
		return fmt.Errorf("unsupported MCP transport %q, use stdio, http, both or sse", s.cfg.Transport)
	}
}

func (s *Service) newMCPServer() *server.MCPServer {
	mcpServer := server.NewMCPServer(
		"fofax-mcp",
		s.cfg.Version,
		server.WithInstructions("Use tools to query FOFA in a read-only and structured way."),
		server.WithToolCapabilities(true),
		server.WithRecovery(),
	)
	s.registerTools(mcpServer)
	return mcpServer
}

func (s *Service) serveStdio() error {
	log.Printf("[fofax-mcp] transport=stdio")
	return server.ServeStdio(s.newMCPServer())
}

func (s *Service) serveHTTP() error {
	handler := s.requireAuth(server.NewStreamableHTTPServer(s.newMCPServer()))
	srv := &http.Server{
		Addr:              s.cfg.ListenAddr,
		Handler:           s.httpMux(handler),
		ReadHeaderTimeout: 10 * time.Second,
	}
	log.Printf("[fofax-mcp] transport=http listening on %s%s (auth enabled, healthz: /healthz)",
		s.cfg.ListenAddr, s.cfg.BasePath)
	return srv.ListenAndServe()
}

func (s *Service) serveSSE() error {
	sseServer := server.NewSSEServer(
		s.newMCPServer(),
		server.WithBasePath(s.cfg.BasePath),
	)
	srv := &http.Server{
		Addr:              s.cfg.ListenAddr,
		Handler:           s.httpMux(s.requireAuth(sseServer)),
		ReadHeaderTimeout: 10 * time.Second,
	}
	log.Printf("[fofax-mcp] transport=sse listening on %s%s (auth enabled, healthz: /healthz)",
		s.cfg.ListenAddr, s.cfg.BasePath)
	return srv.ListenAndServe()
}

func (s *Service) serveBoth() error {
	go func() {
		if err := s.serveHTTP(); err != nil {
			log.Printf("[fofax-mcp] http transport stopped: %v", err)
		}
	}()
	// stdio 是前台传输：客户端关闭 stdin 后整个进程随之退出。
	return s.serveStdio()
}

func (s *Service) httpMux(mcpHandler http.Handler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", s.handleHealthz)
	// 精确匹配 base path 用于 Streamable HTTP，子路径由 SSE 传输自行匹配
	//（例如 /mcp/sse、/mcp/message）。
	mux.Handle(s.cfg.BasePath, mcpHandler)
	mux.Handle(s.cfg.BasePath+"/", mcpHandler)
	return mux
}

func (s *Service) handleHealthz(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		w.Header().Set("Allow", "GET, HEAD")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	payload := map[string]string{
		"status":    "ok",
		"service":   "fofax-mcp",
		"transport": s.cfg.Transport,
		"version":   s.cfg.Version,
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(payload)
}

// requireAuth enforces "Authorization: Bearer <token>" on network transports.
// stdio runs as a local child process and never reaches this middleware.
func (s *Service) requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !authorized(r.Header, s.cfg.AuthToken) {
			w.Header().Set("WWW-Authenticate", `Bearer realm="fofax-mcp"`)
			http.Error(w, "unauthorized: use header Authorization: Bearer <token>", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (s *Service) cloneOptions() *cli.Options {
	cloned := *s.baseOptions
	cloned.Silent = true
	return &cloned
}
