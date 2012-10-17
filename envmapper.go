package envmapper

import "testing"

func Test_MapEnv(t *testing.T) {
	e := EnvMap{
		"SHELL": "/bin/bash",
		"TERM":  "screen-256color",
	}

	o := MapEnv([]string{
		"SHELL=/bin/bash",
		"TERM=screen-256color",
	})

	if len(o) != len(e) {
		t.Fatalf("%v != %v ", len(o), len(e))
	}

	for k, ev := range e {
		if ov, ok := o[k]; !ok || ov != ev {
			t.Fatalf("%v != %v", ev, ov)
		}
	}
}
