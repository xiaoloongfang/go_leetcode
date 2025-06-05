package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	res := 0
	last_list := make([]*TreeNode, 0)
	next_list := make([]*TreeNode, 0)
	last_list = append(last_list, root)

	for {
		res = 0
		for i := 0; i < len(last_list); i++ {
			node := last_list[i]
			res += node.Val
			if node.Left != nil {
				next_list = append(next_list, node.Left)
			}
			if node.Right != nil {
				next_list = append(next_list, node.Right)
			}
		}
		if len(next_list) == 0 {
			return res
		}
		last_list = next_list
		next_list = make([]*TreeNode, 0)
	}

}
