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
	N, _ := strconv.Atoi(s[0])
	A, _ := strconv.ParseFloat(s[1], 64)

	bad, good := 0.0, A*A

	cand := bad + (good-bad)/2
	for good != cand && bad != cand {
		m := findLastOrLowerTheFloor(A, cand, N-1)
		// fmt.Println(cand, m)
		if m >= 0.00000000001 {
			good = cand
		} else {
			bad = cand
		}
		cand = bad + (good-bad)/2
	}
	fmt.Printf("%f\n", findLastOrLowerTheFloor(A, good, N-1))

}

func findLastOrLowerTheFloor(current float64, next float64, lightNum int) float64 {
	last := next
	for ; lightNum > 1; lightNum-- {
		last = 2*next - current + 2
		if last <= 0 {
			break
		}
		// fmt.Println(last, lightNum)
		current = next
		next = last
	}

	return last
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
