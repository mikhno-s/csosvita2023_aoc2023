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

	n, _ := strconv.Atoi(readLine(reader))

	a := make([]string, n)
	for i := range a {
		s := strings.TrimSpace(readLine(reader))
		a[i] = s
	}
	mergeSort(a, 0, len(a))
	for _, el := range a {
		fmt.Printf("%s\n", el)
	}

}

func mergeSort(a []string, l int, r int) {
	if r-l < 2 {
		return
	}
	m := (l + r) / 2
	mergeSort(a, l, m)
	mergeSort(a, m, r)
	merge(a, l, m, r)
	return
}

func merge(a []string, l int, m int, r int) {
	aux := make([]string, m-l)

	for i := range aux {
		aux[i] = a[l+i]
	}

	i, j, k := 0, m, l
	for i < len(aux) || j < r {
		if j == r || (i < len(aux) && isBigger(a[j], aux[i])) {
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

func isBigger(s1 string, s2 string) bool {
	if len(s1) == len(s2) {
		return s1 > s2
	}
	return len(s1) > len(s2)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
