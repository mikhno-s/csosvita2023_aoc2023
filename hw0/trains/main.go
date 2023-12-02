package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Кожну добу на вокзал прибуває n електричок. По заданому розкладу їх прибуття визначте мінімальний час між прибуттям двох різних електричок.

// Input Format

// В першому рядку задано число n — кількість електричок У другому рядку задано n моментів часу в форматі HH:MM через пробіл.

// Constraints

// 1 ≤ n ≤ 2 × 10^4
// 0 ≤ HH ≤ 23
// 0 ≤ MM ≤ 59
// Output Format

// Виведіть одне число — мінімальний час в хвилинах між прибуттям двох електричок.

// Sample Input 0

// 2
// 23:59 00:00
// Sample Output 0

// 1
// Sample Input 1

// 3
// 00:00 23:59 00:00
// Sample Output 1

// 0

func main() {
	var numOfTrains int
	fmt.Scanln(&numOfTrains)

	in := bufio.NewReader(os.Stdin)
	arrivingTimesString, _ := in.ReadString('\n')

	arrivingTimesInMinutes := make([]int, 0, 0)
	for _, s := range strings.Split(arrivingTimesString, " ") {
		if len(strings.Split(s, ":")) == 2 {
			h, _ := strconv.Atoi(strings.Split(s, ":")[0])
			m, _ := strconv.Atoi(strings.Split(s, ":")[1])
			// Calculate total amount in minutes and append to the list
			arrivingTimesInMinutes = append(arrivingTimesInMinutes, h*60+m)
		}
	}

	// max diff
	diff := 1440
	if len(arrivingTimesInMinutes) >= 2 {
		for i, m1 := range arrivingTimesInMinutes {
			if i < len(arrivingTimesInMinutes) {
				for _, m2 := range arrivingTimesInMinutes[i+1:] {
					localDiff := 1440
					if m1 == m2 {
						fmt.Println(0)
						os.Exit(0)
					}
					// calculate diff
					if m1 > m2 {
						localDiff = m1 - m2
					} else {
						localDiff = m2 - m1
					}
					// if difference more then 12 hours then count bound it
					if localDiff > 720 {
						localDiff = 1440 - localDiff
					}

					if localDiff < diff {
						diff = localDiff
					}
				}
			}
		}
	}
	fmt.Println(diff)
}

//23:59 00:00
