/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	mapIdx := map[int]int{}
	for i, num := range nums {
		if preIdx, ok := mapIdx[num]; ok {
			if i <= preIdx+k {
				return true
			}
		}
		mapIdx[num] = i
	}
	return false
}

// @lc code=end
