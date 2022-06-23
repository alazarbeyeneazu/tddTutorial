package main

import "testing"

func TestAddr(t *testing.T) {
	expect := 3
	sum := Addr(1, 2)
	if expect != sum {
		t.Errorf("%d what and get %d", expect, sum)
	}
}
