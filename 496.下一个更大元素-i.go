/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	res := make([]int, len(nums1))
	for i := range nums2 {
		m[nums2[i]] = -1
		for j := i + 1; j < len(nums2); j++ {
			if nums2[j] > nums2[i] {
				m[nums2[i]] = nums2[j]
				break
			}
		}
	}
	for i := range nums1 {
		res[i] = m[nums1[i]]
	}
	return res
}

