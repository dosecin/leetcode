/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	doConvertBST(root, 0)
	return root
}

func doConvertBST(root *TreeNode, pVal int) int {
	if root == nil {
		return pVal
	}
	r := doConvertBST(root.Right, pVal)
	root.Val += r
	if root.Left == nil {
		return root.Val
	}
	return doConvertBST(root.Left, root.Val)
}

