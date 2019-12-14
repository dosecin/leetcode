/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	pm, pn, curr := m-1, n-1, len(nums1)-1
	for pm >= 0 && pn >= 0 {
		if nums1[pm] >= nums2[pn] {
			nums1[curr] = nums1[pm]
			pm--
		} else {
			nums1[curr] = nums2[pn]
			pn--
		}
		curr--
	}
	for pn >= 0 {
		nums1[curr] = nums2[pn]
		pn--
		curr--
	}
}

// @lc code=end
