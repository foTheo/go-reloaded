package fonction

import (
	"reflect"
	"strings"
	"testing"
)

func TestAn(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"voyelle", "a apple", "an apple"},
		{"h", "a honest person", "an honest person"},
		{"majuscule", "A elephant", "An elephant"},
		{"consonne", "a banana", "a banana"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := An(strings.Fields(tt.input))
			gotPhrase := strings.Join(got, " ")

			if !reflect.DeepEqual(gotPhrase, tt.want) {
				t.Errorf("An(%q) = %q, attendu %q", tt.input, gotPhrase, tt.want)
			}
		})
	}
}
