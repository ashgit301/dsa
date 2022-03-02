package greedy

import(
	"fmt"
	"sort"
)
type Meeting struct {
	Start int
	End int
}

func MeetingRoom(start, end []int, N int) {
	intervalsOrderedByStart := make([]Meeting, N)
	for i := 0 ; i < len(intervalsOrderedByStart) ; i ++ {
		intervalsOrderedByStart[i].Start , intervalsOrderedByStart[i].End = start[i],end[i]
	}
	sort.Slice(intervalsOrderedByStart, func (i,j int)bool {
		return intervalsOrderedByStart[i].End <= intervalsOrderedByStart[j].End
	})
	fmt.Println(intervalsOrderedByStart)
	count := 1
	limit := intervalsOrderedByStart[0].End
	for i := 1 ; i < N ; i ++ {
		if intervalsOrderedByStart[i].Start > limit{
			count++
			limit = intervalsOrderedByStart[i].End
		}
	}
	fmt.Println(count)
}