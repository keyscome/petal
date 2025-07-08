

package env

var DefaultSecretKeys = map[string]bool{
	"API_KEY":      true,
	"DB_PASSWORD":  true,
	"ACCESS_TOKEN": true,
}

func AutoMarkSecrets(vars map[string]EnvVar) map[string]EnvVar {
	for k := range vars {
		if _, isSecret := DefaultSecretKeys[k]; isSecret && vars[k].Type == "" {
			val := vars[k]
			val.Type = "secret"
			vars[k] = val  // ✅ 正确地覆盖回去
		}
	}
	return vars
}
