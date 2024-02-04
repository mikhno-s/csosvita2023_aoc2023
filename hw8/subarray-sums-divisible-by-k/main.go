package main

import (
	"fmt"
)

func subarraysDivByK(nums []int, k int) int {
	moduloMap := map[int]int{0: 1}
	var sum int
	var result int
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			sum += nums[i]%k + k
		} else {
			sum += nums[i]
		}
		reminder := sum % k
		if _, ok := moduloMap[reminder]; ok {
			result += moduloMap[reminder]
		}
		if _, ok := moduloMap[reminder]; !ok {
			moduloMap[reminder] = 1
		} else {
			moduloMap[reminder]++
		}
	}
	// fmt.Println(moduloMap)
	return result
}
func main() {
	fmt.Println(subarraysDivByK([]int{-1, 2, 9}, 2))
	fmt.Println(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5))
}
