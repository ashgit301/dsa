package tree

import (
	//"fmt"
)

func DiagnolTraversal(root *Node, result *[]int) {
	queue := []*Node{}
	node := root
	for node != nil {
		*result = append(*result, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			node = node.Right
		}else if len(queue) != 0{
			node = queue[0]
			queue = queue[1:]
		}else {
			node= nil
		}
	}
}