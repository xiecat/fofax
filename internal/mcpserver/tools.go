package mcpserver

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/xiecat/fofax/internal/cli"
	"github.com/xiecat/fofax/internal/fofa"
)

type searchResult struct {
	Query     string              `json:"query"`
	Fields    []string            `json:"fields"`
	Count     int                 `json:"count"`
	Total     int32               `json:"total"`
	Items     []map[string]string `json:"items"`
	FetchedAt string              `json:"fetched_at"`
}

func (s *Service) registerTools(mcpServer *server.MCPServer) {
	searchTool := mcp.NewTool(
		"fofa_search",
		mcp.WithDescription("Search FOFA assets and return structured records for MCP clients."),
		mcp.WithReadOnlyHintAnnotation(true),
		mcp.WithString("query", mcp.Required(), mcp.Description("FOFA query string, for example: app=\"nginx\" && country=\"CN\"")),
		mcp.WithString("fields", mcp.Description("Comma-separated fields to return. Default: host,port,ip,country"), mcp.DefaultString("host,port,ip,country")),
		mcp.WithNumber("size", mcp.Description("Result size, range 1-10000"), mcp.Min(1), mcp.Max(10000), mcp.DefaultNumber(100)),
		mcp.WithBoolean("next", mcp.Description("Use /api/v1/search/next API"), mcp.DefaultBool(false)),
		mcp.WithBoolean("include", mcp.Description("Include fraud data"), mcp.DefaultBool(false)),
		mcp.WithBoolean("full", mcp.Description("Include historical data"), mcp.DefaultBool(false)),
		mcp.WithNumber("interval_ms", mcp.Description("Request interval in milliseconds"), mcp.Min(0), mcp.DefaultNumber(0)),
	)
	mcpServer.AddTool(searchTool, s.handleSearchTool)

	infoTool := mcp.NewTool(
		"fofax_mcp_info",
		mcp.WithDescription("Return local MCP service configuration metadata."),
		mcp.WithReadOnlyHintAnnotation(true),
	)
	mcpServer.AddTool(infoTool, s.handleInfoTool)
}

func (s *Service) handleInfoTool(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	_ = ctx
	if !authorized(req.Header, s.cfg.AuthToken) {
		return mcp.NewToolResultError("unauthorized: use header Authorization: Bearer <token>"), nil
	}
	data := map[string]any{
		"service":       "fofax-mcp-sse",
		"listen_addr":   s.cfg.ListenAddr,
		"base_path":     s.cfg.BasePath,
		"version":       s.cfg.Version,
		"supports":      []string{"fofa_search", "fofax_mcp_info"},
		"authenticated": true,
	}
	return structured(data), nil
}

func (s *Service) handleSearchTool(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	_ = ctx
	if !authorized(req.Header, s.cfg.AuthToken) {
		return mcp.NewToolResultError("unauthorized: use header Authorization: Bearer <token>"), nil
	}

	query, err := req.RequireString("query")
	if err != nil {
		return mcp.NewToolResultErrorFromErr("query is required", err), nil
	}
	query = strings.TrimSpace(query)
	if query == "" {
		return mcp.NewToolResultError("query cannot be empty"), nil
	}

	fields := strings.TrimSpace(req.GetString("fields", "host,port,ip,country"))
	fieldList := splitFields(fields)
	if len(fieldList) == 0 {
		return mcp.NewToolResultError("fields cannot be empty"), nil
	}

	size := req.GetInt("size", 100)
	if size < 1 || size > 10000 {
		return mcp.NewToolResultError("size must be in range 1-10000"), nil
	}

	options := s.cloneOptions()
	options.FetchSize = size
	options.Include = req.GetBool("include", false)
	options.OldData = req.GetBool("full", false)
	options.Next = req.GetBool("next", false)
	options.ReqIntervalTime = req.GetInt("interval_ms", 0)
	if options.ReqIntervalTime < 0 {
		options.ReqIntervalTime = 0
	}

	records, total, fetchErr := fetchByFields(options, query, fieldList, size)
	if fetchErr != nil {
		return mcp.NewToolResultError(fmt.Sprintf("fofa search failed: %s", fetchErr.Error())), nil
	}

	result := searchResult{
		Query:     query,
		Fields:    fieldList,
		Count:     len(records),
		Total:     total,
		Items:     records,
		FetchedAt: time.Now().Format(time.RFC3339),
	}
	return structured(result), nil
}

func fetchByFields(options *cli.Options, query string, fields []string, maxItems int) ([]map[string]string, int32, error) {
	fo := fofa.NewFoFa(options)
	results := make([]map[string]string, 0, maxItems)
	var total int32
	fo.SetFetchCallback(func(values []string, allSize int32) bool {
		total = allSize
		row := make(map[string]string, len(fields))
		for idx, field := range fields {
			if idx < len(values) {
				row[field] = values[idx]
				continue
			}
			row[field] = ""
		}
		results = append(results, row)
		return len(results) < maxItems
	})

	if ok := fo.FetchField(strings.Join(fields, ","), query); !ok {
		return nil, 0, fmt.Errorf("remote api returned error")
	}
	return results, total, nil
}

func splitFields(value string) []string {
	raw := strings.Split(value, ",")
	fields := make([]string, 0, len(raw))
	for _, item := range raw {
		field := strings.TrimSpace(item)
		if field == "" {
			continue
		}
		fields = append(fields, field)
	}
	return fields
}

func structured(data any) *mcp.CallToolResult {
	text := "ok"
	b, err := json.MarshalIndent(data, "", "  ")
	if err == nil {
		text = string(b)
	}
	return mcp.NewToolResultStructured(data, text)
}
