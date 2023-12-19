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
	// add map with the [element]index of element
	m map[int]int
}

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

	// reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

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
	fmt.Println(h.array[:h.Size])
}

func (h *Heap) Min() int {
	if h.Size < 1 {
		return 0
	}
	return h.array[0]
}

func (h *Heap) ExtractByValue(v int) {
	for i := 0; i < h.Size; i++ {
		// add map to quickly found the index of element
		if h.array[i] == v {
			h.Size -= 1
			h.array[i], h.array[h.Size] = h.array[h.Size], h.array[i]
			h.siffDown(i)
			break
		}
	}
}

func (h *Heap) Insert(v int) {
	if h.Size == 0 {
		h.array = make([]int, 0)
		h.m = make(map[int]int)
	}

	if h.Size >= len(h.array) {
		h.array = append(h.array, v)
		h.Size += 1
	} else {
		h.array[h.Size-1] = v
	}

	h.siffUp(h.Size - 1)

}

func (h *Heap) siffUp(i int) {
	for i > 0 && h.array[i/2] > h.array[i] {
		h.array[i/2], h.array[i] = h.array[i], h.array[i/2]
		h.m[i/2] = h.array[i/2]
		h.m[i] = h.array[i]
		i = i / 2
	}
}

func (h *Heap) siffDown(i int) {
	for h.Size > 0 && 2*i+1 < h.Size {
		j := i

		if h.array[2*i+1] < h.array[j] {
			j = 2*i + 1
		}

		if 2*i+2 < h.Size && h.array[2*i+2] < h.array[j] {
			j = 2*i + 2
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
