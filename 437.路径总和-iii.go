/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
    if root == nil {
		return 0
	}
	cnt := 0
	if sum == root.Val {
		cnt = 1
	}
	return cnt + pathSum(root.Left, sum) +
		doPathSum(root.Left, sum-root.Val) +
		pathSum(root.Right, sum) +
		doPathSum(root.Right, sum-root.Val)
}

func doPathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	cnt := 0
	if sum == root.Val {
		cnt = 1
	}
	return cnt + doPathSum(root.Left, sum-root.Val) +
		doPathSum(root.Right, sum-root.Val)
}
