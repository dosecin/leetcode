/*
 * @lc app=leetcode.cn id=729 lang=golang
 *
 * [729] 我的日程安排表 I
 */

// @lc code=start
type MyCalendar struct {
	cals [][2]int
}

func Constructor() MyCalendar {
	return MyCalendar{
		cals: [][2]int{},
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for _, cal := range this.cals {
		if start < cal[1] && end > cal[0] {
			return false
		}
	}
	this.cals = append(this.cals, [2]int{start, end})
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
// @lc code=end
