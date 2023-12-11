package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// 467..114..
// ...*......
// ..35..633.
// ......#...
// 617*......
// .....+.58.
// ..592.....
// ......755.
// ...$.*....
// .664.598..

func main() {
	fi, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	reader := bufio.NewReader(fi)

	engine := make([]string, 10)

	for i := 0; i < len(engine); i++ {
		engine[i] = readLine(reader)
	}

	var counter int

	for _, line := range engine {
		d := 0
		// var resultNumber int
		result := ""
		for i := range line {
			if strings.ContainsAny(string(line[i]), "*#+$") {
				result += string(line[i])
			} else if len(result) > 0 {
				// resultNumber, _ = strconv.Atoi(result)
				fmt.Println(result)
				// resultNumber = 0
				result = ""
			}
		}

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
