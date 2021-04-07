package q0863

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

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	result := []int{}
	findDisWithTarget(root, target, K, &result)
	return result
}

func findDisWithTarget(root *TreeNode, target *TreeNode, K int, result *[]int) int {
	if root == nil {
		return -1
	}
	if root.Val == target.Val {
		findChidK(root, K, result)
		return K - 1 // 剩余距离
	}
	leftDistance := findDisWithTarget(root.Left, target, K, result)
	if leftDistance == 0 {
		//findChidK(root, leftDistance, result)
		*result = append(*result, root.Val)
	}
	if leftDistance > 0 {
		findChidK(root.Right, leftDistance-1, result)
		return leftDistance - 1
	}

	rightDistance := findDisWithTarget(root.Right, target, K, result)
	if rightDistance == 0 {
		//findChidK(root, rightDistance, result)
		*result = append(*result, root.Val)
	}
	if rightDistance > 0 {
		findChidK(root.Left, rightDistance-1, result)
		return rightDistance - 1
	}

	return -1
}

func findChidK(root *TreeNode, K int, result *[]int) {
	if root == nil {
		return
	}
	if K == 0 {
		*result = append(*result, root.Val)
		return
	}
	findChidK(root.Left, K-1, result)
	findChidK(root.Right, K-1, result)
}

// 父节点
func distanceK2(root *TreeNode, target *TreeNode, K int) []int {
	parentMap := map[*TreeNode]*TreeNode{}
	targetNode := dfs(root, parentMap, target)
	result := []int{}
	findDisK(targetNode, parentMap, K, &result)
	return result
}

func dfs(root *TreeNode, parentMap map[*TreeNode]*TreeNode, target *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var targetNode *TreeNode

	if root.Left != nil {
		t := dfs(root.Left, parentMap, target)
		if t != nil {
			targetNode = t
		}
		parentMap[root.Left] = root
	}
	if root.Right != nil {
		t := dfs(root.Right, parentMap, target)
		if t != nil {
			targetNode = t
		}
		parentMap[root.Right] = root
	}
	if target.Val == root.Val {
		return root
	}
	return targetNode
}

var (
	visit = map[*TreeNode]struct{}{}
)

func findDisK(root *TreeNode, parentMap map[*TreeNode]*TreeNode, k int, result *[]int) {
	if root == nil || k < 0 {
		return
	}
	if _, found := visit[root]; !found {
		visit[root] = struct{}{}
	} else {
		return
	}
	if k == 0 {
		*result = append(*result, root.Val)
	}

	findDisK(parentMap[root], parentMap, k-1, result)
	findDisK(root.Left, parentMap, k-1, result)
	findDisK(root.Right, parentMap, k-1, result)
}
