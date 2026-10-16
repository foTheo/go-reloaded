package fonction

import (
	"reflect"
	"testing"
)

func TestHex(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{"hex", []string{"1A", "(hex)"}, []string{"26"}},
		{"non-hex", []string{"1A", "hello"}, []string{"1A", "hello"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hex(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hex(%v) = %v, attendu %v", tt.input, got, tt.want)
			}
		})
	}
}
