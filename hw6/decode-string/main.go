package main

import (
	"fmt"
	"strconv"
	"strings"
)

func decodeString(s string) string {
	result := ""
	// resultStack := &EndodedStringStack{}
	last := len(s)
	// go over each symbol
	for L := 0; L < last; L++ {
		nStr := ""
		var start int
		// find the encoded substing and find start of the actual substring
		if strings.Contains("1234567890", s[L:L+1]) {
			for L < last && strings.Contains("1234567890", s[L:L+1]) {
				nStr += s[L : L+1]
				L++
			}
			// skip first opening breackets
			start = L + 1

			j := 1 // amount of breakets
			substringEnd := start
			// find the matched bracket and the end index of substring
			for substringEnd < last && j > 0 {
				if s[substringEnd:substringEnd+1] == "[" {
					j++
				}
				if s[substringEnd:substringEnd+1] == "]" {
					j--
				}
				if j == 0 {
					break
				}
				substringEnd++
			}
			n, _ := strconv.Atoi(nStr)
			for i := 0; i < n-1; i++ {
				result += decodeString(s[start:substringEnd])
			}
		} else {
			if s[L:L+1] != "]" {
				result += s[L : L+1]
			}
		}
	}
	return result
}

func printNtimes(n int, s string) string {
	var res string
	for i := 0; i < n; i++ {
		res += s
	}
	return res
}

func main() {

	s := "3[q2[a]]2[ef]"
	fmt.Println(decodeString(s))
}
