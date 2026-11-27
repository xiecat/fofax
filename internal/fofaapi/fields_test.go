package fofaapi

import (
	"testing"
)

func TestValidateOutputFields(t *testing.T) {
	fields, unknown := ValidateOutputFields("host, ip ,port")
	if len(unknown) != 0 {
		t.Fatalf("unexpected unknown fields: %v", unknown)
	}
	if len(fields) != 3 || fields[0] != "host" || fields[1] != "ip" || fields[2] != "port" {
		t.Fatalf("unexpected fields: %v", fields)
	}

	fields, unknown = ValidateOutputFields("host,bogus,port")
	if len(fields) != 2 {
		t.Fatalf("expected 2 valid fields, got %v", fields)
	}
	if len(unknown) != 1 || unknown[0] != "bogus" {
		t.Fatalf("expected bogus to be rejected, got %v", unknown)
	}

	if fields, _ := ValidateOutputFields(" , ,"); len(fields) != 0 {
		t.Fatalf("expected no fields, got %v", fields)
	}
}

func TestLookupQueryField(t *testing.T) {
	if f, ok := LookupQueryField("TITLE"); !ok || f.Operators != OpEqualFuzzy {
		t.Fatalf("title lookup failed: %+v ok=%v", f, ok)
	}
	if f, ok := LookupQueryField("is_fraud"); !ok || f.Permission != "Pro+" {
		t.Fatalf("is_fraud permission mismatch: %+v", f)
	}
	if _, ok := LookupQueryField("nope"); ok {
		t.Fatal("unknown field should not resolve")
	}
}

func TestHasOperatorCapability(t *testing.T) {
	title, _ := LookupQueryField("title")
	if !HasOperatorCapability(title, "=") || !HasOperatorCapability(title, "!=") || !HasOperatorCapability(title, "=~") {
		t.Fatal("title should support all operators")
	}

	app, _ := LookupQueryField("app")
	if !HasOperatorCapability(app, "=") {
		t.Fatal("app should support =")
	}
	if HasOperatorCapability(app, "!=") || HasOperatorCapability(app, "=~") {
		t.Fatal("app must reject != and =~")
	}

	ip, _ := LookupQueryField("ip")
	if !HasOperatorCapability(ip, "!=") || HasOperatorCapability(ip, "=~") {
		t.Fatal("ip should support != but not =~")
	}
}

func TestFieldCatalogConsistency(t *testing.T) {
	seen := make(map[string]bool, len(QueryFields))
	for _, f := range QueryFields {
		if seen[f.Name] {
			t.Errorf("duplicate query field %q", f.Name)
		}
		seen[f.Name] = true
		if f.Group == "" {
			t.Errorf("query field %q has no group", f.Name)
		}
		switch f.Operators {
		case OpEqualOnly, OpEqualNotFuzzy, OpEqualFuzzy:
		default:
			t.Errorf("query field %q has invalid operators %q", f.Name, f.Operators)
		}
	}

	seenOutput := make(map[string]bool, len(OutputFields))
	for _, f := range OutputFields {
		if seenOutput[f.Name] {
			t.Errorf("duplicate output field %q", f.Name)
		}
		seenOutput[f.Name] = true
	}

	if len(QueryFieldNames()) != len(QueryFields) || len(OutputFieldNames()) != len(OutputFields) {
		t.Error("field name helpers must list every field")
	}
}
