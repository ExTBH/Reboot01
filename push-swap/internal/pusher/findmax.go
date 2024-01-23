package pusher

import "sort"

func Findmax(StackA []int) (int, int, int, int) {
	var sortedStack []int
	sortedStack = append(sortedStack, StackA...)
	sort.Ints(sortedStack)

	min_num := sortedStack[0]
	med_num := sortedStack[1]
	max_num := sortedStack[2]
	Size := len(StackA)
	return min_num, med_num, Size, max_num
}
