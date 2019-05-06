/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */
func intersection(nums1 []int, nums2 []int) []int {
	nm1 := make(map[int]struct{})
	nm2 := make(map[int]struct{})
	for _, n := range nums1 {
		nm1[n] = struct{}{}
	}
	ret := []int{}
	for _, n := range nums2 {
		_, ok := nm2[n]
		if ok {
			continue
		}
		nm2[n] = struct{}{}
		_, ok = nm1[n]
		if !ok {
			continue
		}
		ret = append(ret, n)
	}
	return ret
}

