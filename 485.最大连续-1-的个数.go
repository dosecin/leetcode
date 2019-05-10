/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续1的个数
 */
func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	cnt := 0
	for _, v := range nums {
		if v == 1 {
			cnt++
			continue
		}
		if cnt > max {
			max = cnt
		}
		cnt = 0
	}
	if cnt > max {
		max = cnt
	}
	return max
}

