#
# @lc app=leetcode.cn id=429 lang=python3
#
# [429] N叉树的层序遍历
#
"""
# Definition for a Node.
class Node:
    def __init__(self, val, children):
        self.val = val
        self.children = children
"""
class Solution:
    def levelOrder(self, root: 'Node') -> List[List[int]]:
        if root == None:
            return []
        nodes = [root]
        ret = []
        while len(nodes) != 0:
            nl = []
            tmp = []
            for n in nodes:
                nl.append(n.val)
                if n.children is not None:
                    tmp.extend(n.children)
            ret.append(nl)
            nodes = tmp
        return ret

