/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */
type MyStack struct {
	vals []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
    return MyStack{}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
    this.vals = append(this.vals, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	l := len(this.vals)-1
	x := this.vals[l]
	this.vals = this.vals[:l]
	return x
}


/** Get the top element. */
func (this *MyStack) Top() int {
	l := len(this.vals)-1
	return this.vals[l]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    return len(this.vals) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

