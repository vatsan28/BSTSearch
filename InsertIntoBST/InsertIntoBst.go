package InsertIntoBST

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

	fmt.Println(insertIntoBSTIterative(node1, 5))
	fmt.Println(insertIntoBSTRecursive(node1, 10))
}


func insertIntoBSTIterative(root *TreeNode, val int) *TreeNode {
	nodeToInsert := TreeNode{Val: val}
	curr := root
	for curr != nil {
		if curr.Val > val {
			if curr.Left != nil {
				curr = curr.Left
			} else {
				curr.Left = &nodeToInsert
				curr = nil
			}
		} else if curr.Val < val {
			if curr.Right != nil {
				curr = curr.Right
			} else {
				curr.Right = &nodeToInsert
				curr = nil
			}
		}
	}
	return root
}

func insertIntoBSTRecursive(root *TreeNode, val int) *TreeNode {
	nodeToInsert := TreeNode{}
	if root == nil {
		root.Val = val
	} else {
		if val < root.Val {
			if root.Left == nil {
				nodeToInsert.Val = val
				root.Left = &nodeToInsert
				return root
			}else{
				root.Left = insertIntoBSTRecursive(root.Left, val)
			}
		} else {
			if root.Right == nil {
				nodeToInsert.Val = val
				root.Right = &nodeToInsert
				return root
			}else{
				root.Right = insertIntoBSTRecursive(root.Right, val)
			}
		}
	}
	return root
}

