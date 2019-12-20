import "sort"

/*
 * @lc app=leetcode.cn id=719 lang=golang
 *
 * [719] 找出第 k 小的距离对
 */

// @lc code=start
func smallestDistancePair(nums []int, k int) int {
	possible := func(diff int) bool {
		var lef, rig int
		cnt := 0
		for rig < len(nums) {
			for nums[rig]-nums[lef] > diff {
				lef++
			}
			cnt += rig - lef
			rig++
		}
		return cnt >= k
	}

	sort.Ints(nums)
	lef, rig := 0, nums[len(nums)-1]-nums[0]
	for lef < rig {
		mid := lef + (rig-lef)/2
		if possible(mid) {
			rig = mid
		} else {
			lef = mid + 1
		}
	}
	return lef
}

// @lc code=end
