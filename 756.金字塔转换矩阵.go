/*
 * @lc app=leetcode.cn id=756 lang=golang
 *
 * [756] 金字塔转换矩阵
 */

// @lc code=start
func pyramidTransition(bottom string, allowed []string) bool {
	mapAllowed := map[string][]byte{}
	for _, a := range allowed {
		pre := a[:2]
		mapAllowed[pre] = append(mapAllowed[pre], a[2])
	}

	var check func(curr string, index int, list []byte) bool
	check = func(curr string, index int, list []byte) bool {
		if len(curr) == 2 {
			_, ok := mapAllowed[curr]
			return ok
		}

		if index+1 == len(curr) {
			return check(string(list), 0, make([]byte, len(list)-1))
		}

		pre := curr[index : index+2]
		if al, ok := mapAllowed[pre]; ok {
			for _, a := range al {
				list[index] = a
				if check(curr, index+1, list) {
					return true
				}
			}
		}
		return false
	}
	return check(bottom, 0, make([]byte, len(bottom)-1))
}

// @lc code=end
