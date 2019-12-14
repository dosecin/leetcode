/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	val := postorder[len(postorder)-1]
	root := &TreeNode{
		Val: val,
	}
	rootIdx := 0
	for rootIdx < len(inorder) {
		if val == inorder[rootIdx] {
			break
		}
		rootIdx++
	}
	root.Left = buildTree(inorder[:rootIdx], postorder[:rootIdx])
	root.Right = buildTree(inorder[rootIdx+1:], postorder[rootIdx:len(postorder)-1])
	return root
}

// @lc code=end
