package env_test

import (
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v2"
	"keyscome.io/petal/internal/env"
)

func TestLoadEnvFile(t *testing.T) {
	path := filepath.Join("testdata", "sample.env")
	t.Logf("Loading env file: %s", path)
	vars, err := env.LoadEnvFile(path)
	if err != nil {
		t.Fatalf("failed to load env file: %v", err)
	}
	t.Logf("Before AutoMarkSecrets: %+v", vars)
	vars = env.AutoMarkSecrets(vars)
	t.Logf("After  AutoMarkSecrets: %+v", vars)

	if vars["TMP_DIR"].Value != "/tmp/sample" {
		t.Errorf("expected TMP_DIR to be /tmp/sample, got %s", vars["TMP_DIR"].Value)
	}
	if vars["DB_PASSWORD"].Type != "secret" {
		t.Errorf("expected DB_PASSWORD to be marked secret")
	}
}

func TestLoadYAMLEnvWithSecret(t *testing.T) {
	yamlContent := `
DB_PASSWORD:
  value: abc123
  type: secret
TMP_DIR:
  value: /tmp/yaml
  type: plain
`
	t.Log("Parsing YAML env block with secret")
	m := make(map[string]env.EnvVar)
	if err := yaml.Unmarshal([]byte(yamlContent), &m); err != nil {
		t.Fatalf("yaml unmarshal error: %v", err)
	}
	if m["DB_PASSWORD"].Type != "secret" || m["TMP_DIR"].Type != "plain" {
		t.Errorf("expected proper type decoding: got %+v", m)
	}
}

func TestResolver_Render(t *testing.T) {
	global := map[string]env.EnvVar{
		"A": {Value: "1", Type: "plain"},
	}
	file := map[string]env.EnvVar{
		"B": {Value: "${A}2", Type: "plain"},
	}
	task := map[string]env.EnvVar{
		"C": {Value: "${B}3", Type: "plain"},
	}
	r := env.NewResolver(global, file, task)
	raw := "ABC={{ .C }}"
	rendered := r.Render(raw, false)
	t.Logf("Render input: %s", raw)
	t.Logf("Flat env: %+v", r.Flat(false))
	t.Logf("Rendered result: %s", rendered)
	if rendered != "ABC=123" {
		t.Errorf("expected rendered result to be ABC=123, got %s", rendered)
	}
}

func TestResolver_Mask(t *testing.T) {
	vars := map[string]env.EnvVar{
		"SECRET":      {Value: "sensitive", Type: "secret"},
		"PLAIN":       {Value: "public", Type: "plain"},
		"DB_PASSWORD": {Value: "abc123", Type: ""},
	}
	vars = env.AutoMarkSecrets(vars)
	r := env.NewResolver(nil, nil, vars)
	masked := r.GetMaskedEnv()
	t.Logf("Masked output: %+v", masked)
	if masked["SECRET"] != "******" {
		t.Errorf("expected SECRET to be masked, got %s", masked["SECRET"])
	}
	if masked["PLAIN"] != "public" {
		t.Errorf("expected PLAIN to be public, got %s", masked["PLAIN"])
	}
	if masked["DB_PASSWORD"] != "******" {
		t.Errorf("expected DB_PASSWORD to be masked, got %s", masked["DB_PASSWORD"])
	}
}

func TestCheckMissing(t *testing.T) {
	r := env.NewResolver(nil, nil, map[string]env.EnvVar{
		"A": {Value: "1", Type: "plain"},
	})
	required := []string{"A", "B"}
	missing := r.CheckMissing(required)
	t.Logf("Missing check for %v → %v", required, missing)
	if len(missing) != 1 || missing[0] != "B" {
		t.Errorf("expected missing [B], got %v", missing)
	}
}

func TestExpandEnvVarsRecursive(t *testing.T) {
	input := map[string]env.EnvVar{
		"A": {Value: "/base", Type: "plain"},
		"B": {Value: "${A}/dir", Type: "secret"},
		"C": {Value: "${B}/file", Type: "plain"},
	}
	result := env.ExpandEnvVarsRecursive(input)

	if result["B"].Value != "/base/dir" {
		t.Errorf("Expected B to be /base/dir, got %s", result["B"].Value)
	}
	if result["C"].Value != "/base/dir/file" {
		t.Errorf("Expected C to be /base/dir/file, got %s", result["C"].Value)
	}
}
