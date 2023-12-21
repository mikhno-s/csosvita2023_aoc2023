package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// 56. Merge Intervals

// Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

// Example:

// Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].

// Constraints:

// 1 <= intervals.length <= 104
// intervals[i].length == 2
// 0 <= starti <= endi <= 104

type intervals [][]int

func (l intervals) Len() int { return len(l) }
func (l intervals) Less(i, j int) bool {
	for x := range l[i] {
		if l[i][x] == l[j][x] {
			continue
		}
		return l[i][x] < l[j][x]
	}
	return false
}

func (l *intervals) Swap(i, j int) { (*l)[i], (*l)[j] = (*l)[j], (*l)[i] }

func main() {

	fi, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	reader := bufio.NewReader(fi)

	intervals := make([][]int, 0)
	fmt.Println(reflect.TypeOf(intervals))
	// Reading intervals from an input file
	for s := readLine(reader); s != ""; s = readLine(reader) {
		strLine := strings.Split(s, " ")
		start, _ := strconv.Atoi(strLine[0])
		end, _ := strconv.Atoi(strLine[1])
		intervals = append(intervals, []int{start, end})
	}

	// Sorting intervals by start and finish
	sort.Slice(intervals, func(i, j int) bool {
		for x := range intervals[i] {
			if intervals[i][x] == intervals[j][x] {
				continue
			}
			return intervals[i][x] < intervals[j][x]
		}
		return false
	})

	result := make([][]int, 0)

	L, R := 0, 1

	// Adding the very first interval to result
	if len(intervals) > 0 {
		result = append(result, []int{intervals[L][0], intervals[L][1]})
	}

	// Start from the second interval
	for ; R < len(intervals); R += 1 {
		// Check if start of next interval is a part of previous interval by checking that start_i <= finish_i-1
		if intervals[R][0] <= result[L][1] {
			if intervals[R][1] > result[L][1] {
				result[L][1] = intervals[R][1]
			}
		} else {
			result = append(result, []int{intervals[R][0], intervals[R][1]})
			L += 1
		}
	}
	fmt.Println(result)

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
