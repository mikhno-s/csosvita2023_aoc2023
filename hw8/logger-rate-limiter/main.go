package main

import "fmt"

type Logger struct {
	cache map[string]int
}

func Constructor() Logger {
	return Logger{map[string]int{}}
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if this.cache[message] > timestamp {
		return false
	}

	this.cache[message] = timestamp + 10

	return true
}

func main() {
	obj := Constructor()
	fmt.Println(obj.ShouldPrintMessage(1, "foo"))  // return true, next allowed timestamp for "foo" is 1 + 10 = 11
	fmt.Println(obj.ShouldPrintMessage(2, "bar"))  // return true, next allowed timestamp for "bar" is 2 + 10 = 12
	fmt.Println(obj.ShouldPrintMessage(3, "foo"))  // 3 < 11, return false
	fmt.Println(obj.ShouldPrintMessage(8, "bar"))  // 8 < 12, return false
	fmt.Println(obj.ShouldPrintMessage(10, "foo")) // 10 < 11, return false
	fmt.Println(obj.ShouldPrintMessage(11, "foo")) // 11 >= 11, return true, next allowed timestamp for "foo" is 11 + 10 = 21

}
