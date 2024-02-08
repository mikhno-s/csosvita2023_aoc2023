package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// file, _ := os.Open("input.txt")
	// reader := bufio.NewReaderSize(file, 16*1024*1024)
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	n, _ := strconv.Atoi(readLine(reader))

	player, computer := []int{}, []int{}

	s := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	for i := 0; i < n; i++ {
		u, _ := strconv.Atoi(s[i])
		player = append(player, u)
	}

	s = strings.Split(strings.TrimSpace(readLine(reader)), " ")
	for i := 0; i < n; i++ {
		u, _ := strconv.Atoi(s[i])
		computer = append(computer, u)
	}

	sort.Ints(player)
	sort.Ints(computer)

	P, C := n-1, n-1
	var total int
	// two pointers
	for P >= 0 && C >= 0 {
		// fmt.Printf("[%d]%d / [%d]%d: ", C, computer[C], P, player[P])
		if computer[C] < player[P] {
			// fmt.Println("WIN")
			total += player[P]
			C--
			P--
		} else {
			// fmt.Println("LOSE")
			C--
		}
	}

	fmt.Println(total)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
