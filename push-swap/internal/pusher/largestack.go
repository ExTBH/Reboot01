package pusher

func LargeStack(a, b []int) []int {
	for len(a) != 3 {
		min := min(a)
		index := getIndex(a, min)
		proximity := len(a) / 2
		if proximity > index {
			for min != a[0] {
				if min == a[1] && a[0]-1 == min {
					sa(a)
				} else {
					ra(a)
				}
			}
		} else {
			for min != a[0] {
				rra(a)
			}
		}
		pb(&a, &b)
	}
	SmallStack(a)

	for len(b) != 0 {
		pa(&a, &b)
	}
	return a
}

func getIndex(a []int, n int) int {
	for i := 0; i < len(a); i++ {
		if a[i] == n {
			return i
		}
	}
	return -1
}

func min(a []int) int {
	min := a[0]
	for i := 0; i < len(a); i++ {
		if min > a[i] {
			min = a[i]
		}
	}
	return min
}

func reverse(a []int) {
	for i := len(a) - 1; i > 0; i-- {
		a[i], a[i-1] = a[i-1], a[i]
	}
}
