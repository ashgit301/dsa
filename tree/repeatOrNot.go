package tree
import (
	"fmt"
)
//wrong solution - where is the false condition
func RepeatOrNot(root *Node) bool {
	if root == nil {
		return true
	}
	if IsSubTree(root, root.Left) ||  IsSubTree(root, root.Left) {
		return true
	}

	return RepeatOrNot(root.Left) || RepeatOrNot(root.Right)
}




func findDuplicateSubtrees(root *Node) []*Node {
	hashAll := map[string]int{}
	duplicate := []*Node{}
	dfsForRepeat(root, hashAll, &duplicate)
	return duplicate
}

func dfsForRepeat(node *Node, hashAll map[string]int, duplicate *[]*Node) string {
	if node == nil {
		return "nil"
	}
    lString := dfsForRepeat(node.Left, hashAll, duplicate)
    rString := dfsForRepeat(node.Right, hashAll, duplicate)
    buildString := fmt.Sprintf("(%s)(%v)(%s)", lString, node.Val, rString) //bottom up approach
	hashAll[buildString]++
	if hashAll[buildString] == 2 {
		*duplicate = append(*duplicate, node)
	}
    return buildString
}