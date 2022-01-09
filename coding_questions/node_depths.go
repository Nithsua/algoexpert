package main

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func TraverseTree(root *BinaryTree, count *int, level int) int {
	if root == nil {
		return 0
	} else {
		*count += TraverseTree(root.Left, count, level+1)
		*count += TraverseTree(root.Right, count, level+1)
		return level
	}
}

func NodeDepths(root *BinaryTree) int {
	count := 0
	TraverseTree(root, &count, 0)
	return count
}
