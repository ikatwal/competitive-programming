package main

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

// Solved 589
func preorder(root *Node) []int {
	// print root
	// do preorder of its children
	res := make([]int, 0)
	helperPreorder(root, &res)
	return res
}

func helperPreorder(root *Node, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	for i := range root.Children {
		helperPreorder(root.Children[i], res)
	}
}
