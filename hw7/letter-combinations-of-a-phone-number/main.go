package main

import (
	"fmt"
)

var keyboard = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func gen(prefix string, a []string, n int) []string {
	// check if object is ready
	if len(prefix) >= n {
		return []string{prefix}
	}

	result := []string{}
	for _, c := range a[0] {
		// without backtracking?
		result = append(result, gen(prefix+string(c), a[1:], n)...)
	}
	return result
}

func letterCombinations(digits string) []string {

	a := make([]string, 0)
	for _, d := range digits {
		a = append(a, keyboard[string(d)])
	}
	if len(digits) == 0 {
		return []string{}
	}
	return gen("", a, len(digits))
}

func main() {
	fmt.Println(letterCombinations("222"))

}
