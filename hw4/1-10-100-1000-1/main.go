package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 12 4  7   11   16
// 1101001000100001
//
//	3  6
//
// 1101001000100001
func main() {
	// same time can be done with hashmap that reduses algo time complexity to O(1)
	max := 65536
	prefixSum := make([]int, max)
	prefixSum[0] = 1
	for i := 1; i < max; i++ {
		prefixSum[i] = prefixSum[i-1] + i
	}
	fmt.Println(prefixSum)
	// fmt.Println("ready")
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	numlen, _ := strconv.Atoi(readLine(reader))
	nums := make([]int, numlen)
	for i := range nums {
		num, _ := strconv.Atoi(readLine(reader))
		if num > max {
			max = num
		}
		nums[i] = num
	}

	for i, n := range nums {
		fmt.Printf("%d", showPos(n, prefixSum))
		if i < len(nums)-1 {
			fmt.Printf(" ")
		}
	}
}

func showPos(n int, prefixSum []int) int {
	L, R := -1, len(prefixSum)
	result := -1
	for R-L > 1 {
		m := L + ((R - L) / 2)

		if prefixSum[m] >= n {
			R = m
			result = prefixSum[m]
		} else {
			L = m
		}
	}
	if result == n {
		return 1
	} else {
		return 0
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
