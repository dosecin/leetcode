/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	stack := []*TreeNode{root}
	valSli := []int{}
	for len(stack) > 0 {
		n := stack[0]
		stack = stack[1:]
		valSli = append(valSli, n.Val)
		if n.Left != nil {
			stack = append(stack, n.Left)
		}
		if n.Right != nil {
			stack = append(stack, n.Right)
		}
	}
	if len(valSli) < 2 {
		return 0
	}
	sort.Ints(valSli)
	diff := absIntSub(valSli[0], valSli[1])
	for i := 2; i < len(valSli); i++ {
		d := absIntSub(valSli[i], valSli[i-1])
		if d < diff {
			diff = d
		}
	}
	return diff
}

func absIntSub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

