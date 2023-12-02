package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	X, Y int
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	n, _ := strconv.Atoi(readLine(reader))
	c := make([]coordinate, n)
	for i := 0; i < n; i++ {
		tmp := strings.Split(readLine(reader), " ")
		x, _ := strconv.Atoi(tmp[0])
		y, _ := strconv.Atoi(tmp[1])
		c[i] = coordinate{x, y}
	}

	m, _ := strconv.Atoi(readLine(reader))
	routes := make([][]int, m)
	for i := 0; i < m; i++ {
		tmp := strings.Split(readLine(reader), " ")
		start, _ := strconv.Atoi(tmp[0])
		finish, _ := strconv.Atoi(tmp[1])
		p := []int{start, finish}
		routes[i] = p
	}

	for _, r := range routes {
		fmt.Println(calculateTotalElevation(r[0], r[1], c))
	}

}

func calculateTotalElevation(start int, finish int, a []coordinate) (result int) {
	one := 1
	if start > finish {
		one = -1
	}

	start -= 1
	finish -= 1

	for ; start != finish; start += one {
		if a[start].Y < a[start+one].Y {
			result += a[start+one].Y - a[start].Y
		}
	}

	return result
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
