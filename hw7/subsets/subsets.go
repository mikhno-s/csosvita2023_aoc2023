package main

import (
	"fmt"
)

const n = 1

func genSubsets(prefix []int, a []int) {
	if len(prefix) == n {
		for i, c := range prefix {
			if c == 1 {
				fmt.Print(a[i])
			}
		}
		fmt.Println()
		return
	}
	genSubsets(append(prefix, 0), a)
	genSubsets(append(prefix, 1), a)
}

func subsets(nums []int) {
	genSubsets([]int{}, nums)
}
