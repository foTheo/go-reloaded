package fonction

import (
	"reflect"
	"testing"
)

func TestUp(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{"up", []string{"Hello", "world", "(up)"}, []string{"Hello", "WORLD"}},
		{"non-up", []string{"Hello", "world"}, []string{"Hello", "world"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Up(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Up(%v) = %v, attendu %v", tt.input, got, tt.want)
			}
		})
	}
}
