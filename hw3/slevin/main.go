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
	n, _ := strconv.Atoi(s[0])
	m, _ := strconv.Atoi(s[1])

	regPlates := make([]int, 0, n)

	regPlatesStr := strings.Split(readLine(reader), " ")
	for _, num := range regPlatesStr {
		tmp, _ := strconv.Atoi(num)
		regPlates = append(regPlates, tmp)
	}
	fmt.Println(fast(m, regPlates))

	// var c int
	// for {

	// 	regPlates := make([]int, rand.Int31n(10))

	// 	for i := range regPlates {
	// 		regPlates[i] = rand.Intn(998) + 1
	// 	}

	// 	k := int(rand.Int31n(100) + 1)

	// 	s := slow(k, regPlates)
	// 	f := fast(k, regPlates)
	// 	c++

	// 	if s != f {
	// 		fmt.Println(k)
	// 		fmt.Println(s, f)
	// 		fmt.Println(regPlates)
	// 		break
	// 	}
	// 	if c%10 == 0 {
	// 		fmt.Println(c)
	// 	}

	// }

}

func fast(m int, regPlates []int) (result int) {
	var l, sum int
	for r := range regPlates {
		sum += regPlates[r]
		for sum > m {
			sum -= regPlates[l]
			l++
		}
		if sum == m {
			result++
		}
	}

	return result
}

func slow(m int, regPlates []int) (result int) {
	for i := range regPlates {
		var sum int
		sum = regPlates[i]
		if sum == m {
			result++
		}
		for j := i + 1; j < len(regPlates); j++ {
			sum += regPlates[j]
			if sum == m {
				result++
				break
			}
			if sum > m {
				break
			}
		}
	}
	return
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
