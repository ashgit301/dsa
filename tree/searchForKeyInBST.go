package tree

func searchBST(root *Node, val int) *Node {
	if root == nil  {
        return nil
    }
    if val < root.Val {
        return searchBST(root.Left, val)
    } else if val > root.Val {
        return searchBST(root.Right, val)
    } else if val == root.Val {
        return root
    }
    return nil

}