package main

import (
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
// bfs
func invertTreeBFS(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	stack := []*TreeNode{root}
	for len(stack) != 0 {
		node := stack[0]
		stack = stack[1:]
		if node == nil {
			continue
		}
		node.Left, node.Right = node.Right, node.Left
		stack = append(stack, node.Left, node.Right)
	}
	return root
}

// dfs
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			node.Left, node.Right = node.Right, node.Left
			continue
		}
		// node != nil
		stack = append(stack, node)
		stack = append(stack, nil)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		//// node != nil
		//if node.Left != nil {
		//	stack = append(stack, node.Left)
		//}
		//stack = append(stack, node)
		//stack = append(stack, nil)
		//if node.Right != nil {
		//	stack = append(stack, node.Right)
		//}
		//
		//// node != nil
		//if node.Left != nil {
		//	stack = append(stack, node.Left)
		//}
		//if node.Right != nil {
		//	stack = append(stack, node.Right)
		//}
		//stack = append(stack, node)
		//stack = append(stack, nil)
	}

	return root
}

var (
	result = []int{}
)
