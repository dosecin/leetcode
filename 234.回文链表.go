/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	vals := []int{}
	n := head
	for n != nil {
		vals = append(vals, n.Val)
		n = n.Next
	}
	for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
		if vals[i] != vals[j] {
			return false
		}
	}
	return true
}

