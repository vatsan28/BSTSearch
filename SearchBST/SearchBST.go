package SearchBST

import "fmt"

type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}

func main() {

	var node1 = new(TreeNode)
	var node11 = new(TreeNode)
	var node12 = new(TreeNode)
	var node111 = new(TreeNode)
	var node112 = new(TreeNode)

	node1.Val = 4
	node1.Right = node12
	node1.Left = node11
	node12.Val = 7
	node11.Val = 2
	node11.Right = node112
	node11.Left = node111
	node111.Val = 1
	node112.Val = 3

	fmt.Println(searchBSTIterative(node1, 2))
	fmt.Println(searchBSTRecursive(node1, 3))
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
