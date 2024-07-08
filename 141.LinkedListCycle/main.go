package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head.Next
	if slow == nil {
		return false
	}

	fast := slow.Next
	for fast != nil && slow != nil {
		if fast == slow {
			return true
		}
		slow = slow.Next

		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	return false
}

func main() {
	head := ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 4}
	fmt.Println(hasCycle(&head))
}
