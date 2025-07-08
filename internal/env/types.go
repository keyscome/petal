package env

type EnvVar struct {
	Value string `yaml:"value"`
	Type  string `yaml:"type,omitempty"` // "plain" | "secret"
}