package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return dummy.Next
}

func main() {
	list1 := &ListNode{Val: 1}
	list1.Next = &ListNode{Val: 2}
	list1.Next.Next = &ListNode{Val: 4}
	list2 := &ListNode{Val: 1}
	list2.Next = &ListNode{Val: 3}
	list2.Next.Next = &ListNode{Val: 4}
	// fmt.Println(list1, list2, list1.Next.Val)
	merged := mergeTwoLists(list1, list2)
	fmt.Println("The merged list l1 and l2 is:")
	for merged != nil {
		fmt.Printf("%d ", merged.Val)
		merged = merged.Next
	}
}
