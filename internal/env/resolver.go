package env

import (
	"bytes"
	"os"
	"text/template"
)

type Resolver struct {
	merged map[string]EnvVar
}

// 创建合并后的 Resolver
func NewResolver(global, file, task map[string]EnvVar) *Resolver {
	global = ExpandEnvVarsRecursive(global)
	file = ExpandEnvVarsRecursive(file)
	task = ExpandEnvVarsRecursive(task)
	merged := make(map[string]EnvVar)
	for _, src := range []map[string]EnvVar{global, file, task} {
		for k, v := range src {
			merged[k] = v
		}
	}
	return &Resolver{merged: merged}
}

// Flat 返回扁平变量 map（可选是否 mask secret）
func (r *Resolver) Flat(maskSecret bool) map[string]string {
	flat := make(map[string]string)
	for k, v := range r.merged {
		if maskSecret && v.Type == "secret" {
			flat[k] = "*****"
		} else {
			flat[k] = v.Value
		}
	}
	return flat
}

// 展开 ${VAR} 引用，支持递归替换
func expandVars(val string, vars map[string]string) string {
	const maxDepth = 10
	for i := 0; i < maxDepth; i++ {
		next := os.Expand(val, func(key string) string {
			return vars[key]
		})
		if next == val {
			break
		}
		val = next
	}
	return val
}

// Render 渲染输入字符串，支持 ${VAR} 与 {{ .VAR }} 替换，并可脱敏 secret
func (r *Resolver) Render(input string, maskSecret bool) string {
	flat := r.Flat(maskSecret)

	// 先进行 ${VAR} 替换
	expanded := make(map[string]string)
	for k, v := range flat {
		expanded[k] = expandVars(v, flat)
	}

	// 再进行 {{ .VAR }} 模板替换
	tmpl, err := template.New("env").Parse(input)
	if err != nil {
		return input
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, expanded); err != nil {
		return input
	}
	return buf.String()
}

// 检查缺失的必需变量
func (r *Resolver) CheckMissing(keys []string) []string {
	var missing []string
	for _, k := range keys {
		if _, ok := r.merged[k]; !ok {
			missing = append(missing, k)
		}
	}
	return missing
}
