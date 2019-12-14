/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	pathSum := make([]int, len(triangle), len(triangle))
	minPath := -1
	for i := 0; i < len(triangle); i++ {
		if i == 0 {
			minPath = triangle[0][0]
			pathSum[0] = minPath
			continue
		}
		data := triangle[i]
		pre := pathSum[0]
		for j := 0; j < len(data); j++ {
			if j == 0 {
				minPath = pre + data[0]
				pathSum[0] = minPath
				continue
			}
			if j == len(data)-1 {
				pathSum[j] = pre + data[j]
			} else {
				left := pre + data[j]
				right := pathSum[j] + data[j]
				pre = pathSum[j]
				pathSum[j] = left
				if pathSum[j] > right {
					pathSum[j] = right
				}
			}
			if pathSum[j] < minPath {
				minPath = pathSum[j]
			}
		}
	}
	return minPath
}

// @lc code=end
