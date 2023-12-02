package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type trip struct {
	m       int
	flavors []int
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	t, _ := strconv.Atoi(readLine(reader))

	trips := make([]*trip, t)
	for i := range trips {
		m, _ := strconv.Atoi(readLine(reader))
		n, _ := strconv.Atoi(readLine(reader))
		flavors := make([]int, n)
		flavorsStr := strings.Split(readLine(reader), " ")
		for f := range flavorsStr {
			num, _ := strconv.Atoi(flavorsStr[f])
			flavors[f] = num
		}
		trips[i] = &trip{m, flavors}
	}

	for i := range trips {
		fmt.Println(findPair(trips[i].m, trips[i].flavors))
	}

}

func findPair(n int, a []int) (int, int) {
	nums := make(map[int]int)
	min, max := -1, -1
	for i := range a {
		c := n - a[i]
		if _, ok := nums[c]; ok {
			min = nums[c] + 1
			max = i + 1
			break
		}
		nums[a[i]] = i
	}
	return min, max
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
