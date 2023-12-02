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
	// n, _ := strconv.Atoi(s[0])
	k, _ := strconv.Atoi(s[1])

	str := strings.Split(readLine(reader), "")
	fmt.Println(fast(k, str))
	// fmt.Println(fast(1, []string{"a", "b", "b"}))                // 2 1
	// fmt.Println(fast(2, []string{"a", "b", "a", "b", "a"})) // 4 1
	// fmt.Println(fast(1, []string{"a", "b", "b", "c", "d", "e"})) // 4 3
	// fmt.Println(fast(3, []string{"1", "2", "2", "3"}))           // 4 1

}

func fast(k int, a []string) (int, int) {
	var index, R, L int
	chars := make(map[string]int)
	var currentLen, maxLen int
	good := func() bool {
		// fmt.Print(chars)
		for _, i := range chars {
			if i > k {
				return false
			}
		}
		return true
	}
	add := func(s string) {
		currentLen++
		chars[s] += 1
	}
	remove := func(s string) {
		currentLen--
		chars[s] -= 1
	}

	for R < len(a) || L < len(a) {
		if good() {
			if currentLen > maxLen {
				maxLen = currentLen
				index = L
			}
		}
		if R < len(a) && good() {
			add(a[R])
			R++
		} else {
			remove(a[L])
			L++

		}

	}

	return maxLen, index + 1
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
