package fonction

import (
	"reflect"
	"testing"
)

func TestCap(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{"cap", []string{"hello", "world", "(cap)"}, []string{"hello", "World"}},
		{"non-cap", []string{"hello", "world"}, []string{"hello", "world"}},
		{"cour-test", []string{"Welcome", "to", "the", "Brooklyn", "bridge", "(cap)"}, []string{"Welcome", "to", "the", "Brooklyn", "Bridge"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cap(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cap(%v) = %v, attendu %v", tt.input, got, tt.want)
			}
		})
	}
}
