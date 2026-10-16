package fonction

import (
	"reflect"
	"testing"
)

func TestBin(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{"bin", []string{"101", "(bin)"}, []string{"5"}},
		{"non-bin", []string{"101", "hello"}, []string{"101", "hello"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bin(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bin(%v) = %v, attendu %v", tt.input, got, tt.want)
			}
		})
	}
}
