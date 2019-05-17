/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	_, d := dobt(root)
	return d - 1
}

func dobt(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	l, dl := dobt(root.Left)
	r, dr := dobt(root.Right)
	d := l + r + 1
	l++
	r++
	ml := l
	if r > l {
		ml = r
	}
	if dl > d {
		d = dl
	}
	if dr > d {
		d = dr
	}
	return ml, d
}

