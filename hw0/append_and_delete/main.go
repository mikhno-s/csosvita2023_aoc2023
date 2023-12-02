package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'appendAndDelete' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. STRING t
 *  3. INTEGER k
 */

func appendAndDelete(s string, t string, k int) string {
	var minSteps int
	lenDiff := len(s) - len(t)

	if lenDiff < 0 {
		lenDiff *= -1
	}

	// Calculate the difference
	minSteps += lenDiff
	// search for different symbols in arrays to calculate the total delete and append operations needed
	for i := range s {
		if i < len(t) {
			if s[i] != t[i] {
				minSteps += 2
				// The rest of the string must be deleted so we dont check the diff
				if len(t) > len(s) {
					minSteps += 2 * (len(s) - i - 1)
				} else {
					minSteps += 2 * (len(t) - i - 1)
				}
				break
			}
		}
	}

	// if minimum amount of steps need to be done to transform string is more than k then string cannot be transformed in k amount of steps
	// minSteps - lower bound
	if minSteps > k {
		return "No"
	}

	// We have ariphmetic progression [minSteps, minSteps+2, minSteps+4, ..., len(t)*2+(len(s)-len(t)] that represents the number of operations
	// possible values
	for i := 1; i <= len(t); i++ {
		if k == i*2+(len(s)-len(t)) {
			return "Yes"
		}
	}

	// upper bound
	if k >= len(t)*2+(len(s)-len(t)) {
		return "Yes"
	}

	return "No"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	s := readLine(reader)
	t := readLine(reader)

	k, err := strconv.Atoi(strings.TrimSpace(readLine(reader)))
	checkError(err)
	// result := appendAndDelete("11", "111", 7)
	result := appendAndDelete(s, t, k)

	fmt.Printf("%s\n", result)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
