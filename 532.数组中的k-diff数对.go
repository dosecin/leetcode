/*
 * @lc app=leetcode.cn id=532 lang=golang
 *
 * [532] 数组中的K-diff数对
 */

// @lc code=start
func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	mapNums := map[int]int{}
	ans := 0
	for _, num := range nums {
		if cnt, ok := mapNums[num]; ok {
			if k == 0 && cnt == 1 {
				ans++
			}
			mapNums[num]++
			continue
		}
		diff := num + k
		if _, ok := mapNums[diff]; ok {
			ans++
		}
		diff = num - k
		if _, ok := mapNums[diff]; ok {
			ans++
		}
		mapNums[num] = 1
	}
	return ans
}

// @lc code=end
