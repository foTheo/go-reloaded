package main

import (
	"reflect"
	"testing"
)

func TestCap(t *testing.T) {
	mots := []string{"hELLO", "(cap)", "école", "(cap)"}
	want := []string{"HELLO", "École"}

	if got := Cap(mots); !reflect.DeepEqual(got, want) {
		t.Fatalf("Cap(%v) = %v, want %v", mots, got, want)
	}
}
