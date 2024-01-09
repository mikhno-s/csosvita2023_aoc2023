package main

import (
	"fmt"
)

func gen(prefix []int, a []int, result *[][]int) {
	// check if object is ready
	if len(prefix) == len(a) {
		// need to create a new obj to append it to the result list
		r := []int{}
		for i, c := range prefix {
			if c == 1 {
				r = append(r, a[i])
			}
		}
		*result = append(*result, r)
		return
	}
	gen(append(prefix, 0), a, result)
	gen(append(prefix, 1), a, result)
	return
}

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	gen([]int{}, nums, &res)
	return res
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
