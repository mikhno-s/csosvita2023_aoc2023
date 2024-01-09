package main

import "fmt"

type RecentCounter struct {
	a    []int
	size int
	head int
	tail int
}

func Constructor() RecentCounter {
	rc := &RecentCounter{}
	rc.a = make([]int, 0)
	return *rc

}

func (this *RecentCounter) Ping(t int) int {
	// enqueue
	this.a = append(this.a, t)
	this.tail++
	this.size++
	// dequeue
	for this.a[this.head] < t-3000 {
		this.head++
		this.size--
	}

	return this.size
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Ping(642))
	fmt.Println(obj.Ping(1849))
	fmt.Println(obj.Ping(4921))
	fmt.Println(obj.Ping(5936))
	fmt.Println(obj.Ping(5957))
}
