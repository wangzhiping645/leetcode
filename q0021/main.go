package main

import "github.com/halfrost/LeetCode-Go/structures"

func main() {

}

// ListNode define
type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	} else {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	l := dummy

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			l.Next = l2
			l2 = l2.Next
		} else {
			l.Next = l1
			l1 = l1.Next
		}
		l = l.Next
	}
	if l1 != nil {
		l.Next = l1
	}
	if l2 != nil {
		l.Next = l2
	}
	return dummy.Next
}
