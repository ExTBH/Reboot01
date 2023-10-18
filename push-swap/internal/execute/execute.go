package execute

import (
	"fmt"
	"push-swap/internal/structs"
)

func Start(Rui *structs.Rui) {
	for _, instruction := range Rui.Instructions {
		switch *instruction {
		case structs.InstructionPA:
			pa(Rui)
		case structs.InstructionPB:
			pb(Rui)
		case structs.InstructionSA:
			sa(Rui)
		case structs.InstructionSB:
			sb(Rui)
		case structs.InstructionSS:
			ss(Rui)
		case structs.InstructionRA:
			ra(Rui)
		case structs.InstructionRB:
			rb(Rui)
		case structs.InstructionRR:
			rr(Rui)
		case structs.InstructionRRA:
			rra(Rui)
		case structs.InstructionRRB:
			rrb(Rui)
		case structs.InstructionRRR:
			rrr(Rui)
		}
	}
}

func pa(Rui *structs.Rui) {
	Rui.StackB, Rui.StackA = push(Rui.StackB, Rui.StackA)
}
func pb(Rui *structs.Rui) {
	Rui.StackA, Rui.StackB = push(Rui.StackA, Rui.StackB)
}
func sa(Rui *structs.Rui) {
	if len(Rui.StackA) < 2 {
		fmt.Println("Error : length under two! (sa)")
		return
	}
	Rui.StackA[0], Rui.StackA[1] = Rui.StackA[1], Rui.StackA[0]
}
func sb(Rui *structs.Rui) {
	if len(Rui.StackB) < 2 {
		fmt.Println("Error : length under two! (sb)")
		return
	}
	Rui.StackB[0], Rui.StackB[1] = Rui.StackB[1], Rui.StackB[0]
}
func ss(Rui *structs.Rui) {
	sa(Rui)
	sb(Rui)
}
func ra(Rui *structs.Rui) {
	if len(Rui.StackA) < 2 {
		fmt.Println("Error : length under two! (ra)")
		return
	}
	for i := 0; i < len(Rui.StackA)-1; i++ {
		Rui.StackA[i], Rui.StackA[i+1] = Rui.StackA[i+1], Rui.StackA[i]
	}
}
func rb(Rui *structs.Rui) {
	if len(Rui.StackB) < 2 {
		fmt.Println("Error : length under two! (rb)")
		return
	}
	for i := 0; i < len(Rui.StackB)-1; i++ {
		Rui.StackB[i], Rui.StackB[i+1] = Rui.StackB[i+1], Rui.StackB[i]
	}
}
func rr(Rui *structs.Rui) {
	ra(Rui)
	rb(Rui)
}
func rra(Rui *structs.Rui) {
	if len(Rui.StackA) < 2 {
		fmt.Println("Error : length under two! (rra)")
		return
	}
	for i := len(Rui.StackA) - 1; i > 0; i-- {
		Rui.StackA[i], Rui.StackA[i-1] = Rui.StackA[i-1], Rui.StackA[i]
	}
}
func rrb(Rui *structs.Rui) {
	if len(Rui.StackB) < 2 {
		fmt.Println("Error : length under two! (rrb)")
		return
	}
	for i := len(Rui.StackB) - 1; i > 0; i-- {
		Rui.StackB[i], Rui.StackB[i-1] = Rui.StackB[i-1], Rui.StackB[i]
	}
}
func rrr(Rui *structs.Rui) {
	rra(Rui)
	rrb(Rui)
}
