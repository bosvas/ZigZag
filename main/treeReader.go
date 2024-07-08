package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func treeReader() {

	reader := bufio.NewReader(os.Stdin)

	//Enter example is :
	//1 null 1 1 1 null null 1 1 null 1 null null null 1
	//solution - 3
	//1 1 1 null 1 null null 1 1 null 1
	//solution - 4
	//1
	//solution - 0
	//1 1 1
	//solution - 1

	fmt.Println("Ilya, hi! Please enter data in one string in format:'1 null 1 1 1 null null 1 1 null 1 null null null 1':")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	values := strings.Split(input, " ")

	root := buildTree(values)
	if root == nil {
		fmt.Println("You can't make tree from your input")
		return
	}

	length := longestZigZag(root)
	fmt.Printf("Solution to ZigZug Leetcode problem is : %d\n", length)

}

func buildTree(values []string) *TreeNode {

	rootVal, _ := strconv.Atoi(values[0])

	if len(values) == 0 || rootVal != 1 {
		return nil
	}

	root := &TreeNode{Val: rootVal}
	queue := []*TreeNode{root}

	i := 1
	for i < len(values) {
		node := queue[0]
		queue = queue[1:]

		if i < len(values) {
			if values[i] != "null" {
				leftVal, _ := strconv.Atoi(values[i])
				node.Left = &TreeNode{Val: leftVal}
				queue = append(queue, node.Left)
			} else {
				node.Left = nil
			}
			i++
		}

		if i < len(values) {
			if values[i] != "null" {
				rightVal, _ := strconv.Atoi(values[i])
				node.Right = &TreeNode{Val: rightVal}
				queue = append(queue, node.Right)
			} else {
				node.Right = nil
			}
			i++
		}
	}

	return root
}
