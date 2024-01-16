package main

import (
	"fmt"
	"strconv"
)

type TimeMap struct {
	// key -> timestamp (as string, because I don't want to use interfaces in slice) -> value
	// also [][]string can be changed by []Val{timestamp, value}
	data map[string][][]string // key: [ [timestamp, value], ... ]
}

func Constructor() TimeMap {
	return TimeMap{
		make(map[string][][]string),
	}

}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.data[key] = append(this.data[key], []string{fmt.Sprint(timestamp), value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	data := this.data[key]

	// if key is empty
	if len(data) == 0 {
		return ""
	}

	// if min timestamp is > given timestamp
	minTs, _ := strconv.Atoi(data[0][0])
	if timestamp < minTs {
		return ""
	}

	min, max := 0, len(data)
	val := data[0][1]
	for max-min > 1 {
		m := min + (max-min)/2
		current, _ := strconv.Atoi(data[m][0])
		if current > timestamp {
			max = m
		} else {
			val = data[m][1]
			min = m
		}
	}
	return val
}

func main() {
	obj := Constructor()
	obj.Set("foo", "bar", 1)
	fmt.Println(obj.Get("foo", 1)) //bar
	fmt.Println(obj.Get("foo", 3)) //bar
	obj.Set("foo", "bar2", 4)
	fmt.Println(obj.Get("foo", 4)) //bar2
	fmt.Println(obj.Get("foo", 5)) //bar2
	obj.Set("foo2", "buzz", 100)
	obj.Set("foo2", "buzz2", 101)
	fmt.Println(obj.Get("foo2", 199)) //buzz2

}
