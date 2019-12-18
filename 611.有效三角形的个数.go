/*
 * @lc app=leetcode.cn id=611 lang=golang
 *
 * [611] 有效三角形的个数
 */

// @lc code=start
import "sort"

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := 0; i < len(nums); i++ {
		j, k := i+1, i+2
		for k < len(nums) {
			sum := nums[i] + nums[j]
			if sum > nums[k] {
				ans += k - j
				k++
			} else {
				j++
				if j == k {
					k++
				}
			}
		}
	}
	return ans
}

// @lc code=end
