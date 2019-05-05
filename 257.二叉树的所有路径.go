import "strconv"

/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	nstr := strconv.Itoa(root.Val)
	lstr := binaryTreePaths(root.Left)
	rstr := binaryTreePaths(root.Right)
	str := make([]string, 0, len(lstr)+len(rstr))
	for i := range lstr {
		str = append(str, nstr+"->"+lstr[i])
	}
	for i := range rstr {
		str = append(str, nstr+"->"+rstr[i])
	}
	return str
}
