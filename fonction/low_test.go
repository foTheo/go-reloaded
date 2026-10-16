package fonction

import (
	"reflect"
	"testing"
)

func TestLow(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{"low", []string{"Hello", "WORLD", "(low)"}, []string{"Hello", "world"}},
		{"non-low", []string{"Hello", "WORLD"}, []string{"Hello", "WORLD"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Low(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Low(%v) = %v, attendu %v", tt.input, got, tt.want)
			}
		})
	}
}
