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
	inputArray := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	n, _ := strconv.Atoi(inputArray[0])
	nAB, _ := strconv.Atoi(inputArray[1])
	nBC, _ := strconv.Atoi(inputArray[2])
	nAC, _ := strconv.Atoi(inputArray[3])

	// checking the kAC kC kBC
	sumA := n - nBC
	sumB := n - nAC
	sumC := n - nAB

	innerN := sumA + sumB + sumC - n
	possibleIntermediateVariants := findTriplets(innerN)

	for _, i := range possibleIntermediateVariants {
		kAB := i[0]
		kBC := i[1]
		kAC := i[2]
		kA := sumA - kAB - kAC
		kB := sumB - kAB - kBC
		kC := sumC - kAC - kBC

		if kA >= 0 && kB >= 0 && kC >= 0 {
			if kA+kAB+kB == nAB && kB+kBC+kC == nBC && kC+kAC+kA == nAC {
				fmt.Println("YES")
				fmt.Println(kA, kAB, kB, kBC, kC, kAC)
				os.Exit(0)
			}
		}
	}
	fmt.Println("NO")
}

func findTriplets(target int) [][]int {
	var result [][]int

	for i := 0; i <= target; i++ {
		for j := 0; j <= target; j++ {
			for k := 0; k <= target; k++ {
				if i+j+k == target {
					result = append(result, []int{i, j, k})
				}
			}
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
