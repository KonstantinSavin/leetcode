package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{}
	res := dummy

	for head != nil {
		if head.Val != val {
			res.Next = head
			res = res.Next
		}
		head = head.Next
		res.Next = nil
	}
	return dummy.Next
}

func main() {
	head := ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 4}
	fmt.Println(removeElements(&head, 2))
}
