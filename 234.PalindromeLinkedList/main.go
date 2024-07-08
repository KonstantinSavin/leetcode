package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	prev := &ListNode{}
	prev = nil
	dir := []int{}

	for head != nil {
		dir = append(dir, head.Val)
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	rev := []int{}
	for prev != nil {
		rev = append(rev, prev.Val)
		prev = prev.Next
	}
	// fmt.Println(dir, rev)

	for i := range dir {
		if dir[i] != rev[i] {
			return false
		}
	}

	return true
}

func main() {
	head := ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	fmt.Println(isPalindrome(&head))
	head = ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 1}
	fmt.Println(isPalindrome(&head))
}
