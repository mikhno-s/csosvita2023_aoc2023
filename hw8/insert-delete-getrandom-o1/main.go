package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	inserted map[int]int // [value]: index in allKeys
	removed  map[int]int
	allKeys  []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{make(map[int]int), make(map[int]int), make([]int, 0)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.inserted[val]; ok {
		return false
	}
	index := len(this.allKeys)
	if _, ok := this.removed[val]; ok {
		index = this.removed[val]
		delete(this.removed, val)
		this.allKeys[index] = val
	} else {
		this.allKeys = append(this.allKeys, val)
	}

	this.inserted[val] = index

	return true

}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.inserted[val]; !ok {
		return false
	}
	this.removed[val] = this.inserted[val]
	delete(this.inserted, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {

	i := rand.Intn(len(this.allKeys))
	for ; ; i = rand.Intn(len(this.allKeys)) {
		if _, ok := this.inserted[this.allKeys[i]]; ok {
			break
		}
	}

	return this.allKeys[i]
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Insert(1))
	fmt.Println(obj.Remove(2))
	fmt.Println(obj.Insert(2))
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.Remove(1))
	fmt.Println(obj.Insert(2))
	fmt.Println(obj.GetRandom())

}
