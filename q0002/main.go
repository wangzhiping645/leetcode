package main

import (
	"fmt"
	"github.com/halfrost/LeetCode-Go/structures"
)

// ListNode define
type ListNode = structures.ListNode

func main() {

	i := 10000
	j := 80009
	for {
		fmt.Println(j / i)
		if j/i != 0 {
			j = j % (j / i * i)
		}

		i /= 10
		if i == 0 {
			break
		}
	}

	i = 10
	j = 80009
	for {
		fmt.Println(j % i)
		//if j/i != 0 {
		j = j / i
		//}

		//i *= 10
		if j == 0 {
			break
		}
	}

	//a1 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	//a2 := []int{5, 6, 4}
	a1 := []int{2, 4, 3}
	a2 := []int{5, 6, 4}
	fmt.Println(structures.List2Ints(addTwoNumbers(structures.Ints2List(a1), structures.Ints2List(a2))))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ll := &ListNode{}
	t := ll
	before := 0
	for l1 != nil || l2 != nil || before != 0 {
		tmp := 0
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		tmp += before

		t.Next = &ListNode{
			Val: tmp % 10,
		}
		if tmp >= 10 {
			before = 1
		} else {
			before = 0
		}
		t = t.Next
	}
	return ll.Next
}
