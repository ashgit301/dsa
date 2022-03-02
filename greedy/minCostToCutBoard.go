package greedy

import (
	"sort"
)

func minCostTocutBoard(a,b []int) {
	i := 0 
	j := 0 
	sort.Slice(a, func(i, j int)bool {
		return a[i] > a[j]
	})
	sort.Slice(b, func(i, j int)bool {
		return b[i] > b[j]
	})
	rowEdge := len(a)-1
	colEdge := len(b)-1
	horizontalPiece := 1
	verticalPiece := 1
	cost := 0 
	for i < rowEdge && j < colEdge {
		if a[i]>= b[j] {
			cost += a[i]*verticalPiece
			horizontalPiece++
			i++
		}else {
			cost += b[j]*horizontalPiece
			verticalPiece++
			j++
		}
	}
	for i < rowEdge {
		cost += a[i]*verticalPiece
		horizontalPiece++
	}
	for j < colEdge {
		cost += b[j]*horizontalPiece
		verticalPiece++
	}
}