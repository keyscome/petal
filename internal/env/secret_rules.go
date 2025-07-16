package env

// 默认的敏感字段标识
var defaultSecrets = map[string]bool{
	"DB_PASSWORD":  true,
	"API_KEY":      true,
	"ACCESS_TOKEN": true,
}

// 自动将默认敏感字段标记为 secret 类型
func AutoMarkSecrets(vars map[string]EnvVar) map[string]EnvVar {
	for k, v := range vars {
		if _, isSecret := defaultSecrets[k]; isSecret && v.Type == "" {
			v.Type = "secret"
			vars[k] = v
		}
	}
	return vars
}
