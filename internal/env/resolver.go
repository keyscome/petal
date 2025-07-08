package env

import (
	"os"
	"strings"
	"text/template"
)

type Resolver struct {
	merged map[string]EnvVar
}

func NewResolver(global, file, task map[string]EnvVar) *Resolver {
	merged := make(map[string]EnvVar)
	for _, src := range []map[string]EnvVar{global, file, task} {
		for k, v := range src {
			merged[k] = v
		}
	}
	return &Resolver{merged: merged}
}

func (r *Resolver) Flat() map[string]string {
	flat := make(map[string]string)
	for k, v := range r.merged {
		flat[k] = v.Value
	}
	return flat
}

func (r *Resolver) Render(input string) string {
	rendered := os.Expand(input, func(k string) string {
		if v, ok := r.merged[k]; ok {
			return v.Value
		}
		return ""
	})
	tmpl, err := template.New("tmpl").Parse(rendered)
	if err != nil {
		return rendered
	}
	var sb strings.Builder
	tmpl.Execute(&sb, r.Flat())
	return sb.String()
}

func (r *Resolver) CheckMissing(keys []string) []string {
	var missing []string
	for _, k := range keys {
		if _, ok := r.merged[k]; !ok {
			missing = append(missing, k)
		}
	}
	return missing
}
