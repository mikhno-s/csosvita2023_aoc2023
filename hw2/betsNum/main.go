package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	var arr []int
	for i := 0; i < len(arrTemp); i++ {
		arrItem, err := strconv.Atoi(arrTemp[i])
		checkError(err)
		arr = append(arr, arrItem)
	}
	var bestPrime, bestPrimeSum int
	for L := arr[0]; L <= arr[1]; L++ {

		if big.NewInt(int64(L)).ProbablyPrime(20) {
			var sum int

			for i := L; i > 0; i /= 10 {
				sum += int(math.Pow(float64(i%10), 2))
			}
			if bestPrimeSum < sum {
				bestPrime = L
				bestPrimeSum = sum
			}
		}
	}

	if bestPrime == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(bestPrime)
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

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"math"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
// 	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
// 	var arr []int
// 	for i := 0; i < len(arrTemp); i++ {
// 		arrItem, err := strconv.Atoi(arrTemp[i])
// 		checkError(err)
// 		arr = append(arr, arrItem)
// 	}

// 	var primeArray []int
// 	var bestPrime, bestPrimeSum int
// 	for L := arr[0]; L <= arr[1]; L++ {

// 		// min prime
// 		if L == 2 {
// 			primeArray = append(primeArray, L)
// 			continue
// 		}

// 		// check if last num is 0 or 2/4/6/8 or last num is 5
// 		if L%10 == 0 || L%10%2 == 0 || L != 5 && L%10 == 5 {
// 			continue
// 		}

// 		// check if the sum of numbers is divisible by 3
// 		var sum int
// 		for i := L; i > 0; i /= 10 {
// 			sum += i % 10
// 			if L != 3 && sum%3 == 0 {
// 				continue
// 			}

// 		}

// 		root := int(math.Sqrt(float64(L)))
// 		var divisible bool
// 		// Dividing the given number by all the prime numbers below its square root value.
// 		for i := 0; i < len(primeArray) && primeArray[i] <= root; i++ {
// 			if L%root == 0 {
// 				divisible = true
// 				break
// 			}
// 		}
// 		if !divisible {
// 			primeArray = append(primeArray, L)
// 			var currentSum int
// 			for i := L; i > 0; i /= 10 {
// 				currentSum += int(math.Pow(float64(i%10), 2))
// 			}
// 			if bestPrimeSum < currentSum {
// 				bestPrime = L
// 				bestPrimeSum = currentSum
// 			}
// 			if bestPrimeSum == currentSum {
// 				if bestPrime > L {
// 					bestPrime = L
// 				}
// 			}
// 		}
// 	}

// 	if bestPrime == 0 {
// 		fmt.Println(-1)
// 	} else {
// 		fmt.Println(bestPrime)
// 	}
// }

// func readLine(reader *bufio.Reader) string {
// 	str, _, err := reader.ReadLine()
// 	if err == io.EOF {
// 		return ""
// 	}

// 	return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
