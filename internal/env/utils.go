package env

import (
	"bufio"
	"os"
	"strings"
)

func LoadEnvFile(path string) (map[string]EnvVar, error) {
	result := make(map[string]EnvVar)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		k := strings.TrimSpace(parts[0])
		v := strings.TrimSpace(parts[1])
		// result[k] = EnvVar{Value: v, Type: "plain"}
		result[k] = EnvVar{Value: v} // 默认 Type 留空
	}
	return result, scanner.Err()
}

func MergeFromStringMap(in map[string]string, asType string) map[string]EnvVar {
	res := make(map[string]EnvVar)
	for k, v := range in {
		res[k] = EnvVar{Value: v, Type: asType}
	}
	return res
}
