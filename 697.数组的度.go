/*
 * @lc app=leetcode.cn id=697 lang=golang
 *
 * [697] 数组的度
 */

// @lc code=start
func findShortestSubArray(nums []int) int {
	lef, rig, cntMap := map[int]int{}, map[int]int{}, map[int]int{}
	maxCnt := 0
	for i, num := range nums {
		cntMap[num]++
		cnt := cntMap[num]
		if cnt > maxCnt {
			maxCnt = cnt
		}
		lefIdx, rigIdx := i, len(nums)-1-i
		if _, ok := lef[num]; !ok {
			lef[num] = lefIdx
		}
		rigNum := nums[rigIdx]
		if _, ok := rig[rigNum]; !ok {
			rig[rigNum] = rigIdx
		}
	}
	ans := len(nums)
	for num, cnt := range cntMap {
		if cnt != maxCnt {
			continue
		}
		subLen := rig[num] - lef[num] + 1
		if subLen < ans {
			ans = subLen
		}
	}
	return ans
}

// @lc code=end
