package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var n, m int

	fmt.Scanln(&n)
	tshirts := make([]float64, n)
	fmt.Fscan(os.Stdin, ReadTo(tshirts)...)

	fmt.Scanln(&m)
	pants := make([]float64, m)
	fmt.Fscan(os.Stdin, ReadTo(pants)...)

	// fmt.Println(pants)
	fmt.Println(fast(tshirts, pants))

}

func slow(tshirts []float64, pants []float64) (int, int) {
	minDiff := math.Abs(tshirts[0] - pants[0])
	minP := pants[0]
	minT := tshirts[0]
	for _, T := range tshirts {
		for _, P := range pants {
			currentDiff := math.Abs(T - P)
			if currentDiff < minDiff {
				minT, minP = T, P
				minDiff = currentDiff
			}
		}
	}
	return int(minT), int(minP)
}

func fast(tshirts []float64, pants []float64) (int, int) {
	var minT, minP, T, P int
	n := len(tshirts)
	m := len(pants)

	for T < n && P < m {
		if math.Abs(tshirts[minT]-pants[minP]) > math.Abs(tshirts[T]-pants[P]) {
			minT, minP = T, P
		}
		if tshirts[T] < pants[P] {
			T++
		} else {
			P++
		}
	}

	return int(tshirts[minT]), int(pants[minP])
}

func ReadTo(n []float64) []interface{} {
	p := make([]interface{}, len(n))
	for i := range n {
		p[i] = &n[i]
	}
	return p
}
