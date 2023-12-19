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

type CellsGroup struct {
	Sum   int
	Cells []int
}

type Cells []*CellsGroup

func main() {

	// fi, err := os.Open("input")
	// if err != nil {
	// 	panic(err)
	// }

	// defer func() {
	// 	if err := fi.Close(); err != nil {
	// 		panic(err)
	// 	}
	// }()
	// reader := bufio.NewReader(fi)

	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	N, _ := strconv.Atoi(readLine(reader))
	groupSize := int(math.Sqrt(float64(N)))
	cells := make(Cells, 0)
	groupID, i := -1, -1
	for _, cellClients := range strings.Split(strings.TrimSpace(readLine(reader)), " ") {
		// init new kuchka
		if i == groupSize || i == -1 {
			i = 0
			groupID++
			cells = append(cells, &CellsGroup{})
			cells[groupID].Cells = make([]int, groupSize)
		}
		cells[groupID].Cells[i], _ = strconv.Atoi(cellClients)
		cells[groupID].Sum += cells[groupID].Cells[i]
		i++
	}

	Q, _ := strconv.Atoi(readLine(reader))
	for i := 0; i < Q; i++ {
		c := strings.Split(strings.TrimSpace(readLine(reader)), " ")
		if c[0] == "COUNT" {
			L, _ := strconv.Atoi(c[1])
			R, _ := strconv.Atoi(c[2])
			fmt.Println(cells.sum(L-1, R-1))
		}

		if c[0] == "ENTER" {
			groupID, _ := strconv.Atoi(c[1])
			cells.set(groupID-1, 1)
		}

		if c[0] == "LEAVE" {
			groupID, _ := strconv.Atoi(c[1])
			cells.set(groupID-1, -1)
		}
	}
}

func (c *Cells) sum(L int, R int) (sum int) {
	if len(*c) == 0 {
		return sum
	}

	groupSize := len((*c)[0].Cells)
	LGroupID := L / groupSize
	RGroupID := R / groupSize

	for i := L - LGroupID*groupSize; i < groupSize; i++ {
		sum += (*c)[LGroupID].Cells[i]
	}

	if LGroupID != RGroupID {
		for i := LGroupID + 1; i < RGroupID; i++ {
			sum += (*c)[i].Sum
		}
		for i := 0; i <= R-RGroupID*groupSize; i++ {
			sum += (*c)[RGroupID].Cells[i]
		}
	}

	return sum
}

func (c *Cells) set(i int, v int) {
	if len(*c) == 0 {
		return
	}
	groupSize := len((*c)[0].Cells)
	groupID := i / groupSize

	(*c)[groupID].Cells[i-groupID*groupSize] += v
	(*c)[groupID].Sum += v
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
