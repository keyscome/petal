package env

var DefaultSecretKeys = map[string]bool{
	"DB_PASSWORD":  true,
	"API_KEY":       true,
	"ACCESS_TOKEN":  true,
}

func AutoMarkSecrets(vars map[string]EnvVar) map[string]EnvVar {
	for k := range vars {
		if _, isSecret := DefaultSecretKeys[k]; isSecret && vars[k].Type == "" {
			val := vars[k]
			val.Type = "secret"
			vars[k] = val
		}
	}
	return vars
}
