package main

import "fmt"

type Node struct {
	t    int
	next *Node
}

type RecentCounter struct {
	size int
	head *Node
	tail *Node
}

func Constructor() RecentCounter {
	rc := &RecentCounter{}
	node := &Node{}
	rc.head = node
	rc.tail = rc.head
	return *rc

}

func (this *RecentCounter) Ping(t int) int {
	// enqueue
	next := &Node{t, nil}
	this.tail.next = next
	this.tail = next
	if this.size == 0 {
		this.head = this.head.next
	}
	this.size++
	// dequeue
	for this.head.t < t-3000 {
		this.head = this.head.next
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
