package SearchBST

type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}

func searchBSTIterative(root *TreeNode, val int) *TreeNode {

	for root != nil {
		if root.Val < val {
			root = root.Right
		} else if root.Val > val {
			root = root.Left
		} else {
			return root
		}
	}
	return nil
}

func searchBSTRecursive(root *TreeNode, val int) *TreeNode {
	cur := root
	if cur == nil {
		return nil
	} else if cur.Val > val {
		return searchBSTRecursive(cur.Left, val)
	} else if cur.Val < val {
		return searchBSTRecursive((cur.Right, val)
	}
	return cur
}
