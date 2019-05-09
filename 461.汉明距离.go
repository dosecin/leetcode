/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */
func hammingDistance(x int, y int) int {
	z := x ^ y
	res := 0
	for z != 0 {
		res++
		z &= z-1
	}
	return res
}

