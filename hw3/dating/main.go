package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	s := strings.Split(readLine(reader), " ")
	n, _ := strconv.Atoi(s[0])
	r, _ := strconv.Atoi(s[1])

	distsStr := strings.Split(readLine(reader), " ")
	dists := make([]int, 0, n)
	for _, num := range distsStr {
		tmp, _ := strconv.Atoi(num)
		dists = append(dists, tmp)
	}

	fmt.Println(fast(r, dists))

	// var c int
	// for {

	// 	a := make([]int, rand.Int31n(10))

	// 	for i := range a {
	// 		a[i] = rand.Intn(300000) + 1
	// 	}

	// 	sort.Ints(a)

	// 	r := int(rand.Intn(999999999)) + 1

	// 	s := slow(r, a)
	// 	f := fast(r, a)
	// 	c++

	// 	if s != f {
	// 		fmt.Println(r, s, "!=", f)
	// 		fmt.Println(a)
	// 		break
	// 	}
	// 	if c%10 == 0 {
	// 		fmt.Println(c)
	// 	}

	// }

}

func fast(m int, a []int) (result int) {
	n := len(a)
	var l int
	for r := 1; r < len(a); r++ {
		for a[r]-a[l] > m {
			result += n - r
			l++
		}
	}

	return result
}

func slow(r int, a []int) (result int) {
	for i := range a {
		for j := i + 1; j < len(a); j++ {
			diff := int(math.Abs(float64(a[i] - a[j])))

			if diff > r {
				result += len(a) - j
				break
			}
		}
	}
	return result
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
