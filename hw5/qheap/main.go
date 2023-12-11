package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Heap struct {
	Size  int
	array []int
}

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

	qN, _ := strconv.Atoi(readLine(reader))

	h := &Heap{}
	for i := 0; i < qN; i++ {
		c := strings.Split(strings.TrimSpace(readLine(reader)), " ")
		// fmt.Print("before ")
		// h.Print()
		if c[0] == "1" {
			v, _ := strconv.Atoi(c[1])
			h.Insert(v)
		}

		if c[0] == "2" {
			v, _ := strconv.Atoi(c[1])
			h.ExtractByValue(v)
		}

		if c[0] == "3" {
			fmt.Println(h.Min())
		}
		// fmt.Print("after ", h.Size)
		// h.Print()
	}
}

func (h *Heap) Print() {
	if h.Size < 1 || len(h.array) < 2 {
		fmt.Println(h.array[:h.Size])
	} else {
		fmt.Println(h.array[1 : h.Size+1])
	}

}

func (h *Heap) Min() int {
	if h.Size < 1 || len(h.array) < 2 {
		return 0
	}
	return h.array[1]
}

func (h *Heap) ExtractByValue(v int) {
	for i := 1; i <= h.Size && i < len(h.array); i++ {
		if h.array[i] == v {
			h.array[i], h.array[h.Size] = h.array[h.Size], h.array[i]
			h.Size -= 1
			h.siffDown(i)
			break
		}
	}
}

func (h *Heap) Insert(v int) {
	if h.Size == 0 {
		h.array = make([]int, 2)
	}

	h.Size += 1
	if h.Size >= len(h.array) {
		h.array = append(h.array, v)
	} else {
		h.array[h.Size] = v
	}
	h.siffUp(h.Size)
}

func (h *Heap) siffUp(i int) {
	for i > 1 && h.array[i/2] > h.array[i] {
		h.array[i/2], h.array[i] = h.array[i], h.array[i/2]
		i = i / 2
	}
}

func (h *Heap) siffDown(i int) {
	for 2*i <= h.Size {
		j := i

		if h.array[2*i] < h.array[j] {
			j = 2 * i
		}

		if 2*i+1 <= h.Size && h.array[2*i+1] < h.array[j] {
			j = 2*i + 1
		}

		if j == i {
			break
		}
		h.array[i], h.array[j] = h.array[j], h.array[i]
		i = j
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
