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

// best
func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	if length == 2 {
		return mergeTwoLists(lists[0], lists[1])
	}
	num := length / 2
	left := mergeKLists(lists[:num])
	right := mergeKLists(lists[num:])
	return mergeTwoLists(left, right)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	l := dummy
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

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

// worst
func mergeKListsWorst(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var dummy *ListNode
	for _, l := range lists {
		dummy = mergeTwoLists(dummy, l)
	}
	return dummy
}
