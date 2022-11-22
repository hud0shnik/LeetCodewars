#!/usr/bin/python3
from typing import Optional

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def checkTree(self, root: Optional[TreeNode]) -> bool:
        return root.left.val+root.right.val == root.val
        
if __name__ == "__main__":
    s = Solution
    print(s.checkTree(s,TreeNode(10,TreeNode(6),TreeNode(4))))
