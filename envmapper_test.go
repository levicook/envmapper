package envmapper

import "testing"

func Test_New(t *testing.T) {
	e := Map{
		"SHELL": "/bin/bash",
		"TERM":  "screen-256color",
		"FOO":   "bar",
	}

	o := New([]string{
		"SHELL=/bin/bash",
		"TERM=screen-256color",
		"FOO=bar",
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
