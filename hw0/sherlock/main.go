package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
Watson gives Sherlock an array of integers. His challenge is to find an element of the array such that the sum of all elements to the left is equal to the sum of all elements to the right.

Example

arr = [5,6,8,11]

 8 is between two subarrays that sum to 11.

 arr = 1

The answer is [1] since left and right sum to 0.

You will be given arrays of integers and must determine whether there is an element that meets the criterion. If there is, return YES. Otherwise, return NO.

Function Description

Complete the balancedSums function in the editor below.

balancedSums has the following parameter(s):

int arr[n]: an array of integers
Returns

string: either YES or NO
Input Format

The first line contains , the number of test cases.

The next  pairs of lines each represent a test case.
- The first line contains , the number of elements in the array .
- The second line contains  space-separated integers  where .

 * Complete the 'balancedSums' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER_ARRAY arr as parameter.
*/

func balancedSums(arr []int) string {

	var sumL int
	var sumR int
	for _, n := range arr {
		sumR += n
	}

	for _, n := range arr {
		if sumR-n == sumL {
			return "YES"
		}
		sumL += n
		sumR -= n
	}

	return "NO"
}

func main() {

	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	T, err := strconv.Atoi(strings.TrimSpace(readLine(reader)))
	checkError(err)

	for TItr := 0; TItr < int(T); TItr++ {
		n, err := strconv.Atoi(strings.TrimSpace(readLine(reader)))
		checkError(err)

		arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var arr []int

		for i := 0; i < int(n); i++ {
			arrItem, err := strconv.Atoi(arrTemp[i])
			checkError(err)
			arr = append(arr, arrItem)
		}

		result := balancedSums(arr)

		fmt.Printf("%s\n", result)
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
