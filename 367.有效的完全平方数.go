/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */
func isPerfectSquare(num int) bool {
	min, max := 0, num
	for min <= max {
		mid := min + (max - min)/2
		if mid*mid == num {
			return true
		}
		if mid*mid > num {
			max = mid-1
		} else {
			min = mid+1
		}
	}
	return false
}

