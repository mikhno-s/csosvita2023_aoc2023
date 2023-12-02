package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b, c, d int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	fmt.Scanln(&d)
	result := ""
	for x := 0; x <= 1000; x++ {
		if a*x*x*x+b*x*x+c*x+d == 0 {
			result = result + strconv.Itoa(x) + " "
		}
	}
	fmt.Println(result)
}
