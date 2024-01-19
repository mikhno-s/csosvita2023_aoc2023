package main

import "fmt"

const (
	maxloadfactor        = 2
	deafultHashTableSize = 16
)

type Node struct {
	value int
	next  *Node
}
type HashMap struct {
	size       int
	loadFactor float32
	data       []*Node
}

func (this *HashMap) hasChode(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func (this *HashMap) resize() {
	rehashedData := make([]*Node, len(this.data)*2)
	this.size = 0
	for _, el := range this.data {
		for ; el != nil; el = el.next {
			index := this.hasChode(el.value) % len(rehashedData)
			this.setByIndex(index, el.value, rehashedData)
			this.size++
		}
	}
	this.data = rehashedData

}

func (this *HashMap) setByIndex(index int, v int, data []*Node) bool {
	if data[index] == nil {
		data[index] = &Node{v, nil}
	} else {
		current := data[index]
		for ; current.value != v && current.next != nil; current = current.next {
		}
		if current.value == v {
			return false
		}
		current.next = &Node{v, nil}
	}

	this.loadFactor = float32(this.size) / float32(len(data))

	return true
}

func (this *HashMap) Set(i int) bool {
	if this.loadFactor >= maxloadfactor {
		this.resize()
	}

	index := this.hasChode(i) % len(this.data)
	res := this.setByIndex(index, i, this.data)
	this.size++
	return res
}

func containsDuplicate(nums []int) bool {
	hashMap := &HashMap{0, 0, make([]*Node, deafultHashTableSize)}
	for _, num := range nums {
		if !hashMap.Set(num) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{274, -735, -911, 80, 454, -511, 922, -775, 985, -669, -463, -896, -629, -586, 910, -361, 288, -375, 88, 556, -578, -406, -87, 25, 377, -635, -445, -289, 646, -487, -924, -968, -962, 502, 36, 129, -611, 54, -27, -496, 915, -84, -782, 349, 678, 332, -114, 345, 14, 119, 710, 821, -194, 988, 38, -369, 409, -559, -529, -298, -593, 705}))
}
