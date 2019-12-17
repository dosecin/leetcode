/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */

// @lc code=start
func findDuplicate(nums []int) int {
	lef, rig := 1, len(nums)-1
	for lef < rig {
		mid := lef + (rig-lef)/2
		cnt := 0
		for _, num := range nums {
			if num <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			lef = mid + 1
		} else {
			rig = mid
		}
	}
	return rig
}

// @lc code=end
