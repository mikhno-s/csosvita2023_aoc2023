package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type MinHeap []int
type MaxHeap []int

func add(max *MaxHeap, min *MinHeap, n int) {
	heap.Push(max, n)
	if len(*max)-len(*min) > 1 {
		heap.Push(min, heap.Pop(max))
	}

	if len(*min) > 0 && (*max)[0] > (*min)[0] {
		_min := heap.Pop(min)
		_max := heap.Pop(max)
		heap.Push(max, _min)
		heap.Push(min, _max)
	}
	return
}

func findMedian(max *MaxHeap, min *MinHeap) float64 {
	if len(*min) == 0 {
		return float64((*max)[0])
	}
	if (len(*max)+len(*min))%2 == 0 {
		return float64(((*max)[0] + (*min)[0])) / 2
	}
	return float64((*max)[0])
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

	aCount, _ := strconv.Atoi(strings.TrimSpace(readLine(reader)))

	minheap := &MinHeap{}
	maxheap := &MaxHeap{}
	for i := 0; i < aCount; i++ {
		n, _ := strconv.Atoi(strings.TrimSpace(readLine(reader)))
		add(maxheap, minheap, n)
		// fmt.Println(*maxheap, *minheap)
		fmt.Printf("%.1f\n", findMedian(maxheap, minheap))
	}
}

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

// Push and Pop use pointer receivers because they modify the slice's length,
// not just its contents.
func (h *MinHeap) Push(x any) { *h = append(*h, x.(int)) }

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }

// This version from examples is very unoptimized. a lot of memmove
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
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
