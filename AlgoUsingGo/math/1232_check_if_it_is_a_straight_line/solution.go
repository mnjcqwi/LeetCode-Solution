package main

import "strconv"

func checkStraightLine(coordinates [][]int) bool {
	slop := getSlop(coordinates[0], coordinates[1])
	for i := 2; i < len(coordinates); i++ {
		if slop != getSlop(coordinates[i], coordinates[i-1]) {
			return false
		}
	}
	return true

}

func getSlop(curt, prev []int) string {
	dy := curt[1] - prev[1]
	dx := curt[0] - prev[0]
	if dy == 0 {
		return "/"
	}
	gc := gcd(dy, dx)
	return strconv.Itoa(dy/gc) + "/" + strconv.Itoa(dx/gc)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

