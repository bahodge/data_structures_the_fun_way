package ch05

import (
	"fmt"
)

type BinarySearchTree struct {
	Root *TreeNode
}

type TreeNode struct {
	Value  int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

func (t BinarySearchTree) FindTreeNode(target int) *TreeNode {
	current := t.Root
	for current != nil && current.Value != target {
		if target < current.Value {
			current = current.Left
		} else {
			current = current.Right
		}

	}
	return current
}

func (t *BinarySearchTree) InsertTreeNode(value int) {
	if t.Root == nil {
		t.Root = &TreeNode{
			Value: value,
		}
	} else {
		InsertNode(t.Root, value)
	}
}

func (t *BinarySearchTree) RemoveTreeNode(value int) {
	node := t.FindTreeNode(value)

	if t.Root == nil || node == nil {
		return
	}

	// Deleting a leaf node
	if node.Left == nil && node.Right == nil {
		if node.Parent == nil {
			t.Root = nil
		} else if node.Parent.Left == node {
			node.Parent.Left = nil
		} else {
			node.Parent.Right = nil
		}

		return
	}

	// Delete a node with one child
	if node.Left == nil || node.Right == nil {
		child := node.Left
		if node.Left == nil {
			child = node.Right
		}
		child.Parent = node.Parent
		if node.Parent == nil {
			t.Root = child
		} else if node.Parent.Left == node {
			node.Parent.Left = child
		} else {
			node.Parent.Right = child
		}
		return
	}

	// Delete a node with 2 children
	successor := node.Right
	for successor.Left != nil {
		successor = successor.Left
	}
	t.RemoveTreeNode(node.Value)

	// Insert the successor in the deleted node's place
	if node.Parent == nil {
		t.Root = successor
	} else if node.Parent.Left == node {
		node.Parent.Left = successor
	} else {
		node.Parent.Right = successor
	}

	successor.Parent = node.Parent
	successor.Left = node.Left
	node.Left.Parent = successor

	successor.Right = node.Right

	if node.Right != nil {
		node.Right.Parent = successor
	}
}

func InsertNode(current *TreeNode, value int) {
	// no duplicates
	if value == current.Value {
		return
	}

	if value < current.Value {
		if current.Left != nil {
			InsertNode(current.Left, value)
		} else {
			current.Left = &TreeNode{
				Value:  value,
				Parent: current,
			}
		}
	} else {
		if current.Right != nil {
			InsertNode(current.Right, value)
		} else {
			current.Right = &TreeNode{
				Value:  value,
				Parent: current,
			}
		}
	}
}

func Run() {
	tree := &BinarySearchTree{}
	tree.InsertTreeNode(10)
	tree.InsertTreeNode(20)
	tree.InsertTreeNode(9)
	tree.InsertTreeNode(10)
	tree.InsertTreeNode(12)

	fmt.Println("----------------ch05 start-----------------")
	fmt.Printf("binary tree FindTreeNode 9. Found %d\n", tree.FindTreeNode(9).Value)
	fmt.Printf("binary tree FindTreeNode 20. Found %d\n", tree.FindTreeNode(20).Value)
	fmt.Println("binary tree RemoveTreeNode 9")
	tree.RemoveTreeNode(9)
	fmt.Printf("binary tree FindTreeNode 9. Found %+v\n", tree.FindTreeNode(9))

	fmt.Println("----------------ch05 end-----------------")
}
