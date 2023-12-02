// Month | + Interest | - Payment | = Balance
// ------------------------------------------
//       |            |           |  2000.00
//    1  |     15.94  |   510.00  |  1505.94
//    2  |     12.00  |   510.00  |  1007.94
//    3  |      8.03  |   510.00  |   505.97
//    4  |      4.03  |   510.00  |     0.00

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
	price, _ := strconv.ParseFloat(s[0], 64)
	monthlyPayment, _ := strconv.ParseFloat(s[1], 64)
	loanTerm, _ := strconv.Atoi(s[2])

	// totalInterest := calculateLastBalanceByInterest(price, monthlyPayment, 9.562054624583681/100/12, loanTerm)
	// totalInterest := calculateInterstsBody(price, monthlyPayment, 9.562054624583681/100/12, loanTerm)
	// fmt.Println(totalInterest)

	interestBody := monthlyPayment*float64(loanTerm) - price

	bad, good := 0.0, 100.0/100/12
	var cand float64
	cand = (bad + good) / 2
	for good != cand && bad != cand {
		m := calculateInterstsBody(price, monthlyPayment, cand, loanTerm)
		// fmt.Println(good, bad, m)
		if m >= interestBody {
			good = cand
		} else {
			bad = cand
		}
		cand = (bad + good) / 2
	}

	fmt.Println(good * 100 * 12)
}

func calculateInterstsBody(balance float64, monthlyPayment float64, interest float64, monthLeft int) float64 {
	interestBodySum := 0.0

	for ; monthLeft > 0; monthLeft-- {
		interestBody := balance * interest
		interestBodySum += interestBody
		balance = balance - monthlyPayment + interestBody
	}
	return interestBodySum
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
