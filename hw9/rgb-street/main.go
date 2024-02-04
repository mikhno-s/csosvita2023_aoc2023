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

	file, _ := os.Open("input.txt")
	reader := bufio.NewReaderSize(file, 16*1024*1024)
	// reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	n, _ := strconv.Atoi(readLine(reader))

	input := []map[string]int{}

	for i := 0; i < n; i++ {
		s := strings.Split(strings.TrimSpace(readLine(reader)), " ")
		R, _ := strconv.Atoi(s[0])
		G, _ := strconv.Atoi(s[1])
		B, _ := strconv.Atoi(s[2])
		input = append(input, map[string]int{"R": R, "G": G, "B": B})
	}

	prefixSum := map[string]int{}

	// calculate default values
	for c := range input[0] {
		prefixSum[c] = input[0][c]
	}
	minSum := math.MaxInt
	// go over the all houses
	for i := 1; i < n; i++ {
		prefixSumIntermidiate := map[string]int{}
		// append the current house's prices to each sum of prices calculated before
		for prefix := range prefixSum {
			// check each color
			for current := range input[i] {
				lastColor := prefix[len(prefix)-1:]
				// calculate the prefix sum (cache) if last colors are different or current total sum is less then already calculated min sum
				if current != lastColor && len(prefix+current) <= n && prefixSum[prefix]+input[i][current] < minSum {
					prefixSumIntermidiate[prefix+current] += prefixSum[prefix] + input[i][current]
					// check the minimum price if it's end of the street
					if len(prefix+current) == n && prefixSumIntermidiate[prefix+current] < minSum {
						minSum = prefixSumIntermidiate[prefix+current]
					}
				}
			}
			// delete prev prefixes
			delete(prefixSum, prefix)
		}
		// copy to prefix sum (cache)
		for prefix := range prefixSumIntermidiate {
			prefixSum[prefix] = prefixSumIntermidiate[prefix]
		}
	}
	fmt.Println(minSum)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
