package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// main task functionality
func longestZigZag(root *TreeNode) int {
	maxLen := 0

	//We can't make time complexity less then О(n) because we need to go through all graphs anyway
	//Solution with recursion seems optimal because it helps us to go through all the nodes one time exactly
	// and the code is much clean and easy to understand
	if root != nil {
		checkPath(root.Left, "right", 1, &maxLen)
		checkPath(root.Right, "left", 1, &maxLen)
	}

	return maxLen
}

// Function for getting the longest zigzag from current node
// We call this function recursively
func checkPath(node *TreeNode, direction string, length int, maxLen *int) {
	if node == nil {
		return
	}
	if length > *maxLen {
		*maxLen = length
	}

	if direction == "left" {
		checkPath(node.Left, "right", length+1, maxLen)
		checkPath(node.Right, "left", 1, maxLen)
	} else {
		checkPath(node.Right, "left", length+1, maxLen)
		checkPath(node.Left, "right", 1, maxLen)
	}
}
