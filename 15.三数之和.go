/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
import "sort"

func threeSum(nums []int) [][]int {
	sortNums := make([]int, len(nums), len(nums))
	copy(sortNums, nums)
	sort.Ints(sortNums)
	res := make([][]int, 0, 0)
	for k, v := range sortNums {
		i, j := k+1, len(sortNums)-1
		if k > 0 && sortNums[k] == sortNums[k-1] {
			continue
		}
		for i < j {
			sum := v + sortNums[i] + sortNums[j]
			if sum == 0 {
				res = append(res, []int{v, sortNums[i], sortNums[j]})
				for i < j && sortNums[i+1] == sortNums[i] {
					i++
				}
				for i < j && sortNums[j-1] == sortNums[j] {
					j--
				}
				i++
				j--
			} else if sum < 0 {
				i++
			} else {
				j--
			}
		}
	}
	return res
}

// @lc code=end
