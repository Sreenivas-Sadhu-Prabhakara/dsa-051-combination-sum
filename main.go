package main

import (
	"fmt"
	"sort"
)

// CombinationSum returns all unique combinations of candidates (each usable
// unlimited times) that sum to target. Sorting enables the early cutoff.
func CombinationSum(candidates []int, target int) [][]int {
	a := append([]int{}, candidates...)
	sort.Ints(a)
	out := [][]int{}
	cur := []int{}
	var walk func(start, remaining int)
	walk = func(start, remaining int) {
		if remaining == 0 {
			out = append(out, append([]int{}, cur...))
			return
		}
		for i := start; i < len(a) && a[i] <= remaining; i++ {
			cur = append(cur, a[i])
			walk(i, remaining-a[i]) // i again, not i+1: reuse allowed
			cur = cur[:len(cur)-1]
		}
	}
	walk(0, target)
	return out
}

func main() {
	fmt.Println(CombinationSum([]int{2, 3, 6, 7}, 7))
}
