/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	stack := []int{-1}
	maxArea := 0
	for i, h := range heights {
		if len(stack) == 1 {
			stack = append(stack, i)
			continue
		}
		end := stack[len(stack)-1]
		if h >= heights[end] {
			stack = append(stack, i)
			continue
		}
		for len(stack) > 1 {
			stackLast := len(stack) - 1
			end := stack[stackLast]
			if h >= heights[end] {
				break
			}
			area := heights[end] * (i - stack[stackLast-1] - 1)
			if area > maxArea {
				maxArea = area
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for i := 1; i < len(stack); i++ {
		idx := stack[i]
		area := heights[idx] * (len(heights) - 1 - stack[i-1])
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

// @lc code=end
