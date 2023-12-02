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
	s := readLine(reader)

	n, _ := strconv.Atoi(s)
counting:
	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			for c := 0; c < 100; c++ {
				for d := 0; d < 100; d++ {
					if a*a+b*b+c*c+d*d == n {
						if a != 0 {
							fmt.Printf("%d ", a)
						}
						if b != 0 {
							fmt.Printf("%d ", b)
						}
						if c != 0 {
							fmt.Printf("%d ", c)
						}
						if d != 0 {
							fmt.Printf("%d", d)
						}
						break counting
					}
				}
			}
		}
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
