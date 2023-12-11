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

	// fi, err := os.Open("input")
	// if err != nil {
	// 	panic(err)
	// }

	// defer func() {
	// 	if err := fi.Close(); err != nil {
	// 		panic(err)
	// 	}
	// }()
	// reader := bufio.NewReader(fi)

	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	// read first array
	n1, _ := strconv.Atoi(readLine(reader))
	arrayStr := strings.Split(readLine(reader), " ")
	a1 := make([]int, n1)
	for i := range a1 {
		a1[i], _ = strconv.Atoi(arrayStr[i])
	}
	// read second array
	n2, _ := strconv.Atoi(readLine(reader))
	arrayStr = strings.Split(readLine(reader), " ")
	a2 := make([]int, n2)
	for i := range a2 {
		a2[i], _ = strconv.Atoi(arrayStr[i])
	}

	k, _ := strconv.Atoi(readLine(reader))
	mergeSort(a1, 0, len(a1))
	mergeSort(a2, 0, len(a2))

	max, min := a1[len(a1)-1]+a2[len(a2)-1]+1, a1[0]+a2[0]-1
	// fmt.Println(numOfPairsLessOrEqual(-30, a1, a2))

	for max-min > 1 {
		m := min + (max-min)/2
		if numOfPairsLessOrEqual(m, a1, a2) >= k {
			max = m
		} else {
			min = m
		}
	}

	fmt.Println(max)

}

// -28
// -24
// -22
// -20
// -15
// -11
// -9
// -7
// 4
// 5
// 6
// 8
// 9
// 10
// 10
// 10
// 11
// 12
// 12
// 13
// 14
// 23
// 42
// 43
// 44
// 82
// 86
// 88
// 90
// 120

func numOfPairsLessOrEqual(nk int, a []int, b []int) int {
	c := 0

	for i := range a {
		R := len(b)
		L := -1

		for R-L > 1 {
			m := L + (R-L)/2
			res := a[i] + b[m]
			if res > nk {
				R = m
			} else {
				L = m
			}

		}

		c += R
	}

	return c
}

func mergeSort(a []int, l int, r int) {
	if r-l < 2 {
		return
	}
	m := (l + r) / 2
	mergeSort(a, l, m)
	mergeSort(a, m, r)
	merge(a, l, m, r)
	return
}

func merge(a []int, l int, m int, r int) {
	aux := make([]int, m-l)

	for i := range aux {
		aux[i] = a[l+i]
	}

	i, j, k := 0, m, l
	for i < len(aux) || j < r {
		if j == r || (i < len(aux) && aux[i] < a[j]) {
			a[k] = aux[i]
			i++
		} else {
			a[k] = a[j]
			j++
		}
		k++
	}

	return
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
