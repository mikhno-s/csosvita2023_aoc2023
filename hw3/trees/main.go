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

	s := strings.Split(readLine(reader), " ")
	n, _ := strconv.Atoi(s[0])
	k, _ := strconv.Atoi(s[1])

	treesStr := strings.Split(readLine(reader), " ")
	trees := make([]int, 0, n)
	for _, num := range treesStr {
		tmp, _ := strconv.Atoi(num)
		trees = append(trees, tmp)
	}
	fmt.Println(fast(k, trees))

	// fmt.Println(fast(3, []int{1, 2, 1, 3, 2})) // 2 4
	// fmt.Println(fast(4, []int{2, 4, 2, 3, 3, 1})) // 2 6
	// fmt.Println(fast(3, []int{1, 2, 2, 3, 2, 1})) // 4 6
}

func fast(k int, a []int) (int, int) {
	trees := make(map[int]int)
	var minR, L, R int
	minL := len(a)
	minLen := len(a)
	good := func() bool { return len(trees) >= k }
	add := func(n int) {
		trees[n] += 1
	}
	remove := func(n int) {
		trees[n] -= 1
		if trees[n] == 0 {
			delete(trees, n)
		}
	}

	for R < len(a) || L < len(a) {
		if R < len(a) && !good() {
			add(a[R])
			R++
		} else {
			if good() && R-L < minLen {
				minLen = R - L
				minL, minR = L, R
			}
			remove(a[L])
			L++
		}

	}

	return minL + 1, minR
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
