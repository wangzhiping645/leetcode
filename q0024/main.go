package main

import (
	"fmt"

	"github.com/halfrost/LeetCode-Go/structures"
)

func main() {
	fmt.Println(structures.List2Ints(swapPairs(structures.Ints2List([]int{1, 2, 3}))))
	fmt.Println(structures.List2Ints(swapPairs(structures.Ints2List([]int{1, 2, 3, 4}))))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// ListNode define
type ListNode = structures.ListNode

// 多一步
func swapPairs1(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	x := dummy
	for dummy != nil {
		if dummy.Next == nil || dummy.Next.Next == nil {
			break
		}
		tmp := dummy.Next
		dummy.Next = dummy.Next.Next
		tmp.Next = tmp.Next.Next
		dummy.Next.Next = tmp
		dummy = tmp
	}
	return x.Next
}

// 递归 内存最佳
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}
