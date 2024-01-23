package pusher

func BubbleSort(a []int) []int {
	k := make([]int, len(a))
	copy(k, a)

	for i := 0; i < len(k)-1; i++ {
		swapped := false
		for j := 0; j < len(k)-i-1; j++ {
			if k[j+1] < k[j] {
				k[j], k[j+1] = k[j+1], k[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	return k
}
