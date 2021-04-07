package main

import (
	"fmt"

	"github.com/halfrost/LeetCode-Go/structures"
)

// TreeNode define
type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
        3
    5        1
  6   2    0   8
     7 4
*/
func main() {
	t := structures.Ints2TreeNode([]int{3, 5, 1, 6, 2, 0, 8, structures.NULL, structures.NULL, 7, 4})
	fmt.Println("前序")
	preOrderTraverse(t)
	fmt.Println()
	preOrderTraverse2(t)
	fmt.Println()
	fmt.Println("中序")
	inOrderTraverse(t)
	fmt.Println()
	inOrderTraverse2(t)
	fmt.Println()
	fmt.Println("后序")
	postOrderTraverse(t)
	fmt.Println()
	postOrderTraverse2(t)
	fmt.Println()
	fmt.Println("bfs")
	bfs(t)
	fmt.Println()
	bfsKLevel(t, 1)
	fmt.Println()
	bfsKLevel(t, 2)
	fmt.Println()
	bfsKLevel(t, 4)

}
func preOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val)
	preOrderTraverse(root.Left)
	preOrderTraverse(root.Right)

}

func preOrderTraverse2(root *TreeNode) {
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			fmt.Print(root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) > 0 {
			root = stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]

		}
	}
}
func inOrderTraverse2(root *TreeNode) {
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) > 0 {
			root = stack[len(stack)-1]
			fmt.Print(root.Val)
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
}

func postOrderTraverse2(root *TreeNode) {
	stack := []*TreeNode{}
	var pre *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		if (top.Right == nil && top.Left == nil) || (top.Right == nil && top.Left == pre) || pre == top.Right {
			fmt.Print(top.Val)
			pre = top
			stack = stack[:len(stack)-1]

		} else {
			root = top.Right
		}

	}
}

// 中左右
func preOrderTraverse1(root *TreeNode) {
	//stack := []*TreeNode{root}
	//for len(stack) > 0 {
	//	node := stack[len(stack)-1]
	//	stack = stack[:len(stack)-1]
	//	if node == nil {
	//		continue
	//	}
	//	stack = append(stack, node.Right)
	//	stack = append(stack, node.Left)
	//	fmt.Print(node.Val)
	//}

	//stack := []*TreeNode{}
	//for len(stack) > 0 || root != nil {
	//	if root == nil {
	//		root = stack[len(stack)-1]
	//		stack = stack[:len(stack)-1]
	//		continue
	//	} else {
	//		fmt.Print(root.Val)
	//		stack = append(stack, root.Right)
	//		stack = append(stack, root.Left)
	//		root = stack[len(stack)-1]
	//		stack = stack[:len(stack)-1]
	//	}
	//}

	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			fmt.Print(root.Val)
			stack = append(stack, root)
			root = root.Left
		}

		if len(stack) > 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
}

func inOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	inOrderTraverse(root.Left)
	fmt.Print(root.Val)
	inOrderTraverse(root.Right)

}

// 左中右
func inOrderTraverse1(root *TreeNode) {
	//stack := []*TreeNode{}
	//for len(stack) > 0 || root != nil {
	//	if root == nil {
	//		root = stack[len(stack)-1]
	//		stack = stack[:len(stack)-1]
	//		fmt.Print(root.Val)
	//		root = root.Right
	//	} else {
	//		stack = append(stack, root)
	//		root = root.Left
	//	}
	//}

	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		if len(stack) > 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Print(root.Val)
			root = root.Right
		}
	}
}

func postOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	postOrderTraverse(root.Left)
	postOrderTraverse(root.Right)
	fmt.Print(root.Val)

}

// 左右中
func postOrderTraverse1(root *TreeNode) {
	//stack := []*TreeNode{}
	//var pre *TreeNode
	//for len(stack) > 0 || root != nil {
	//	for root != nil {
	//		stack = append(stack, root)
	//		root = root.Left
	//	}
	//	root = stack[len(stack)-1]
	//	stack = stack[:len(stack)-1]
	//	if root
	//}

	stack := []*TreeNode{}
	var pre *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]

		//}
		//fmt.Println(root.Left == nil && root.Right == nil, root.Right == nil && pre == root.Left, pre == root.Right)
		if (top.Left == nil && top.Right == nil) || (top.Right == nil && pre == top.Left) || pre == top.Right {
			fmt.Print(top.Val)
			pre = top
			stack = stack[:len(stack)-1]
		} else {
			root = top.Right
		}
	}

}

func bfs(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Print(node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

func bfsKLevel(root *TreeNode, level int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	currentLevel := 0
	for len(queue) > 0 {
		if currentLevel > level {
			return
		}
		currentLevelSize := len(queue)
		currentCount := 0
		currentLevel++
		for currentCount < currentLevelSize { // 一次性比遍历完当前层
			currentCount++
			node := queue[0]
			queue = queue[1:]
			if currentLevel == level {
				fmt.Print(node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

	}
}
