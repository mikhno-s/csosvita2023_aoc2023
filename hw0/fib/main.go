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

	t, err := strconv.Atoi(strings.TrimSpace(readLine(reader)))
	checkError(err)

	for tItr := 0; tItr < t; tItr++ {
		n, err := strconv.Atoi(strings.TrimSpace(readLine(reader)))
		checkError(err)
		fmt.Println(fibEvenSum(n))
	}
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

func fibEvenSum(max int) (sum int) {
	for n, n_last := 1, 1; n <= max; n, n_last = n+n_last, n {
		if n%2 == 0 {
			sum += n
		}
	}
	return sum
}
