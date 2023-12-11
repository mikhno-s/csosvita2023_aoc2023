// 3 2
// 0 5
// -3 2
// 7 10
// 1 6
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Lines [][]int

func (l Lines) Len() int { return len(l) }
func (l Lines) Less(i, j int) bool {
	for x := range l[i] {
		if l[i][x] == l[j][x] {
			continue
		}
		return l[i][x] < l[j][x]
	}
	return false
}

func (l *Lines) Swap(i, j int) { (*l)[i], (*l)[j] = (*l)[j], (*l)[i] }

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	str := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	n, _ := strconv.Atoi(str[0])
	m, _ := strconv.Atoi(str[1])
	lines := make(Lines, n*2)
	for i := 0; i < len(lines); i += 2 {
		strLine := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		start, _ := strconv.Atoi(strLine[0])
		end, _ := strconv.Atoi(strLine[1])
		lines[i] = []int{start, 1}
		lines[i+1] = []int{end + 1, -1}
	}

	// nlogn
	sort.Sort(&lines)

	pointsStr := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	for i := 0; i < m; i++ {
		var total int
		p, _ := strconv.Atoi(pointsStr[i])
		for i := range lines {
			if p >= lines[i][0] {
				total += lines[i][1]
			}
		}
		fmt.Printf("%d ", total)
	}

}

func test(k int, a [][]int) (c int) {
	for _, el := range a {
		if el[0] <= k && k <= el[1] {
			c++
		}
	}
	return c
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
