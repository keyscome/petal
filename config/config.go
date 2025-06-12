// config/config.go
package config

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type RemoteTask struct {
	Name  string            `yaml:"name"`
	Hosts []string          `yaml:"hosts"`
	Env   map[string]string `yaml:"env,omitempty"`
	Cmd   string            `yaml:"cmd"`
}

type TaskFile struct {
	Env   map[string]string `yaml:"env,omitempty"`
	Tasks []RemoteTask      `yaml:"tasks"`
}

func LoadTaskFile(path string) (*TaskFile, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var tf TaskFile
	if err := yaml.Unmarshal(data, &tf); err != nil {
		return nil, err
	}
	return &tf, nil
}

func ValidateTaskFile(tf *TaskFile) error {
	seen := make(map[string]struct{})
	for _, t := range tf.Tasks {
		if t.Name == "" {
			return fmt.Errorf("task missing name")
		}
		if len(t.Hosts) == 0 {
			return fmt.Errorf("task '%s' missing hosts", t.Name)
		}
		if _, ok := seen[t.Name]; ok {
			return fmt.Errorf("duplicate task name: %s", t.Name)
		}
		seen[t.Name] = struct{}{}
	}
	return nil
}
