package fofaapi

import (
	"fmt"
	"sort"
	"strings"
)

// OperatorKind mirrors the FOFA grammar capabilities of a query field.
type OperatorKind string

const (
	// OpEqualNotFuzzy allows only "=" and "!=".
	OpEqualNotFuzzy OperatorKind = "EN"
	// OpEqualFuzzy allows "=", "!=" and fuzzy matching.
	OpEqualFuzzy OperatorKind = "ENF"
	// OpEqualOnly allows only "=".
	OpEqualOnly OperatorKind = "E"
)

// QueryField describes one field usable on the left side of a FOFA query.
type QueryField struct {
	Name       string
	Operators  OperatorKind
	Permission string // empty means no extra membership required
	Group      string
	Summary    string
}

// OutputField describes one field usable in the "fields" argument of a search.
type OutputField struct {
	Name string
	Note string
}

// QueryFields is the catalog of query-side field names derived from the FOFA
// grammar shipped with the repository (grammer.json).
var QueryFields = []QueryField{
	{Name: "ip", Operators: OpEqualNotFuzzy, Group: "assets", Summary: "IPv4/IPv6 address or C-class subnet"},
	{Name: "port", Operators: OpEqualFuzzy, Group: "assets", Summary: "Open port number"},
	{Name: "domain", Operators: OpEqualFuzzy, Group: "assets", Summary: "Root domain"},
	{Name: "host", Operators: OpEqualFuzzy, Group: "assets", Summary: "Hostname or IP:port"},
	{Name: "os", Operators: OpEqualFuzzy, Group: "assets", Summary: "Operating system"},
	{Name: "server", Operators: OpEqualFuzzy, Group: "assets", Summary: "Server banner product"},
	{Name: "asn", Operators: OpEqualFuzzy, Group: "assets", Summary: "Autonomous system number"},
	{Name: "org", Operators: OpEqualFuzzy, Group: "assets", Summary: "Organization owning the IP"},
	{Name: "is_domain", Operators: OpEqualOnly, Group: "assets", Summary: "true/false: asset is under a domain"},
	{Name: "is_ipv6", Operators: OpEqualOnly, Group: "assets", Summary: "true/false: asset is IPv6"},

	{Name: "app", Operators: OpEqualOnly, Group: "app", Summary: "Application fingerprint name"},
	{Name: "fid", Operators: OpEqualNotFuzzy, Group: "app", Summary: "FOFA fingerprint id"},
	{Name: "product", Operators: OpEqualNotFuzzy, Group: "app", Summary: "Product name"},
	{Name: "product.version", Operators: OpEqualNotFuzzy, Group: "app", Summary: "Product version"},
	{Name: "category", Operators: OpEqualNotFuzzy, Group: "app", Summary: "Application category"},
	{Name: "type", Operators: OpEqualOnly, Group: "app", Summary: "Service type"},
	{Name: "cloud_name", Operators: OpEqualFuzzy, Group: "app", Summary: "Cloud provider name"},
	{Name: "is_cloud", Operators: OpEqualOnly, Group: "app", Summary: "true/false: asset is on a cloud"},
	{Name: "is_fraud", Operators: OpEqualOnly, Permission: "Pro+", Group: "app", Summary: "true/false: flagged as fraud"},
	{Name: "is_honeypot", Operators: OpEqualOnly, Permission: "Pro+", Group: "app", Summary: "true/false: flagged as honeypot"},

	{Name: "protocol", Operators: OpEqualFuzzy, Group: "protocol", Summary: "Application protocol"},
	{Name: "banner", Operators: OpEqualNotFuzzy, Group: "protocol", Summary: "Full protocol banner"},
	{Name: "banner_hash", Operators: OpEqualNotFuzzy, Permission: "Personal+", Group: "protocol", Summary: "Banner hash"},
	{Name: "banner_fid", Operators: OpEqualNotFuzzy, Permission: "Personal+", Group: "protocol", Summary: "Banner fingerprint id"},
	{Name: "base_protocol", Operators: OpEqualNotFuzzy, Group: "protocol", Summary: "Transport protocol, e.g. tcp/udp"},

	{Name: "title", Operators: OpEqualFuzzy, Group: "web", Summary: "HTML title"},
	{Name: "header", Operators: OpEqualNotFuzzy, Group: "web", Summary: "HTTP response headers"},
	{Name: "header_hash", Operators: OpEqualFuzzy, Permission: "Personal+", Group: "web", Summary: "Header hash"},
	{Name: "body", Operators: OpEqualNotFuzzy, Group: "web", Summary: "HTML body"},
	{Name: "body_hash", Operators: OpEqualNotFuzzy, Group: "web", Summary: "Body hash"},
	{Name: "js_name", Operators: OpEqualFuzzy, Group: "web", Summary: "Referenced JavaScript path"},
	{Name: "js_md5", Operators: OpEqualFuzzy, Group: "web", Summary: "Referenced JavaScript md5"},
	{Name: "cname", Operators: OpEqualFuzzy, Group: "web", Summary: "CNAME record"},
	{Name: "cname_domain", Operators: OpEqualFuzzy, Group: "web", Summary: "Domain of the CNAME record"},
	{Name: "icon_hash", Operators: OpEqualNotFuzzy, Group: "web", Summary: "favicon hash"},
	{Name: "status_code", Operators: OpEqualNotFuzzy, Group: "web", Summary: "HTTP status code"},
	{Name: "icp", Operators: OpEqualFuzzy, Group: "web", Summary: "ICP filing number"},
	{Name: "sdk_hash", Operators: OpEqualNotFuzzy, Permission: "Business+", Group: "web", Summary: "SDK hash"},

	{Name: "country", Operators: OpEqualNotFuzzy, Group: "geo", Summary: "Country code"},
	{Name: "region", Operators: OpEqualNotFuzzy, Group: "geo", Summary: "Province/region"},
	{Name: "city", Operators: OpEqualNotFuzzy, Group: "geo", Summary: "City"},

	{Name: "cert", Operators: OpEqualNotFuzzy, Group: "cert", Summary: "Certificate blob"},
	{Name: "cert.subject", Operators: OpEqualFuzzy, Group: "cert", Summary: "Certificate subject"},
	{Name: "cert.issuer", Operators: OpEqualFuzzy, Group: "cert", Summary: "Certificate issuer"},
	{Name: "cert.subject.org", Operators: OpEqualFuzzy, Group: "cert", Summary: "Certificate subject organization"},
	{Name: "cert.subject.cn", Operators: OpEqualFuzzy, Group: "cert", Summary: "Certificate subject common name"},
	{Name: "cert.issuer.org", Operators: OpEqualFuzzy, Group: "cert", Summary: "Certificate issuer organization"},
	{Name: "cert.issuer.cn", Operators: OpEqualFuzzy, Group: "cert", Summary: "Certificate issuer common name"},
	{Name: "cert.domain", Operators: OpEqualFuzzy, Group: "cert", Summary: "Domain covered by the certificate"},
	{Name: "cert.is_equal", Operators: OpEqualOnly, Permission: "Personal+", Group: "cert", Summary: "true/false: certificate matches asset"},
	{Name: "cert.is_valid", Operators: OpEqualOnly, Permission: "Personal+", Group: "cert", Summary: "true/false: certificate is valid"},
	{Name: "cert.is_match", Operators: OpEqualOnly, Permission: "Personal+", Group: "cert", Summary: "true/false: certificate domain matches"},
	{Name: "cert.is_expired", Operators: OpEqualOnly, Permission: "Personal+", Group: "cert", Summary: "true/false: certificate expired"},
	{Name: "jarm", Operators: OpEqualFuzzy, Group: "cert", Summary: "JARM fingerprint"},
	{Name: "tls.version", Operators: OpEqualNotFuzzy, Group: "cert", Summary: "TLS version"},
	{Name: "tls.ja3s", Operators: OpEqualFuzzy, Group: "cert", Summary: "TLS JA3S fingerprint"},
	{Name: "cert.sn", Operators: OpEqualNotFuzzy, Group: "cert", Summary: "Certificate serial number"},
	{Name: "cert.not_after.after", Operators: OpEqualOnly, Group: "cert", Summary: "Certificate expiry after date"},
	{Name: "cert.not_after.before", Operators: OpEqualOnly, Group: "cert", Summary: "Certificate expiry before date"},
	{Name: "cert.not_before.after", Operators: OpEqualOnly, Group: "cert", Summary: "Certificate validity start after date"},
	{Name: "cert.not_before.before", Operators: OpEqualOnly, Group: "cert", Summary: "Certificate validity start before date"},

	{Name: "after", Operators: OpEqualOnly, Permission: "Personal+", Group: "time", Summary: "Asset first seen after date (YYYY-MM-DD)"},
	{Name: "before", Operators: OpEqualOnly, Permission: "Personal+", Group: "time", Summary: "Asset first seen before date (YYYY-MM-DD)"},
	{Name: "after&before", Operators: OpEqualOnly, Group: "time", Summary: "Range syntax; prefer after=... && before=..."},

	{Name: "ip_filter()", Operators: OpEqualOnly, Permission: "Business+", Group: "aggregation", Summary: "Same-origin IP filter"},
	{Name: "ip_exclude()", Operators: OpEqualOnly, Permission: "Business+", Group: "aggregation", Summary: "Same-origin IP exclusion"},
	{Name: "port_size", Operators: OpEqualNotFuzzy, Permission: "Business+", Group: "aggregation", Summary: "Number of open ports on the IP"},
	{Name: "port_size_gt", Operators: OpEqualOnly, Permission: "Business+", Group: "aggregation", Summary: "Port count greater than"},
	{Name: "port_size_lt", Operators: OpEqualOnly, Permission: "Business+", Group: "aggregation", Summary: "Port count less than"},
	{Name: "ip_ports", Operators: OpEqualNotFuzzy, Permission: "Business+", Group: "aggregation", Summary: "All ports of the same IP"},
	{Name: "ip_country", Operators: OpEqualFuzzy, Permission: "Business+", Group: "aggregation", Summary: "Country of the same IP"},
	{Name: "ip_region", Operators: OpEqualFuzzy, Permission: "Business+", Group: "aggregation", Summary: "Region of the same IP"},
	{Name: "ip_city", Operators: OpEqualFuzzy, Permission: "Business+", Group: "aggregation", Summary: "City of the same IP"},
	{Name: "ip_after", Operators: OpEqualNotFuzzy, Permission: "Business+", Group: "aggregation", Summary: "First-seen of the same IP"},
	{Name: "ip_before", Operators: OpEqualNotFuzzy, Permission: "Business+", Group: "aggregation", Summary: "Last-seen of the same IP"},
}

// OutputFields catalogs the field names accepted by the search "fields"
// argument and by the stats "fields" argument.
var OutputFields = []OutputField{
	{Name: "host", Note: "hostname or ip:port"},
	{Name: "ip"},
	{Name: "port"},
	{Name: "protocol"},
	{Name: "base_protocol"},
	{Name: "country", Note: "country code"},
	{Name: "country_name", Note: "localized country name"},
	{Name: "region"},
	{Name: "city"},
	{Name: "longitude"},
	{Name: "latitude"},
	{Name: "asn"},
	{Name: "org"},
	{Name: "domain"},
	{Name: "cname"},
	{Name: "os"},
	{Name: "server"},
	{Name: "product"},
	{Name: "product_category"},
	{Name: "icp"},
	{Name: "title"},
	{Name: "banner"},
	{Name: "cert"},
	{Name: "jarm"},
	{Name: "fid"},
	{Name: "header_hash"},
	{Name: "banner_hash"},
	{Name: "banner_fid"},
	{Name: "tls.version"},
	{Name: "tls.ja3s"},
	{Name: "lastupdatetime"},
	{Name: "header", Note: "size is capped at 2000 when requested"},
	{Name: "body", Note: "may need a higher membership level"},
	{Name: "icon_hash", Note: "may need a higher membership level"},
	{Name: "icon", Note: "may need a higher membership level"},
	{Name: "cname_domain", Note: "may need a higher membership level"},
	{Name: "structinfo", Note: "may need a higher membership level"},
}

var (
	queryFieldIndex  = buildQueryFieldIndex()
	outputFieldIndex = buildOutputFieldIndex()
)

func buildQueryFieldIndex() map[string]QueryField {
	idx := make(map[string]QueryField, len(QueryFields))
	for _, f := range QueryFields {
		idx[strings.ToLower(f.Name)] = f
	}
	return idx
}

func buildOutputFieldIndex() map[string]OutputField {
	idx := make(map[string]OutputField, len(OutputFields))
	for _, f := range OutputFields {
		idx[strings.ToLower(f.Name)] = f
	}
	return idx
}

// LookupQueryField resolves a query field name case-insensitively.
func LookupQueryField(name string) (QueryField, bool) {
	f, ok := queryFieldIndex[strings.ToLower(strings.TrimSpace(name))]
	return f, ok
}

// LookupOutputField resolves an output field name case-insensitively.
func LookupOutputField(name string) (OutputField, bool) {
	f, ok := outputFieldIndex[strings.ToLower(strings.TrimSpace(name))]
	return f, ok
}

// QueryFieldNames returns every known query field name, sorted.
func QueryFieldNames() []string {
	names := make([]string, 0, len(QueryFields))
	for _, f := range QueryFields {
		names = append(names, f.Name)
	}
	sort.Strings(names)
	return names
}

// OutputFieldNames returns every known output field name, sorted.
func OutputFieldNames() []string {
	names := make([]string, 0, len(OutputFields))
	for _, f := range OutputFields {
		names = append(names, f.Name)
	}
	sort.Strings(names)
	return names
}

// ValidateOutputFields splits and validates a comma separated field list.
// It returns the normalized fields and one message per unknown field.
func ValidateOutputFields(raw string) ([]string, []string) {
	parts := strings.Split(raw, ",")
	fields := make([]string, 0, len(parts))
	var unknown []string
	for _, part := range parts {
		name := strings.TrimSpace(part)
		if name == "" {
			continue
		}
		if _, ok := LookupOutputField(name); !ok {
			unknown = append(unknown, name)
			continue
		}
		fields = append(fields, name)
	}
	return fields, unknown
}

// HasOperatorCapability reports whether the field accepts the given operator.
// Only "=", "!=" and fuzzy (=~) are modelled.
func HasOperatorCapability(field QueryField, op string) bool {
	switch op {
	case "=", "==":
		return true
	case "!=":
		return field.Operators != OpEqualOnly
	case "=~":
		return field.Operators == OpEqualFuzzy
	}
	return false
}

// FieldPermissionNote renders the membership requirement of a field.
func FieldPermissionNote(field QueryField) string {
	if field.Permission == "" {
		return ""
	}
	return fmt.Sprintf("%s membership required", field.Permission)
}
