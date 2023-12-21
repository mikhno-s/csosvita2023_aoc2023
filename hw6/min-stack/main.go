package main

import "fmt"

type MinStack struct {
	size int
	a    []int
	min  *MinStack
}

func Constructor() MinStack {
	s := new(MinStack)
	s.min = &MinStack{0, make([]int, 0), nil}
	s.a = make([]int, 0)
	return *s
}

func (this *MinStack) Push(val int) {
	if this.min != nil && (this.size == 0 || this.min.Top() >= val) {
		this.min.Push(val)
	}
	if this.size >= len(this.a) {
		this.a = append(this.a, val)
	} else {
		this.a[this.size] = val
	}

	this.size += 1

}

func (this *MinStack) Pop() {
	if this.min != nil && this.a[this.size-1] == this.min.Top() {
		this.min.Pop()
	}
	this.size -= 1
}

func (this *MinStack) Top() int {
	return this.a[this.size-1]
}

func (this *MinStack) GetMin() int {
	return this.min.Top()
}

func main() {

	minStack := Constructor()
	minStack.Push(0)
	minStack.Push(1)
	minStack.Push(0)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.GetMin())

}
