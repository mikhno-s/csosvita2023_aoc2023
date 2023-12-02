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
	numArray := strings.Split(strings.TrimSpace(readLine(reader)), "")

	var count int
	// I use map to create a tree of numbers
	total := make(map[string]map[string]map[string]bool)
	for i := 0; i < len(numArray); i++ {

		// ignore numbers that start from zero
		if numArray[i] == "0" {
			continue
		}
		// check first number
		if _, ok := total[numArray[i]]; !ok {
			total[numArray[i]] = make(map[string]map[string]bool)
		}

		// go all possible second numbers
		for j := i + 1; j < len(numArray); j++ {

			// check second number
			if _, ok := total[numArray[i]][numArray[j]]; !ok {
				total[numArray[i]][numArray[j]] = make(map[string]bool)
			}

			// go over all possible third numbers
			for k := j + 1; k < len(numArray); k++ {
				// if there is no key than create a new record
				if _, ok := total[numArray[i]][numArray[j]][numArray[k]]; !ok {
					total[numArray[i]][numArray[j]][numArray[k]] = true
					count++
				}
			}
		}
	}
	fmt.Println(count)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

// [1][2][3] = true
// [1][2][1] = true
// ...
