package tree


func findTarget(root *Node, k int) bool {
	return dfs(root, k, map[int]bool{})
}

func dfs(root *Node, k int, m map[int]bool) bool {
	if root == nil {
		return false
	}

	if m[k-root.Val] {
		return true
	}
	m[root.Val] = true

	return dfs(root.Left, k, m) || dfs(root.Right, k, m)
}