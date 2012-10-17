package envmapper

import "strings"

type Map map[string]string

func New(env []string) Map {
	m := make(Map, cap(env))
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
