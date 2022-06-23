package main

import (
	"testing"
)

func TestAdder(t *testing.T) {
	arrayv := [4]int{1, 3, 1, 10}
	get := Adder(arrayv)
	want := 15
	if get != want {
		t.Errorf("Epected %d and Got %d", want, get)
	}

}
