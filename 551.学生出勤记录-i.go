/*
 * @lc app=leetcode.cn id=551 lang=golang
 *
 * [551] 学生出勤记录 I
 */
func checkRecord(s string) bool {
	absent := 0
	late := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			absent++
			if absent > 1 {
				return false
			}
			late = 0
		} else if s[i] == 'L' {
			late++
			if late > 2 {
				return false
			}
		} else {
			late = 0
		}
	}
	return true
}

