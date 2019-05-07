#
# @lc app=leetcode.cn id=427 lang=python3
#
# [427] 建立四叉树
#
"""
# Definition for a QuadTree node.
class Node:
    def __init__(self, val, isLeaf, topLeft, topRight, bottomLeft, bottomRight):
        self.val = val
        self.isLeaf = isLeaf
        self.topLeft = topLeft
        self.topRight = topRight
        self.bottomLeft = bottomLeft
        self.bottomRight = bottomRight
"""
class Solution:
    def is_leaf(self, grid: List[List[int]]) -> bool:
        return all(grid[i][j] == grid[0][0] 
            for i in range(len(grid)) for j in range(len(grid[i])))                
                
    def construct(self, grid: List[List[int]]) -> 'Node':
        if self.is_leaf(grid):
            return Node(True if grid[0][0] else False, True, None, None, None, None)
        half = len(grid) // 2
        return Node('*', False,
                        self.construct([row[:half] for row in grid[:half]]),
                        self.construct([row[half:] for row in grid[:half]]),
                        self.construct([row[:half] for row in grid[half:]]),
                        self.construct([row[half:] for row in grid[half:]]))

