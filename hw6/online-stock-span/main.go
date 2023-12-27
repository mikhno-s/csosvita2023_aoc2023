package main

import "fmt"

type StockSpan struct {
	Price    int
	SpanSize int
}

type StockSpanner struct {
	StackSize      int
	StockMonoStack []*StockSpan
}

func Constructor() StockSpanner {
	return StockSpanner{0, make([]*StockSpan, 0)}
}

func (this *StockSpanner) Next(price int) int {
	stockSpan := &StockSpan{Price: price, SpanSize: 1}
	for this.StackSize > 0 && price >= this.StockMonoStack[this.StackSize-1].Price {
		stockSpan.SpanSize += this.StockMonoStack[this.StackSize-1].SpanSize
		this.Pop()
	}
	this.Push(stockSpan)

	return stockSpan.SpanSize
}

func (this *StockSpanner) Push(v *StockSpan) {
	if this.StackSize >= len(this.StockMonoStack) {
		this.StockMonoStack = append(this.StockMonoStack, v)
	} else {
		this.StockMonoStack[this.StackSize] = v
	}
	this.StackSize += 1
}

func (this *StockSpanner) Pop() {
	this.StackSize -= 1
}

func main() {
	stock := Constructor()

	fmt.Println(stock.Next(28))
	fmt.Println(stock.Next(14))
	fmt.Println(stock.Next(28))
	fmt.Println(stock.Next(35))
	fmt.Println(stock.Next(46))
	fmt.Println(stock.Next(53))

}
