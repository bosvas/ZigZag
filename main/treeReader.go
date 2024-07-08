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
	//1 2 L
	//1 3 R
	//2 4 R
	//4 5 L
	//4 6 R
	//3 7 L

	fmt.Println("Please enter edges in format 'int int (L/R)' where first - is parent, second is child and L or R is position. Write 'stop' to finish):")

	var edges [][]interface{}

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "stop" {
			break
		}
		parts := strings.Fields(input)
		if len(parts) != 3 {
			fmt.Println("Ooops, you did something wrong, try again (not 3 var in this line's input")
			continue
		}
		parent, err1 := strconv.Atoi(parts[0])
		child, err2 := strconv.Atoi(parts[1])
		side := parts[2]

		if err1 != nil || err2 != nil || (side != "L" && side != "R") {
			fmt.Println("Ooops, you did something wrong, try again (not int int L/R type of input)")
			continue
		}

		edges = append(edges, []interface{}{parent, child, side})
	}

	root := buildTree(edges)
	if root == nil {
		fmt.Println("You can't make tree from your input")
		return
	}

	length := longestZigZag(root)
	fmt.Printf("Solution to ZigZug Leetcode problem is : %d\n", length)

}

func buildTree(edges [][]interface{}) *TreeNode {

	nodes := make(map[int]*TreeNode)
	parents := make(map[int]bool)

	var root *TreeNode

	for _, edge := range edges {
		parentVal, childVal := edge[0].(int), edge[1].(int)
		side := edge[2].(string)

		if parents[childVal] {
			fmt.Printf("Node %d already has a parent.\n", childVal)
			return nil
		}
		parents[childVal] = true

		parent, exists := nodes[parentVal]
		if !exists {
			parent = &TreeNode{Val: parentVal}
			nodes[parentVal] = parent
		}

		child, exists := nodes[childVal]
		if !exists {
			child = &TreeNode{Val: childVal}
			nodes[childVal] = child
		}

		if side == "L" {
			if parent.Left != nil {
				fmt.Printf("Node %d already has Left child.\n", parentVal)
				return nil
			}
			parent.Left = child
		} else if side == "R" {
			if parent.Right != nil {
				fmt.Printf("Node %d already has Right child.\n", parentVal)
				return nil
			}
			parent.Right = child
		}

		if root == nil || root == child {
			root = parent
		}
	}

	return root
}
