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
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	str := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	n, _ := strconv.Atoi(str[0])
	m, _ := strconv.Atoi(str[1])
	lines := make([][]int, n)
	for i := range lines {
		lines[i] = make([]int, 2)
		strLine := strings.Split(strings.TrimSpace(readLine(reader)), " ")
		lines[i][0], _ = strconv.Atoi(strLine[0])
		lines[i][1], _ = strconv.Atoi(strLine[1])
		// first element must be a start of line
		if lines[i][1] < lines[i][0] {
			lines[i][0], lines[i][1] = lines[i][1], lines[i][0]
		}
	}

	pointsStr := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	for i := 0; i < m; i++ {
		p, _ := strconv.Atoi(pointsStr[i])
		fmt.Printf("%d ", slow(p, lines))
	}
	fmt.Println()

	// fmt.Println(lines, points)

}

func slow(k int, a [][]int) (c int) {
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
