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
	merged := make(map[string]EnvVar)
	for _, src := range []map[string]EnvVar{global, file, task} {
		for k, v := range src {
			merged[k] = v
		}
	}
	return &Resolver{merged: merged}
}

// 将 EnvVar 映射为扁平的 map[string]string
func (r *Resolver) Flat() map[string]string {
	flat := make(map[string]string)
	for k, v := range r.merged {
		flat[k] = v.Value
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

// 渲染一个字符串，先展开 ${VAR}，再执行 {{ .VAR }}
func (r *Resolver) Render(input string) string {
	flat := r.Flat()

	// 先将 flat 中的值进行 ${VAR} 替换
	expanded := make(map[string]string)
	for k, v := range flat {
		expanded[k] = expandVars(v, flat)
	}

	// 再进行 {{ .VAR }} 替换
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
