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

	x, _ := strconv.Atoi(s[0])
	y, _ := strconv.Atoi(s[1])
	z, _ := strconv.Atoi(s[2])
	w, _ := strconv.Atoi(s[3])

	var result int

	for i := 0; i <= 1000; i++ {
		for j := 0; j <= 1000; j++ {
			for k := 0; x*i+y*j+z*k <= w && k <= 1000; k++ {
				if x*i+y*j+z*k == w {
					result++
				}
			}
		}
	}
	fmt.Println(result)

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
