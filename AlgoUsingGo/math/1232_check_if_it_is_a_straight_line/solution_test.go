package main

import (
	"fmt"
	"testing"
)

func TestCheckStraightLine(t *testing.T) {
	test := []struct {
		data   [][]int
		expect bool
	}{
		{
			[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}},
			true,
		},
		{
			[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {7, 7}},
			false,
		},
	}
	for _, tt := range test {
		if ans := checkStraightLine(tt.data); ans != tt.expect {
			fmt.Println("Solution is Wrong")
		}
	}

}
