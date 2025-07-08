package env

import (
	"bufio"
	"os"
	"strings"
)

func LoadEnvFile(path string) (map[string]EnvVar, error) {
	vars := make(map[string]EnvVar)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || !strings.Contains(line, "=") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		vars[key] = EnvVar{Value: val, Type: ""}
	}
	return vars, scanner.Err()
}
