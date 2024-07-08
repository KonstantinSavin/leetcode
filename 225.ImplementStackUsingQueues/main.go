package main

import "fmt"

type MyStack struct {
	queue []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{[]int{}}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	topEl := this.queue[len(this.queue)-1]
	this.queue = this.queue[:len(this.queue)-1]
	return topEl
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.queue[len(this.queue)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

func main() {
	s := Constructor()
	fmt.Println(s.Empty())
	s.Push(1)
	s.Push(4)
	s.Push(2)
	s.Push(3)
	fmt.Println(s)
	fmt.Println(s.Empty())
	fmt.Println(s.Pop())
	fmt.Println(s.Top())

}
