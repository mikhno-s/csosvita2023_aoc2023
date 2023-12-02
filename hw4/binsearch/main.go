package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	numlen, _ := strconv.Atoi(readLine(reader))
	numsStr := strings.Split(readLine(reader), " ")
	nums := make([]int, numlen)
	for i := range nums {
		num, _ := strconv.Atoi(numsStr[i])
		nums[i] = num
	}
	n, _ := strconv.Atoi(readLine(reader))
	fmt.Println(bs(n, nums))

}

// [L:R)
func bs(n int, a []int) (int, int, int) {
	L, R := -1, len(a)
	first := -1
	for R-L > 1 {
		m := L + (R-L)/2
		if a[m] >= n {
			R = m
		} else {
			L = m
		}
	}

	if R < len(a) && a[R] == n {
		first = R
	} else {
		return first, -1, 0
	}

	L = R // set left bound
	R = len(a)

	for R-L > 1 {
		m := L + ((R - L) / 2)
		if a[m] > n {
			R = m
		} else {
			L = m
		}
	}

	return first, L, L - first + 1
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
