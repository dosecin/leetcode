/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */
func containsNearbyDuplicate(nums []int, k int) bool {
	numk := map[int]struct{}{}
	if k > len(nums) {
		k = len(nums)
	}
	var i int
	for i < k {
		if _, exist := numk[nums[i]]; exist {
			return true
		}
		numk[nums[i]] = struct{}{}
		i++
	}

	for i < len(nums) {
		if _, exist := numk[nums[i]]; exist {
			return true
		}
		numk[nums[i]] = struct{}{}
		delete(numk, nums[i-k])
		i++
	}
	return false
}

