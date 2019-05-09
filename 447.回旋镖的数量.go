/*
 * @lc app=leetcode.cn id=447 lang=golang
 *
 * [447] 回旋镖的数量
 */
func numberOfBoomerangs(points [][]int) int {
	res := 0
	for i := range points {
		disCnt := map[int]int{}
		for j := range points {
			if i == j {
				continue
			}
			x := points[i][0] - points[j][0]
			y := points[i][1] - points[j][1]
			dis := x * x + y * y;
			disCnt[dis]++
		}
		for _, v := range disCnt {
			res += v*(v-1)
		}
	}
	return res
}

