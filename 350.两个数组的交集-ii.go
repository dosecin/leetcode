/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */
func intersect(nums1 []int, nums2 []int) []int {
	nm1 := make(map[int]int)
	for _, n := range nums1 {
		nm1[n]++
	}
	ret := make([]int, 0)
	for _, n := range nums2 {
		k, ok := nm1[n]
		if !ok {
			continue
		}
		k--
		if k <= 0 {
			delete(nm1, n)
		} else {
			nm1[n] = k
		}
		ret = append(ret, n)
	}
	return ret
}

