package env_test

import (
	"testing"

	"keyscome.io/petal/internal/env"
)

func TestResolver_SecretMasking(t *testing.T) {
	global := map[string]env.EnvVar{
		"VISIBLE": {Value: "public", Type: "plain"},
		"SECRET":  {Value: "supersecret", Type: "secret"},
	}
	file := map[string]env.EnvVar{}
	task := map[string]env.EnvVar{
		"COMBINED": {Value: "${VISIBLE}+${SECRET}", Type: "plain"},
	}

	r := env.NewResolver(global, file, task)

	// 测试 flat 展示
	flatMasked := r.Flat(true)
	if flatMasked["VISIBLE"] != "public" {
		t.Errorf("Expected VISIBLE=public, got %s", flatMasked["VISIBLE"])
	}
	if flatMasked["SECRET"] != "*****" {
		t.Errorf("Expected SECRET=*****, got %s", flatMasked["SECRET"])
	}

	flatRaw := r.Flat(false)
	if flatRaw["SECRET"] != "supersecret" {
		t.Errorf("Expected SECRET=supersecret, got %s", flatRaw["SECRET"])
	}

	// 测试渲染
	input := "Combined={{ .COMBINED }}"
	renderedMasked := r.Render(input, true)
	if renderedMasked != "Combined=public+*****" {
		t.Errorf("Expected masked render, got: %s", renderedMasked)
	}

	renderedRaw := r.Render(input, false)
	if renderedRaw != "Combined=public+supersecret" {
		t.Errorf("Expected full render, got: %s", renderedRaw)
	}
}