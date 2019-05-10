/*
 * @lc app=leetcode.cn id=501 lang=golang
 *
 * [501] 二叉搜索树中的众数
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	m := map[int]int{}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		tn := stack[0]
		stack = stack[1:]
		m[tn.Val]++
		if tn.Left != nil {
			stack = append(stack, tn.Left)
		}
		if tn.Right != nil {
			stack = append(stack, tn.Right)
		}
	}
	max, res := 0, []int{}
	for k, v := range m {
		if v > max {
			res = res[:0]
			res = append(res, k)
			max = v
		} else if v == max {
			res = append(res, k)
		}
	}
	return res
}

