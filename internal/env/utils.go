package env

import "os"

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
