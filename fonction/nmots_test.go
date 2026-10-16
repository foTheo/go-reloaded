package fonction

import (
	"reflect"
	"strings"
	"testing"
)

func TestNmots(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"2)", 2},
		{"abc)", 0},
	}
	for _, tt := range tests {
		if got := Nmots(tt.input); got != tt.want {
			t.Errorf("Nmots(%q) = %d, attendu %d", tt.input, got, tt.want)
		}
	}
}

func TestTransformer(t *testing.T) {
	tests := []struct {
		name     string
		marqueur string
		fn       func(string) string
		input    []string
		want     []string
	}{
		{"low N=2", "low", strings.ToLower, []string{"HELLO", "WORLD", "(low,", "2)"}, []string{"hello", "world"}},
		{"up N=2", "up", strings.ToUpper, []string{"hello", "world", "(up,", "2)"}, []string{"HELLO", "WORLD"}},
		{"cap N=2", "cap", Capitaliser, []string{"hello", "world", "(cap,", "2)"}, []string{"Hello", "World"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Transformer(tt.input, tt.marqueur, tt.fn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transformer(%v, %q) = %v, attendu %v", tt.input, tt.marqueur, got, tt.want)
			}
		})
	}
}
