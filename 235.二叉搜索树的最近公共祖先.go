/*
 * @lc app=leetcode.cn id=235 lang=golang
 *
 * [235] 二叉搜索树的最近公共祖先
 */
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if (root.Val-p.Val)*(root.Val-q.Val) <= 0 {
			return root
		}
		if root.Val < p.Val {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return nil
}
