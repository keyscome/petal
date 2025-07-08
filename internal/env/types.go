// File: internal/env/types.go
package env

type EnvVar struct {
	Value string `yaml:"value"`
	Type  string `yaml:"type"` // "plain" or "secret"
}
