/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
import "sort"

func subsetsWithDup(nums []int) [][]int {
	var subset func([]int) [][]int
	subset = func(nums []int) [][]int {
		if len(nums) == 0 {
			return [][]int{{}}
		}
		ans := [][]int{}
		prefix := [][]int{}
		subIdx := 0
		for subIdx < len(nums) {
			if nums[0] == nums[subIdx] {
				subIdx++
				prefix = append(prefix, append([]int{}, nums[:subIdx]...))
				continue
			}
			break
		}
		subans := subset(nums[subIdx:])
		ans = append(ans, subans...)
		for _, sa := range subans {
			for _, pre := range prefix {
				pre = append([]int{}, pre...)
				pre = append(pre, sa...)
				ans = append(ans, pre)
			}
		}
		return ans
	}
	sort.Ints(nums)
	return subset(nums)
}

// @lc code=end
