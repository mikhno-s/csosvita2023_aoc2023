package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	// file, _ := os.Open("input.txt")
	// reader := bufio.NewReaderSize(file, 16*1024*1024)
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	input1 := strings.Split(readLine(reader), " ")
	n, _ := strconv.Atoi(input1[0])
	Zmax, _ := strconv.Atoi(input1[1])
	Omax, _ := strconv.Atoi(input1[2])

	S := [][]int{}

	for i := 0; i < n; i++ {
		line := readLine(reader)
		S = append(S, []int{strings.Count(line, "0"), strings.Count(line, "1")})
	}

	// create a 3-dimension DP table that stores in the [i][z][o] for the iTh item with the max z 0 and max o 1
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, Zmax+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, Omax+1)
		}
	}

	for i := 1; i < len(dp); i++ {
		Zi := S[i-1][0]
		Oi := S[i-1][1]
		for z := 0; z <= Zmax; z++ {
			for o := 0; o <= Omax; o++ {
				if z >= Zi && o >= Oi {
					prev := float64(dp[i-1][z][o])
					prevSmaller := float64(dp[i-1][z-Zi][o-Oi])
					dp[i][z][o] = int(math.Max(prev, prevSmaller+1))
				} else {
					dp[i][z][o] = dp[i-1][z][o]
				}
			}
		}
	}

	fmt.Println(dp[n][Zmax][Omax])
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
