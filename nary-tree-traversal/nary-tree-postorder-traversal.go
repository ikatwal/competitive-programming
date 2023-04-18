package main

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

// Solved 590
func postorder(root *Node) []int {
	res := make([]int, 0)
	helperPostorder(root, &res)
	return res
}

func helperPostorder(root *Node, res *[]int) {
	if root == nil {
		return
	}
	for i := range root.Children {
		helperPostorder(root.Children[i], res)
	}
	*res = append(*res, root.Val)
}
