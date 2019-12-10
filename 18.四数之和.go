/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start

import "sort"

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	if len(nums) < 4 {
		return res
	}
	sort.Ints(nums)
	for m, v := range nums {
		if m > 0 && nums[m] == nums[m-1] {
			continue
		}
		for n := m + 1; n < len(nums); n++ {
			if n > m+1 && nums[n] == nums[n-1] {
				continue
			}
			i, j := n+1, len(nums)-1
			for i < j {
				sum := v + nums[n] + nums[i] + nums[j]
				if sum == target {
					res = append(res, []int{v, nums[n], nums[i], nums[j]})
					for i < j && nums[i] == nums[i+1] {
						i++
					}
					for i < j && nums[j] == nums[j-1] {
						j--
					}
					i++
					j--
				} else if sum < target {
					i++
				} else {
					j--
				}
			}
		}
	}
	return res
}

// @lc code=end
