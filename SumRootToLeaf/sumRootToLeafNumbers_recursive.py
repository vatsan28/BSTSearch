# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def SumHelper(self, node, s, tsum):
        if node is None:
            return 0
    
        if (node.left is None and node.right is None):
            return tsum+int(s+str(node.val))
        
        return self.SumHelper(node.left, s + str(node.val), tsum) + self.SumHelper(node.right, s + str(node.val), tsum)
        
    def sumNumbers(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        return self.SumHelper(root, "", 0)