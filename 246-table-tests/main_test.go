package main

import "testing"

func TestMySum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	type seriesOfTests []test

	tests := seriesOfTests{
		test{[]int{2, 3, 4}, 9},
		test{[]int{2, 3, 5}, 10},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}

	x := mySum(2, 3, 4)

	if x != 9 {
		t.Error("Expected", 9, "Got", x)
	}
}
