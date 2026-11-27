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
	"github.com/xiecat/fofax/internal/fofaapi"
	"github.com/xiecat/fofax/internal/fxparser"
)

// maxSearchSize is the FOFA single-request page limit.
const maxSearchSize = 10000

type searchResult struct {
	Query     string              `json:"query"`
	Fields    []string            `json:"fields"`
	Page      int                 `json:"page"`
	Count     int                 `json:"count"`
	Total     int                 `json:"total"`
	Items     []map[string]string `json:"items"`
	FetchedAt string              `json:"fetched_at"`
}

type statsResult struct {
	Query     string                         `json:"query"`
	Fields    []string                       `json:"fields"`
	Total     int64                          `json:"total"`
	Distinct  map[string]any                 `json:"distinct,omitempty"`
	Aggs      map[string][]fofaapi.StatEntry `json:"aggs,omitempty"`
	FetchedAt string                         `json:"fetched_at"`
}

func (s *Service) registerTools(mcpServer *server.MCPServer) {
	searchTool := mcp.NewTool(
		"fofa_search",
		mcp.WithDescription("Search FOFA assets and return structured records. Use fofa_field_list to learn valid fields and fofa_syntax_check to validate the query first."),
		mcp.WithReadOnlyHintAnnotation(true),
		mcp.WithString("query", mcp.Required(), mcp.Description("FOFA query string, for example: app=\"nginx\" && country=\"CN\"")),
		mcp.WithString("fields", mcp.Description("Comma-separated output fields. Use fofa_field_list for the catalog. Default: host,port,ip,country"), mcp.DefaultString("host,port,ip,country")),
		mcp.WithNumber("size", mcp.Description("Results per page, 1-10000"), mcp.Min(1), mcp.Max(maxSearchSize), mcp.DefaultNumber(100)),
		mcp.WithNumber("page", mcp.Description("Start page, 1-based. Ignored when next=true"), mcp.Min(1), mcp.DefaultNumber(1)),
		mcp.WithBoolean("next", mcp.Description("Use the /api/v1/search/next cursor API instead of page numbers"), mcp.DefaultBool(false)),
		mcp.WithBoolean("include", mcp.Description("Include fraud/honeypot data"), mcp.DefaultBool(false)),
		mcp.WithBoolean("full", mcp.Description("Include historical data"), mcp.DefaultBool(false)),
		mcp.WithNumber("interval_ms", mcp.Description("Delay before the request, in milliseconds; use it to avoid rate limiting"), mcp.Min(0), mcp.DefaultNumber(0)),
	)
	mcpServer.AddTool(searchTool, s.handleSearchTool)

	statsTool := mcp.NewTool(
		"fofa_stats",
		mcp.WithDescription("Aggregate a FOFA query into buckets (country, protocol, port, ...) without paging. Cheaper than paging fofa_search for distributions."),
		mcp.WithReadOnlyHintAnnotation(true),
		mcp.WithString("query", mcp.Required(), mcp.Description("FOFA query string")),
		mcp.WithString("fields", mcp.Description("Comma-separated aggregation fields, for example: country or protocol. Use fofa_field_list."), mcp.DefaultString("country")),
		mcp.WithNumber("size", mcp.Description("Maximum number of buckets per field"), mcp.Min(1), mcp.Max(maxSearchSize)),
		mcp.WithNumber("interval_ms", mcp.Description("Delay before the request, in milliseconds"), mcp.Min(0), mcp.DefaultNumber(0)),
	)
	mcpServer.AddTool(statsTool, s.handleStatsTool)

	fieldTool := mcp.NewTool(
		"fofa_field_list",
		mcp.WithDescription("List valid FOFA query fields and output fields, with operator capability and membership requirements. Use it before writing query or fields arguments."),
		mcp.WithReadOnlyHintAnnotation(true),
		mcp.WithString("kind", mcp.Description("Which catalog to return: query, output, or all"), mcp.Enum("query", "output", "all"), mcp.DefaultString("all")),
		mcp.WithString("group", mcp.Description("Only return query fields of this group, for example assets, app, web, cert, geo, time, aggregation")),
		mcp.WithString("search", mcp.Description("Only return fields whose name contains this text")),
	)
	mcpServer.AddTool(fieldTool, s.handleFieldListTool)

	syntaxTool := mcp.NewTool(
		"fofa_syntax_check",
		mcp.WithDescription("Parse a FOFA query with the bundled grammar and report syntax errors plus non-canonical formatting. Does not call the FOFA API."),
		mcp.WithReadOnlyHintAnnotation(true),
		mcp.WithString("query", mcp.Required(), mcp.Description("FOFA query string to validate")),
	)
	mcpServer.AddTool(syntaxTool, s.handleSyntaxCheckTool)

	infoTool := mcp.NewTool(
		"fofax_mcp_info",
		mcp.WithDescription("Return MCP service configuration and the FOFA account quota."),
		mcp.WithReadOnlyHintAnnotation(true),
		mcp.WithBoolean("include_account", mcp.Description("Also query /api/v1/info/my for quota and membership"), mcp.DefaultBool(false)),
	)
	mcpServer.AddTool(infoTool, s.handleInfoTool)
}

func (s *Service) handleInfoTool(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	_ = ctx
	data := map[string]any{
		"service":     "fofax-mcp",
		"transport":   s.cfg.Transport,
		"listen_addr": s.cfg.ListenAddr,
		"base_path":   s.cfg.BasePath,
		"version":     s.cfg.Version,
		"supports":    []string{"fofa_search", "fofa_stats", "fofa_field_list", "fofa_syntax_check", "fofax_mcp_info"},
	}
	if req.GetBool("include_account", false) {
		options := s.cloneOptions()
		account, err := fofaapi.RequestAccountInfo(options.FoFaURL, options.FoFaKey, 0)
		if err != nil {
			data["account_error"] = err.Error()
		} else {
			data["account"] = map[string]any{
				"username":          account.Username,
				"is_vip":            account.IsVIP,
				"vip_level":         account.VIPLevel,
				"fofa_point":        account.FOFAPoint,
				"remain_free_point": account.RemainFreePnt,
				"expiration":        account.Expiration,
			}
		}
	}
	return structured(data), nil
}

func (s *Service) handleSearchTool(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	_ = ctx
	query, err := req.RequireString("query")
	if err != nil {
		return mcp.NewToolResultErrorFromErr("query is required", err), nil
	}
	query = strings.TrimSpace(query)
	if query == "" {
		return mcp.NewToolResultError("query cannot be empty"), nil
	}

	fieldList, unknown := fofaapi.ValidateOutputFields(req.GetString("fields", "host,port,ip,country"))
	if len(fieldList) == 0 {
		return mcp.NewToolResultError("fields cannot be empty"), nil
	}
	if len(unknown) > 0 {
		return mcp.NewToolResultError(fmt.Sprintf(
			"unknown output fields: %s; call fofa_field_list with kind=output for the catalog",
			strings.Join(unknown, ", "))), nil
	}

	size := req.GetInt("size", 100)
	if size < 1 || size > maxSearchSize {
		return mcp.NewToolResultError(fmt.Sprintf("size must be in range 1-%d", maxSearchSize)), nil
	}

	page := req.GetInt("page", 1)
	if page < 1 {
		return mcp.NewToolResultError("page must be >= 1"), nil
	}

	useNext := req.GetBool("next", false)
	if useNext && page > 1 {
		return mcp.NewToolResultError("page and next are mutually exclusive; next resumes from the cursor returned by FOFA"), nil
	}

	options := s.cloneOptions()
	options.FetchSize = size
	options.StartPage = page
	options.Include = req.GetBool("include", false)
	options.OldData = req.GetBool("full", false)
	options.Next = useNext
	options.ReqIntervalTime = maxInt(0, req.GetInt("interval_ms", 0))

	records, total, fetchErr := fetchByFields(options, query, fieldList, size)
	if fetchErr != nil {
		return mcp.NewToolResultError(fmt.Sprintf("fofa search failed: %s", fetchErr.Error())), nil
	}

	result := searchResult{
		Query:     query,
		Fields:    fieldList,
		Page:      page,
		Count:     len(records),
		Total:     int(total),
		Items:     records,
		FetchedAt: time.Now().Format(time.RFC3339),
	}
	return structured(result), nil
}

func (s *Service) handleStatsTool(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	_ = ctx
	query, err := req.RequireString("query")
	if err != nil {
		return mcp.NewToolResultErrorFromErr("query is required", err), nil
	}
	query = strings.TrimSpace(query)
	if query == "" {
		return mcp.NewToolResultError("query cannot be empty"), nil
	}

	fields, unknown := fofaapi.ValidateOutputFields(req.GetString("fields", "country"))
	if len(fields) == 0 {
		return mcp.NewToolResultError("fields cannot be empty"), nil
	}
	if len(unknown) > 0 {
		return mcp.NewToolResultError(fmt.Sprintf(
			"unknown aggregation fields: %s; call fofa_field_list with kind=output for the catalog",
			strings.Join(unknown, ", "))), nil
	}

	options := s.cloneOptions()
	interval := time.Duration(maxInt(0, req.GetInt("interval_ms", 0))) * time.Millisecond
	stats, err := fofaapi.RequestStats(options.FoFaURL, options.FoFaKey, query, fields, req.GetInt("size", 0), interval)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("fofa stats failed: %s", err.Error())), nil
	}

	result := statsResult{
		Query:     query,
		Fields:    fields,
		Total:     stats.Size,
		Distinct:  stats.Distinct,
		Aggs:      stats.Aggs,
		FetchedAt: time.Now().Format(time.RFC3339),
	}
	return structured(result), nil
}

type queryFieldView struct {
	Name       string `json:"name"`
	Operators  string `json:"operators"`
	Group      string `json:"group"`
	Permission string `json:"permission,omitempty"`
	Summary    string `json:"summary,omitempty"`
}

type outputFieldView struct {
	Name string `json:"name"`
	Note string `json:"note,omitempty"`
}

func (s *Service) handleFieldListTool(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	_ = ctx
	kind := strings.ToLower(strings.TrimSpace(req.GetString("kind", "all")))
	group := strings.ToLower(strings.TrimSpace(req.GetString("group", "")))
	search := strings.ToLower(strings.TrimSpace(req.GetString("search", "")))

	if kind != "query" && kind != "output" && kind != "all" {
		return mcp.NewToolResultError("kind must be query, output, or all"), nil
	}

	matchesSearch := func(name string) bool {
		return search == "" || strings.Contains(strings.ToLower(name), search)
	}

	data := map[string]any{
		"operator_legend": map[string]string{
			"EN":  "supports = and !=",
			"ENF": "supports =, != and fuzzy =~",
			"E":   "supports = only",
		},
	}

	if kind == "query" || kind == "all" {
		views := make([]queryFieldView, 0, len(fofaapi.QueryFields))
		for _, f := range fofaapi.QueryFields {
			if group != "" && !strings.EqualFold(f.Group, group) {
				continue
			}
			if !matchesSearch(f.Name) {
				continue
			}
			views = append(views, queryFieldView{
				Name:       f.Name,
				Operators:  string(f.Operators),
				Group:      f.Group,
				Permission: f.Permission,
				Summary:    f.Summary,
			})
		}
		data["query_fields"] = views
		data["query_field_count"] = len(views)
	}

	if kind == "output" || kind == "all" {
		views := make([]outputFieldView, 0, len(fofaapi.OutputFields))
		for _, f := range fofaapi.OutputFields {
			if !matchesSearch(f.Name) {
				continue
			}
			views = append(views, outputFieldView{Name: f.Name, Note: f.Note})
		}
		data["output_fields"] = views
		data["output_field_count"] = len(views)
	}

	return structured(data), nil
}

func (s *Service) handleSyntaxCheckTool(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	_ = ctx
	query, err := req.RequireString("query")
	if err != nil {
		return mcp.NewToolResultErrorFromErr("query is required", err), nil
	}
	query = strings.TrimSpace(query)
	if query == "" {
		return mcp.NewToolResultError("query cannot be empty"), nil
	}

	normalized, issues := fxparser.ValidateQuery(query)
	data := map[string]any{
		"query":    query,
		"valid":    len(issues) == 0,
		"warnings": []string{},
	}
	if len(issues) > 0 {
		data["issues"] = issues
		return structured(data), nil
	}

	data["normalized"] = normalized
	warnings := make([]string, 0, 2)
	if normalized != query {
		warnings = append(warnings, "query is valid but not in canonical form; the normalized value is what fofax sends to FOFA")
	}
	for _, field := range collectQueryFields(query) {
		if _, ok := fofaapi.LookupQueryField(field); !ok {
			warnings = append(warnings, fmt.Sprintf(
				"field %q is not in the fofa_field_list catalog; confirm it exists on FOFA before using it", field))
		}
	}
	data["warnings"] = warnings
	return structured(data), nil
}

// collectQueryFields extracts the probable left-hand field names of a query
// so callers can be warned about fields missing from the catalog. It is a
// lexical heuristic, not a parse.
func collectQueryFields(query string) []string {
	var fields []string
	seen := map[string]bool{}
	for _, token := range strings.FieldsFunc(query, func(r rune) bool {
		return r == '&' || r == '|' || r == '(' || r == ')' || r == ' ' || r == '\t'
	}) {
		idx := strings.IndexAny(token, "=!~")
		if idx <= 0 {
			continue
		}
		name := strings.TrimSpace(token[:idx])
		name = strings.TrimSuffix(name, "!")
		if name == "" || seen[name] {
			continue
		}
		seen[name] = true
		fields = append(fields, name)
	}
	return fields
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

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func structured(data any) *mcp.CallToolResult {
	text := "ok"
	b, err := json.MarshalIndent(data, "", "  ")
	if err == nil {
		text = string(b)
	}
	return mcp.NewToolResultStructured(data, text)
}
