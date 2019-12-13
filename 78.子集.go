/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	ans := [][]int{}
	subans := subsets(nums[1:])
	ans = append(ans, subans...)
	for _, sa := range subans {
		sa = append([]int{nums[0]}, sa...)
		ans = append(ans, sa)
	}
	return ans
}

// @lc code=end
