# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def sumNumbers(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        if (root is None):
            return 0
        
        numbers = []
        m = [(root, "")]
        while(len(m)!=0):
            node, number = m.pop()
            if (node.left):
                m.append((node.left, number+str(node.val)))
            
            if (node.right):
                m.append((node.right, number+str(node.val)))
            
            if (node.left is None and node.right is None):
                numbers.append(int(number+str(node.val)))

        return sum(numbers)