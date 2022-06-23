package main

import (
	"reflect"
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

func TestSum(t *testing.T) {
	testCase := []int{1, 2, 3, 4, 5}
	t.Run("Printer out the value ", func(t *testing.T) {
		want := 15
		got := Sum(testCase)
		if want != got {
			t.Errorf("What i want %d and What i got %d ", want, got)
		}
	})

}
func TestSumAll(t *testing.T) {
	testcase1, testcase2 := []int{1, 2}, []int{5, 4}
	want := []int{3, 9}
	got := SumAll(testcase1, testcase2)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected %#v and Got %#v", want, got)
	}
}
func TestSumAllTail(t *testing.T) {
	test1, test2 := []int{1, 2}, []int{3, 4}
	expected := []int{2, 4}
	result := SumAllTail(test1, test2)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\nwhat i Expect %#v \nWhat i have got %#v", expected, result)
	}
}
