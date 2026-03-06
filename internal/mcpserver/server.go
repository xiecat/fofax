package mcpserver

import (
	"log"
	"strings"

	"github.com/mark3labs/mcp-go/server"
	"github.com/xiecat/fofax/internal/cli"
	"github.com/xiecat/fofax/internal/printer"
)

type Config struct {
	ListenAddr string
	BasePath   string
	AuthToken  string
	Version    string
}

type Service struct {
	baseOptions *cli.Options
	cfg         Config
}

func Start(options *cli.Options) error {
	cfg := Config{
		ListenAddr: strings.TrimSpace(options.McpListenAddr),
		BasePath:   normalizeBasePath(options.McpBasePath),
		AuthToken:  strings.TrimSpace(options.McpAuthToken),
		Version:    cli.FoFaXVersion,
	}
	if cfg.ListenAddr == "" {
		cfg.ListenAddr = ":18080"
	}
	if cfg.BasePath == "" {
		cfg.BasePath = "/mcp"
	}
	svc := &Service{
		baseOptions: options,
		cfg:         cfg,
	}
	return svc.Start()
}

func (s *Service) Start() error {
	printer.Silent = true
	mcpServer := server.NewMCPServer(
		"fofax-mcp-sse",
		s.cfg.Version,
		server.WithInstructions("Use tools to query FOFA in a read-only and structured way."),
		server.WithToolCapabilities(true),
		server.WithRecovery(),
	)
	s.registerTools(mcpServer)

	sseServer := server.NewSSEServer(
		mcpServer,
		server.WithBasePath(s.cfg.BasePath),
	)

	log.Printf("[fofax-mcp] listening on %s%s (SSE mode, auth enabled)", s.cfg.ListenAddr, s.cfg.BasePath)
	return sseServer.Start(s.cfg.ListenAddr)
}

func (s *Service) cloneOptions() *cli.Options {
	cloned := *s.baseOptions
	cloned.Silent = true
	return &cloned
}

func normalizeBasePath(basePath string) string {
	b := strings.TrimSpace(basePath)
	if b == "" || b == "/" {
		return "/mcp"
	}
	b = "/" + strings.Trim(b, "/")
	return b
}
