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

	programNum, _ := strconv.Atoi(s[0])
	coursesNum, _ := strconv.Atoi(s[1])

	programs := make([][]string, programNum)

	// getting input
	for p := range programs {
		programs[p] = strings.Split(readLine(reader), " ")
	}

	var maxLen, minLen, A, B int
	minLen = programNum + 1
	for i := 0; i < coursesNum; i++ {
		for j := i + 1; j < coursesNum; j++ {
			var ab, anotb, notab, notanotb int
			for k := 0; k < programNum; k++ {
				if minLen == 1 {
					break
				}
				if programs[k][i] == "1" && programs[k][j] == "1" {
					ab++
					if ab >= maxLen {
						maxLen = ab
					}
				} else if programs[k][i] == "1" {
					anotb++
					if anotb >= maxLen {
						maxLen = anotb
					}
				} else if programs[k][j] == "1" {
					notab++
					if notab >= maxLen {
						maxLen = notab
					}
				} else if programs[k][i] == "0" && programs[k][j] == "0" {
					notanotb++
					if notanotb >= maxLen {
						maxLen = notanotb
					}
				}

			}

			if maxLen < minLen {
				minLen = maxLen
				A = i + 1
				B = j + 1
			}
			maxLen = 0
		}
	}

	fmt.Println(minLen)
	fmt.Println(A, B)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
