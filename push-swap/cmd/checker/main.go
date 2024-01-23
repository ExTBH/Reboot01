package main

import (
	"fmt"
	"os"
	"push-swap/internal/execute"
	"push-swap/internal/structs"
	"slices"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(2)
	}
	asm_3adel := os.Args[1]

	arr := strings.Fields(asm_3adel)

	arrint := []int(nil)
	for i := 0; i < len(arr); i++ {
		num, err := strconv.Atoi(arr[i])
		if err != nil {
			fmt.Println("Error: cannot convert string to intger")
			os.Exit(2)
		}
		if slices.Contains(arrint, num) {
			fmt.Println("Error: no  duplicates !!!")
			os.Exit(2)
		}
		arrint = append(arrint, num)
	}
	rui := &structs.Rui{
		Instructions: readInstructions(),
		StackA:       arrint,
		StackB:       []int{},
	}
	execute.Start(rui)
	if slices.IsSorted(rui.StackA) {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
	// fmt.Println(rui.StackA)
}

func readInstructions() []structs.Instruction {
	// check stdin have actual data in it
	fi, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	size := fi.Size()
	if size <= 0 {
		return nil
	}
	// read stdin
	var v string
	arr := []structs.Instruction(nil)
	for {
		_, err := fmt.Scanln(&v)
		if err != nil {
			break
		}
		switch structs.Instruction(v) {
		case structs.InstructionPA, structs.InstructionPB, structs.InstructionRA, structs.InstructionRB, structs.InstructionRR, structs.InstructionRRA, structs.InstructionRRB, structs.InstructionRRR, structs.InstructionSA, structs.InstructionSB, structs.InstructionSS:
			arr = append(arr, structs.Instruction(v))
		default:
			fmt.Println("Error: instruction not vaild")
			os.Exit(2)
		}
	}
	return arr
}
