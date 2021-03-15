package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 链表翻转
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	var next *ListNode = nil

	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}

	return pre
}
