package strings

import (
	"fmt"
)

func LongestCommonSubSequence(a, b string) {
	ans := LongestCommonSubSequenceHelper(0,0,a,b)
	fmt.Println(ans)
}


func LongestCommonSubSequenceHelper(i,j int, a,b string) int {
	if i == len(a) || j == len(b) {
		return 0 
	}
	if a[i] == b[j] {
		return 1 + LongestCommonSubSequenceHelper(i+1, j+1, a, b)
	}
	return Max(LongestCommonSubSequenceHelper(i+1,j,a,b), LongestCommonSubSequenceHelper(i,j+1,a,b))
}


func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}


//using dynamic programming
//tushar roy
func LongestCommonSeqDyn(text1,text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m + 1)
	for i := range dp {
		dp[i] = make([]int, n + 1)
	}
	
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i - 1] == text2[j - 1] {
				dp[i][j] = dp[i - 1][j - 1] + 1
			} else {
				dp[i][j] = max(dp[i - 1][j], dp[i][j - 1])
			}
		}
	}
	return dp[m][n]
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

