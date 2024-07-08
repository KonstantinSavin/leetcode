package main

import "fmt"

type MyQueue struct {
	queue []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{[]int{}}
}

/** Push element x onto queue. */
func (q *MyQueue) Push(x int) {
	q.queue = append(q.queue, x)
}

/** Removes the element on front of the queue and returns that element. */
func (q *MyQueue) Pop() int {
	frontEl := q.queue[0]
	q.queue = q.queue[1:]
	return frontEl
}

/** Get the front element. */
func (q *MyQueue) Peek() int {
	return q.queue[0]
}

/** Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
	return len(q.queue) == 0
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
	fmt.Println(s.Peek())

}
