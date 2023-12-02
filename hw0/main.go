package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	N := 6
	precomputed := make([]int, 0, N)
	denom := 6 // 3!
	for i := 1; i <= N; i++ {
		precomputed = append(precomputed, denom)
		denom *= (2*i + 2) * (2*i + 3)

	}

	testXNum := 100000000 // 1 kk
	floats := make([]float32, testXNum)
	for i := 0; i < testXNum; i++ {
		// Generate a random float between 0.0 and 1.0 and store it in the slice
		floats[i] = rand.Float32()
	}

	start := time.Now()
	_ = sinx(testXNum, 6, floats)
	elapsed := time.Since(start)

	start = time.Now()
	_ = sinxOptimized(testXNum, 6, precomputed, floats)
	elapsedOptimised := time.Since(start)

	fmt.Println(elapsed.Milliseconds() - elapsedOptimised.Milliseconds())
}

func sinx(N int, terms int, x []float32) (result []float32) {
	for i := 0; i < N; i++ {
		value := x[i]
		numer := x[i] * x[i] * x[i]
		denom := 6 // 3!
		sign := -1
		for j := 1; j <= terms; j++ {
			value += float32(sign) * numer / float32(denom)
			numer *= x[i] * x[i]
			denom *= (2*j + 2) * (2*j + 3)
			sign *= -1
		}
		result = append(result, value)
	}
	return result
}

func sinxOptimized(N int, terms int, denom []int, x []float32) (result []float32) {
	for i := 0; i < N; i++ {
		value := x[i]
		numer := x[i] * x[i] * x[i]
		sign := -1
		for j := 1; j <= terms; j++ {
			value += float32(sign) * numer / float32(denom[j-1])
			numer *= x[i] * x[i]
			sign *= -1
		}
		result = append(result, value)
	}
	return result
}
