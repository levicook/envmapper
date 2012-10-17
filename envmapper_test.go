package envmapper

import "strings"

type EnvMap map[string]string

func MapEnv(env []string) EnvMap {
	m := make(EnvMap, cap(env))
	for _, s := range env {
		sp := strings.Split(s, "=")
		if len(sp) == 2 {
			k := sp[0]
			v := sp[1]
			m[k] = v
		}
	}
	return m
}
