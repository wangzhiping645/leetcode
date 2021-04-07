package main

import (
	"fmt"

	"github.com/halfrost/LeetCode-Go/structures"
)

func main() {
	p := []int{1, 2, 3, 4, 5}
	fmt.Printf("【input】:%v 【output】:%v", p, structures.List2Ints(removeNthFromEnd(structures.Ints2List(p), 2)))
}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
type ListNode = structures.ListNode

func getLength(head *ListNode) int {
	l := 0
	for ; head != nil; head = head.Next {
		l++
	}
	return l
}

// removeNthFromEnd 根据长度计算
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	l := 0
	l = getLength(head)

	dummy := &ListNode{0, head}
	t := dummy

	for i := 0; i < l-n; i++ {
		t = t.Next
	}
	t.Next = t.Next.Next
	return dummy.Next
}

// 栈
func removeNthFromEnd2(head *ListNode, n int) *ListNode {

	nodes := []*ListNode{}
	dummy := &ListNode{0, head}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes)-n-1]
	prev.Next = prev.Next.Next
	return dummy.Next
}

// 双指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	dummy := &ListNode{0, head}
	before, after := dummy, dummy
	i := 0
	for ; before != nil; before = before.Next {
		if i >= n+1 {
			after = after.Next
		}

		i++
	}
	//if n == 1 {
	//	after.Next = nil
	//} else {
	after.Next = after.Next.Next
	//}

	return dummy.Next

}
