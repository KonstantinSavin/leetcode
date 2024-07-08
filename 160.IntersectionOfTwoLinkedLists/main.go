package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})

	for headA != nil {
		m[headA] = struct{}{}
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := m[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

func main() {
	headA := ListNode{Val: 1}
	headA.Next = &ListNode{Val: 2}
	headA.Next.Next = &ListNode{Val: 4}
	headA.Next.Next.Next = &ListNode{Val: 4}
	headB := ListNode{Val: 3}
	headB.Next = headA.Next.Next
	fmt.Println(getIntersectionNode(&headA, &headB))
}
