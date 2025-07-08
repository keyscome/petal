package env

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
