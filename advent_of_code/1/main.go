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
	fi, err := os.Open("instructions.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	reader := bufio.NewReader(fi)

	lines := make([]string, 1000)
	for i := 0; i < len(lines); i++ {
		lines[i] = readLine(reader)
	}
	digitsMap := map[string]string{"1": "one", "2": "two", "3": "three", "4": "four", "5": "five", "6": "six", "7": "seven", "8": "eight", "9": "nine"}
	digits := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	var counter int

	for i := 0; i < len(lines); i++ {
		s := lines[i]
		var first, last string
		minIndex, maxIndex := len(s), -1
		for _, j := range digits {
			min := strings.Index(s, j)
			max := strings.LastIndex(s, j)
			if min != -1 && min < minIndex {
				minIndex = min
				first = j
			}
			if max != -1 && max > maxIndex {
				maxIndex = max
				last = j
			}
		}

		for k, v := range digitsMap {
			min := strings.Index(s, v)
			max := strings.LastIndex(s, v)
			if min != -1 && min < minIndex {
				minIndex = min
				first = k
			}
			if max != -1 && max > maxIndex {
				maxIndex = max
				last = k
			}
		}
		d, _ := strconv.Atoi(first + last)
		counter += d
	}

	fmt.Println(counter)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
