package main

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	cases := []struct {
		in     []int
		target int
		want   [][]int
	}{
		{[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
		{[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{[]int{2}, 1, [][]int{}},
		{[]int{1}, 2, [][]int{{1, 1}}},
		{[]int{7, 3, 2}, 7, [][]int{{2, 2, 3}, {7}}}, // input order must not matter
	}
	for _, c := range cases {
		if got := CombinationSum(c.in, c.target); !reflect.DeepEqual(got, c.want) {
			t.Errorf("CombinationSum(%v,%d) = %v, want %v", c.in, c.target, got, c.want)
		}
	}
}
