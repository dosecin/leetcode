/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	i, j := 0, len(height)-1
	calcArea := func(left, right int) int {
		if height[left] > height[right] {
			return (right - left) * height[right]
		}
		return (right - left) * height[left]
	}
	area := calcArea(i, j)
	for i < j {
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
		ta := calcArea(i, j)
		if ta > area {
			area = ta
		}
	}
	return area
}

// @lc code=end
