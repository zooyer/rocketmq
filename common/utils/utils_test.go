package utils

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestIsBlank(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"Hello ": {input: "Hello ", want: false},
		" Hello": {input: " Hello", want: false},
		"He llo ": {input: "He llo ", want: false},
		"  ": {input: "  ", want: true},
		"	": {input: "	", want: true},
		" \n": {input: " \n", want: true},
		"\r ": {input: "\r ", want: true},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := IsBlank(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
