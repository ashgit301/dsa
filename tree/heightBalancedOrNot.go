package tree


func isBalanced(root *Node) bool {
    
    _, is_balanced := helper(root)
    return is_balanced
}

func helper(node *Node) (int, bool) {
    if node == nil {
        return 0, true
    }
    
    left_h, left_is_balanced := helper(node.Left)
    right_h, right_is_balanced := helper(node.Right)
    
    height := max(left_h, right_h) + 1
    if !left_is_balanced || !right_is_balanced {
        return height, false
    }
    
    if abs(left_h - right_h) > 1 {
        return height, false
    }
    
    return height, true
}

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}