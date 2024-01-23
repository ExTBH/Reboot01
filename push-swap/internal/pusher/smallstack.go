package pusher

import (
	"fmt"
)

func SmallStack(StackA []int) {
	if len(StackA) == 2 {
		fmt.Println("sa")
	} else {
		min, med, _, max := Findmax(StackA)
		if med == StackA[0] && min == StackA[1] && max == StackA[2] {
			fmt.Println("sa")
		} else if max == StackA[0] && med == StackA[1] && min == StackA[2] {
			fmt.Println("sa")
			fmt.Println("rra")
		} else if max == StackA[0] && min == StackA[1] && med == StackA[2] {
			fmt.Println("ra")
		} else if min == StackA[0] && max == StackA[1] && med == StackA[2] {
			fmt.Println("sa")
			fmt.Println("ra")
		} else if med == StackA[0] && max == StackA[1] && min == StackA[2] {
			fmt.Println("rra")
		}
	}
}
