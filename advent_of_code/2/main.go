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
	fi, err := os.Open("instructions.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	reader := bufio.NewReader(fi)

	// Game 76: 4 blue, 3 green, 11 red; 4 blue, 1 green, 12 red; 11 blue, 2 green, 4 red; 2 blue, 2 green, 11 red; 12 red, 1 blue

	// only 12 red cubes, 13 green cubes, and 14 blue cubes
	// stones := map[string]int{"red": 12, "green": 13, "blue": 14}
	games := make([]string, 100)

	for i := 0; i < len(games); i++ {
		games[i] = readLine(reader)
	}

	var counter int

	for id := range games {

		// stonesMin := map[string]int{"red": 99999999, "green": 99999999, "blue": 99999999}
		stonesMax := map[string]int{"red": 0, "green": 0, "blue": 0}
		sets := strings.Split(games[id], ";")
		for _, set := range sets {

			set = strings.TrimSpace(set)

			pairs := strings.Split(set, ", ")
			stones := map[string]int{"red": 0, "green": 0, "blue": 0}
			for _, pair := range pairs {
				pair = strings.TrimSpace(pair)
				stone := strings.Split(pair, " ")[1]
				amount, _ := strconv.Atoi(strings.Split(pair, " ")[0])
				stones[stone] = amount
			}

			for k := range stonesMax {
				if stonesMax[k] < stones[k] {
					stonesMax[k] = stones[k]
				}
			}

			// for stone := range stonesMin {
			// 	if stonesMin[stone] > stonesMax[stone] {
			// 		stonesMin[stone] = stonesMax[stone]
			// 	}
			// }
			fmt.Println(stonesMax)
		}

		counter += stonesMax["red"] * stonesMax["green"] * stonesMax["blue"]
	}

	fmt.Println(counter)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
