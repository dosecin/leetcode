/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
import "sort"

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	if res == target {
		return target
	}
	for k, v := range nums {
		i, j := k+1, len(nums)-1
		for i < j {
			newRes := v + nums[i] + nums[j]
			if newRes == target {
				return target
			}
			if newRes < target {
				i++
			} else {
				j--
			}
			if abs(newRes-target) < abs(res-target) {
				res = newRes
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end
