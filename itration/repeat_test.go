package main

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	sum := Repeat("a", 8)
	fmt.Println(sum)
}

func TestRepeat(t *testing.T) {
	expect := "aaaaaaaa"
	got := Repeat("a", 8)
	if expect != got {
		t.Errorf("What Expected %s and What i Got %s", expect, got)
	}
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}
