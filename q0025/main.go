package main

import (
	"fmt"

	"github.com/halfrost/LeetCode-Go/structures"
)

type A struct {
	a int
	b int
}

func a() (A, int) {
	return A{1, 2}, 3
}

func a2() (int, A) {
	return 31, A{11, 21}
}

//func main() {
//	t := A{}
//	t, t.a = a()
//	fmt.Println(t)
//	t.a, t = a2()
//	fmt.Println(t)
//}

func main() {
	//fmt.Println(structures.List2Ints(reverseKGroup(structures.Ints2List([]int{1, 2, 3, 4}), 1)))
	//fmt.Println(structures.List2Ints(reverseKGroup(structures.Ints2List([]int{1, 2, 3, 4}), 2)))
	//fmt.Println(structures.List2Ints(reverseKGroup(structures.Ints2List([]int{1, 2, 3, 4, 5}), 2)))
	//fmt.Println(structures.List2Ints(reverseKGroup(structures.Ints2List([]int{1, 2, 3, 4, 5, 6}), 2)))
	fmt.Println(structures.List2Ints(reverseKGroup(structures.Ints2List([]int{1, 2, 3, 4, 5, 6, 7}), 3)))
	//fmt.Println(structures.List2Ints(reverseKGroup(structures.Ints2List([]int{1, 2, 3, 4}), 4)))
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

// 迭代，内存更优
func reverseKGroup(head *ListNode, k int) *ListNode {

	if k == 0 || k == 1 {
		return head
	}
	dummy := &ListNode{0, head}
	prev := dummy

	for head != nil {
		tmp := head
		for i := 0; i < k; i++ {
			if tmp == nil {
				return dummy.Next
			}
			tmp = tmp.Next
		}

		prev.Next, prev = reverseFirstK(head, k)
		head = prev.Next
		//h, t := reverseFirstK(head, k)
		//prev.Next = h
		//prev = t
		//head = t.Next
		fmt.Println(prev.Val)
	}
	return dummy.Next
}

// 递归
func reverseKGroup1(head *ListNode, k int) *ListNode {

	if k == 0 || k == 1 {
		return head
	}

	tmp := head
	for i := 0; i < k; i++ {
		if tmp == nil {
			return head
		}
		tmp = tmp.Next
	}

	var tail *ListNode
	head, tail = reverseFirstK(head, k)
	tail.Next = reverseKGroup1(tail.Next, k)
	return head
}

// reverseFirstK 翻转前 k 个节点
func reverseFirstK(p *ListNode, k int) (*ListNode, *ListNode) {

	head := p
	for i := 1; i < k && p != nil && p.Next != nil; {
		next := p.Next
		p.Next = next.Next
		next.Next = head
		head = next
		i++
	}
	return head, p
}
