package main

func towSum(nums []int, target int) []int {
	rest := make(map[int]int)
	for i, v := range nums {
		if j, ok := rest[target-v]; !ok {
			rest[v] = j
		} else {

			return []int{i, j}
		}
	}
	return make([]int, 0)
}

