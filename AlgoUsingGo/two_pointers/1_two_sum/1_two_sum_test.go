package main

import (
	"fmt"
	"testing"
)

func twoSumTest(t *testing.T) {

	test := []struct {
		data   []int
		target int
		ans    []int
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
	}
	for _, t := range test {
		actual := towSum(t.data, t.target)
		for i := range actual {
			if actual[i] != t.ans[i] {
				fmt.Println("it is wrong")
			}
		}
	}
}
