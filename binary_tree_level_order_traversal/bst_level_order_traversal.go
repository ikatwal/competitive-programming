package mains

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Solved 102
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	currLevel := make([]*TreeNode, 0)
	currLevel = append(currLevel, root)
	for len(currLevel) != 0 {
		// store each level in currLevel
		temp := make([]int, 0)
		nextLevel := make([]*TreeNode, 0)
		// as long as we traverse currLevel we get all elems of a level
		for range currLevel {
			currRoot := currLevel[0]
			temp = append(temp, currRoot.Val)
			if currRoot.Left != nil {
				nextLevel = append(nextLevel, currRoot.Left)
			}
			if currRoot.Right != nil {
				nextLevel = append(nextLevel, currRoot.Right)
			}
			// pop left element in queue
			currLevel[0] = nil
			currLevel = currLevel[1:]
		}
		res = append(res, temp)
		// assign currLevel to nextLevel
		currLevel = nextLevel
	}

	return res
}
