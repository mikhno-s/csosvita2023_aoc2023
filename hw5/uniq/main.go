package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	str := strings.Split(readLine(reader), "")
	fmt.Println(find(str))
	// fmt.Println(find([]string{"a", "a", "a"}))
	// fmt.Println(find([]string{"a", "b", "a", "b", "a"}))      // 2
	// fmt.Println(find([]string{"a", "b", "b", "c", "d", "e"})) // 4
	// fmt.Println(find([]string{"1", "2", "2", "3"}))           // 2

}

func find(a []string) int {
	var R, L int
	chars := make(map[string]bool)
	var currentLen, maxLen int
	good := func(c string) bool {
		return !chars[c]
	}
	add := func(s string) {
		currentLen++
		chars[s] = true
	}
	remove := func(s string) {
		currentLen--
		chars[s] = false
	}

	for R < len(a) {
		if good(a[R]) {
			if currentLen > maxLen {
				maxLen = currentLen
			}
		}
		if R < len(a) && good(a[R]) {
			add(a[R])
			R++
		} else {
			remove(a[L])
			L++
		}
	}

	return maxLen + 1
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
