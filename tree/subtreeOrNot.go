package tree

func IsSubTree(main , sub *Node) bool {
	if main == nil {
		return true
	}
	if sub == nil {
		return true
	}
	if isSameTree(main,sub) {
		return true
	}
	return IsSubTree(main.Left, sub) || IsSubTree(main.Right, sub)

}