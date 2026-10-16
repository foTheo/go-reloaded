package fonction

import "testing"

func TestPonctuation(t *testing.T) {
	tests := []struct{ name, input, want string }{
		{"simple", "over there ,and then BAMM !!", "over there, and then BAMM!!"},
		{"groupe", "I was thinking ... You were right", "I was thinking... You were right"},
		{"sans espace après", "hello,world", "hello, world"},
		{"décimal intact", "pi vaut 3.14", "pi vaut 3.14"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ponctuation(tt.input); got != tt.want {
				t.Errorf("Ponctuation(%q) = %q, attendu %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestApostrophes(t *testing.T) {
	tests := []struct{ name, input, want string }{
		{"paire", "describe me: ' awesome '", "describe me: 'awesome'"},
		{"deux paires", "' a ' and ' b '", "'a' and 'b'"},
		{"contraction intacte", "don't stop", "don't stop"},
		{"apostrophe seule", "say ' hi", "say ' hi"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Apostrophes(tt.input); got != tt.want {
				t.Errorf("Apostrophes(%q) = %q, attendu %q", tt.input, got, tt.want)
			}
		})
	}
}
