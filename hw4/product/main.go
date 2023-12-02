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

	ns := strings.Split(readLine(reader), " ")
	n := make([]int, len(ns))
	for i := range ns {
		num, _ := strconv.Atoi(ns[i])
		n[i] = num
	}
	fmt.Println(getMaxProduct(n))
}
func getMaxProduct(a []int) (int, int) {
	maxA, maxB := -1, -1
	minA, minB := 1, 1

	for i := range a {
		if a[i] > maxA {
			maxA = a[i]
		}
		if maxA > maxB {
			maxA, maxB = maxB, maxA
		}

		if a[i] < minA {
			minA = a[i]
		}

		if minA < minB {
			minA, minB = minB, minA
		}
	}

	if minA*minB > maxA*maxB {
		maxA, maxB = minA, minB
	}
	if maxA < maxB {
		return maxA, maxB
	}
	return maxB, maxA
}

func getMaxProductSlow(a []int) (int, int) {
	maxA, maxB := -1, -1
	maxProduct := 0
	for i := 0; i < len(a); i++ {

		for j := i + 1; j < len(a); j++ {
			currentProduct := a[i] * a[j]
			if maxProduct < currentProduct {
				maxA = a[i]
				maxB = a[j]
				maxProduct = currentProduct
			}
		}
	}
	if maxA < maxB {
		return maxA, maxB
	}
	return maxB, maxA
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
