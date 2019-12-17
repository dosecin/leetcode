/*
 * @lc app=leetcode.cn id=228 lang=golang
 *
 * [228] 汇总区间
 */

// @lc code=start
import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	ans := []string{}
	lef, rig := 0, 1
	for rig < len(nums) {
		if nums[rig] > nums[rig-1]+1 {
			if lef == rig-1 {
				ans = append(ans, fmt.Sprintf("%d", nums[lef]))
			} else {
				ans = append(ans, fmt.Sprintf("%d->%d", nums[lef], nums[rig-1]))
			}
			lef = rig
		}
		rig++
	}
	if lef == rig-1 {
		ans = append(ans, fmt.Sprintf("%d", nums[lef]))
	} else {
		ans = append(ans, fmt.Sprintf("%d->%d", nums[lef], nums[rig-1]))
	}
	return ans
}

// @lc code=end
