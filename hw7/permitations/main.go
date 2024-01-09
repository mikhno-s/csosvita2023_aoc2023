package main

import (
	"fmt"
)

func gen(prefix []int, a []int, result *[][]int, exists map[int]bool) {
	// check if object is ready
	if len(prefix) == len(a) {
		// need to create a new obj to append it to the result list
		tmp := make([]int, len(prefix))
		copy(tmp, prefix)
		*result = append(*result, tmp)
		return
	}
	for _, c := range a {
		// check if object is valid
		if !exists[c] {
			exists[c] = true                          // choose
			gen(append(prefix, c), a, result, exists) // explore
			exists[c] = false                         // unchoose
		}
	}
	return
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	gen([]int{}, nums, &res, map[int]bool{})
	return res
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
