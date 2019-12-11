/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0, 0)
	stack := []int{0}
	currFrame := 0
	for {
		if stack[currFrame] >= len(candidates) {
			stack = stack[:currFrame]
			currFrame--
			if currFrame < 0 {
				break
			}
			target += candidates[stack[currFrame]]
		} else {
			idx := stack[currFrame]
			if candidates[idx] < target {
				target -= candidates[idx]
				stack = append(stack, idx)
				currFrame++
				continue
			}
			if candidates[idx] == target {
				sumArr := []int{}
				for _, s := range stack {
					sumArr = append(sumArr, candidates[s])
				}
				res = append(res, sumArr)
			}
		}
		stack[currFrame]++
	}
	/*递归实现
	for i, v := range candidates {
		if v == target {
			res = append(res, []int{v})
		} else if v < target {
			subRes := combinationSum(candidates[i:], target-v)
			for _, vRes := range subRes {
				res = append(res, append([]int{v}, vRes...))
			}
		}
	}
	*/
	return res
}

// @lc code=end
