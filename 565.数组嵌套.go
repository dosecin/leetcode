/*
 * @lc app=leetcode.cn id=565 lang=golang
 *
 * [565] 数组嵌套
 */

// @lc code=start
func arrayNesting(nums []int) int {
	ans := 0
	for i := range nums {
		if nums[i] < 0 {
			continue
		}
		cnt, idx := 0, i
		for {
			idx, nums[idx] = nums[idx], -1
			if idx < 0 {
				break
			}
			cnt++
			if cnt > ans {
				ans = cnt
			}
		}
	}
	return ans
}

// @lc code=end
