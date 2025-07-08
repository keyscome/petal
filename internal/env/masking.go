package env

func (r *Resolver) GetMaskedEnv() map[string]string {
	masked := make(map[string]string)
	for k, v := range r.merged {
		if v.Type == "secret" {
			masked[k] = "******"
		} else {
			masked[k] = v.Value
		}
	}
	return masked
}
