package fxparser

import "testing"

func TestValidateQueryAcceptsCanonicalQueries(t *testing.T) {
	cases := map[string]string{
		`title="beijing"`:             `title="beijing"`,
		`app="nginx"`:                 `app="nginx"`,
		`country="CN"`:                `country="CN"`,
		`app="nginx"&&country="CN"`:   `app="nginx"&&country="CN"`,
		`(app="nginx"||app="apache")`: `(app="nginx"||app="apache")`,
	}
	for input, want := range cases {
		normalized, issues := ValidateQuery(input)
		if len(issues) > 0 {
			t.Errorf("ValidateQuery(%q) reported issues: %v", input, issues)
			continue
		}
		if normalized != want {
			t.Errorf("ValidateQuery(%q) = %q, want %q", input, normalized, want)
		}
	}
}

func TestValidateQueryRejectsPartialInput(t *testing.T) {
	cases := []string{
		`app="nginx" &&& country="CN"`,
		`app="nginx" country="CN"`,
		`totally_unknown_field="x"`,
	}
	for _, input := range cases {
		normalized, issues := ValidateQuery(input)
		if len(issues) == 0 {
			t.Errorf("ValidateQuery(%q) should report issues, got normalized %q", input, normalized)
		}
	}
}
