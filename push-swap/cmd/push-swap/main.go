package main

import (
	"fmt"
	"os"
	"push-swap/internal/pusher"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	arg := os.Args[1]
	numbers := strings.Fields(arg)
	StackA := make(pusher.Stack, len(numbers))
	for i, numStr := range numbers {
		num, err := pusher.ParseInt(numStr)
		if err != nil {
			fmt.Println("Error")
			return
		}
		StackA[i] = num
	}

	var StackB []int
	if pusher.CheckDup(StackA) {
		fmt.Println("Error")
		return
	} else if sort.IntsAreSorted(StackA) {
		return
	} else {
		if len(StackA) <= 3 {
			pusher.SmallStack(StackA)
		} else {
			pusher.LargeStack(StackA, StackB)
		}
	}
}
