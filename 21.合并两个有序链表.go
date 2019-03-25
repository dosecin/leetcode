/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 *
 * https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (53.38%)
 * Total Accepted:    53.4K
 * Total Submissions: 100K
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 *
 * 示例：
 *
 * 输入：1->2->4, 1->3->4
 * 输出：1->1->2->3->4->4
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head *ListNode
	insert := &head

	for l1 != nil || l2 != nil {
		if l1 == nil {
			*insert = l2
			l2 = l2.Next
		} else if l2 == nil {
			*insert = l1
			l1 = l1.Next
		} else {
			if l1.Val < l2.Val {
				*insert = l1
				l1 = l1.Next
			} else {
				*insert = l2
				l2 = l2.Next
			}
		}
		insert = &((*insert).Next)
	}

	return head
}
