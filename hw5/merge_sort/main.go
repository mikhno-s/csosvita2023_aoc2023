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

	numsStr := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	nums := make([]int, len(numsStr))
	for i := range numsStr {
		num, _ := strconv.Atoi(numsStr[i])
		nums[i] = num
	}
	mergeSort(nums, 0, len(nums))
	for _, el := range nums {
		fmt.Printf("%d ", el)
	}

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
