package main

import (
	"fmt"
	"testing"
)

func ExampleAddr() {
	sum := Addr(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAddr(t *testing.T) {
	expect := 3
	sum := Addr(1, 2)
	if expect != sum {
		t.Errorf("%d what and get %d", expect, sum)
	}
	expect2 := 3
	sum2 := Addr(1, 2)
	if expect2 != sum2 {
		t.Errorf("%d what and get %d", expect2, sum2)
	}
}
