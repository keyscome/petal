package env

import (
	"log"
	"os"
	"strings"
)

// MergeFromStringMap converts a simple map[string]string into EnvVar map with type=plain
func MergeFromStringMap(m map[string]string) map[string]EnvVar {
	result := make(map[string]EnvVar)
	for k, v := range m {
		result[k] = EnvVar{
			Value: v,
			Type:  "plain",
		}
	}
	return result
}

// ExpandVars 只处理 ${VAR}
func ExpandVars(input string, vars map[string]string) string {
	const maxDepth = 10
	val := input
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

// StripShellComments removes shell comments from each line
func StripShellComments(input string) string {
	if strings.Contains(input, "# type=secret") {
		log.Println("⚠️ Warning: Script contains '# type=secret'. Be sure to strip before execution.")
	}
	lines := strings.Split(input, "\n")
	var result []string
	for _, line := range lines {
		if idx := strings.Index(line, "#"); idx != -1 {
			line = line[:idx]
		}
		result = append(result, strings.TrimRight(line, " \t"))
	}
	return strings.Join(result, "\n")
}

// ExpandEnvVarsRecursive 展开变量中的引用（如 ${TMP_DIR_ES}）
func ExpandEnvVarsRecursive(vars map[string]EnvVar) map[string]EnvVar {
	// 构建原始扁平 map[string]string
	flat := make(map[string]string)
	for k, v := range vars {
		flat[k] = v.Value
	}
	// 对每个变量做递归替换
	for k, v := range vars {
		vars[k] = EnvVar{
			Value: expandVars(v.Value, flat),
			Type:  v.Type,
		}
	}
	return vars
}
