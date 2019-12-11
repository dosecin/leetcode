/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
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
			preIdx := 0
			if currFrame > 0 {
				preIdx = stack[currFrame-1] + 1
			}
			if idx == preIdx || candidates[idx] != candidates[idx-1] {
				if candidates[idx] < target {
					target -= candidates[idx]
					stack = append(stack, idx+1)
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
		}
		stack[currFrame]++
	}
	/*递归实现
	for i, v := range candidates {
		if i > 0 && candidates[i-1] == v {
			continue
		}
		if v == target {
			res = append(res, []int{v})
		} else if v < target {
			subRes := combinationSum2(candidates[i+1:], target-v)
			for _, vRes := range subRes {
				res = append(res, append([]int{v}, vRes...))
			}
		}
	}
	*/
	return res
}

// @lc code=end
