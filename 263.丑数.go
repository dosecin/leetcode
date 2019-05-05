/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */
func isUgly(num int) bool {
	if num == 0 {
		return false
	}
	for _, p := range []int{2, 3, 5} {
		for num%p == 0 {
			num /= p
		}
	}
	return num == 1
}
