/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 *
 * https://leetcode-cn.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (48.07%)
 * Total Accepted:    20.3K
 * Total Submissions: 42.1K
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * 设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
 *
 *
 * push(x) -- 将元素 x 推入栈中。
 * pop() -- 删除栈顶的元素。
 * top() -- 获取栈顶元素。
 * getMin() -- 检索栈中的最小元素。
 *
 *
 * 示例:
 *
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> 返回 -3.
 * minStack.pop();
 * minStack.top();      --> 返回 0.
 * minStack.getMin();   --> 返回 -2.
 *
 *
 */
type MinStack struct {
	datas []int
	min   int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if len(this.datas) == 0 {
		this.min = x
	}
	this.datas = append(this.datas, x)
	if x < this.min {
		this.min = x
	}
}

func (this *MinStack) Pop() {
	if len(this.datas) == 0 {
		return
	}
	li := len(this.datas) - 1
	if this.min >= this.datas[li] {
		this.min = this.datas[0]
		for i := 1; i < li; i++ {
			if this.datas[i] < this.min {
				this.min = this.datas[i]
			}
		}
	}
	this.datas = this.datas[:li]
}

func (this *MinStack) Top() int {
	return this.datas[len(this.datas)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
