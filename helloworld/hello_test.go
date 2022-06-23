package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello, Alazar"
	got := Hello("Alazar")
	if want != got {
		t.Errorf("Expected %s Output %s ", want, got)
	}
}
