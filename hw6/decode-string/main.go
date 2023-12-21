package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type StringStack struct {
	top int
	a   []string

	Empty bool
}

type Encoded struct {
	t       int
	encoded []Encoded
	decoded string
}

func main() {

	// stack := NewStringStack()
	// fmt.Println(stack.Pop())
	// fmt.Println(stack.Pop())
	s := "3[q]2[a]"
	// s := strings.Split("2[abc]3[cd]ef", "")
	// s := strings.Split("3[a2[c]]", "")

	// 3 [ a 2 [ c ]  ]as
	// e := &Encoded{}
	// L, R := 0, len(s)
	fmt.Println(decodeString(s))
	// stack.Print()
}

func decodeString(s string) string {
	result := ""
	last := len(s)
	for L := 0; L < last; L++ {
		if strings.Contains("1234567890", s[L:L+1]) {
			n, _ := strconv.Atoi(s[L : L+1])
			R := strings.Index(s[L:last], "]")
			for i := 0; i < n; i++ {
				result += decodeString(s[L+2 : R])
			}
			L = R
		} else {
			result += s[L : L+1]
		}
	}

	// for !stack.Empty {
	// 	result += stack.Pop()
	// }
	return result
}

func (s *StringStack) Push(o string) {
	if s.top >= len(s.a) {
		s.a = append(s.a, o)
	} else {
		s.a[s.top] = o
	}
	s.Empty = false
	s.top += 1

}

func (s *StringStack) Pop() string {
	if s.Empty {
		return ""
	}
	s.top -= 1
	if s.top == 0 {
		s.Empty = true
	}
	return s.a[s.top]
}

func (s *StringStack) Print() {
	for i := s.top - 1; i >= 0; i-- {
		fmt.Printf("%s", s.a[i])
	}
	fmt.Println()
}

func NewStringStack() *StringStack {
	return &StringStack{0, make([]string, 0), true}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
