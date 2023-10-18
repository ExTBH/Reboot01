package structs

type Instruction string

type Rui struct {
	Instructions   []*Instruction
	StackA, StackB []int
}

const (
	// InstructionPA pushes the top element from stack B to stack A
	InstructionPA Instruction = "pa"

	// InstructionPB pushes the top element from stack A to stack B
	InstructionPB Instruction = "pb"

	// InstructionSA swaps the first two elements of stack A
	InstructionSA Instruction = "sa"

	// InstructionSB swaps the first two elements of stack B
	InstructionSB Instruction = "sb"

	// InstructionSS executes both InstructionSA and InstructionSB
	InstructionSS Instruction = "ss"

	// InstructionRA rotates stack A (shifts all elements up by 1, the first element becomes the last)
	InstructionRA Instruction = "ra"

	// InstructionRB rotates stack B
	InstructionRB Instruction = "rb"

	// InstructionRR executes both InstructionRA and InstructionRB
	InstructionRR Instruction = "rr"

	// InstructionRRA reverse rotates stack A (shifts all elements down by 1, the last element becomes the first)
	InstructionRRA Instruction = "rra"

	// InstructionRRB reverse rotates stack B
	InstructionRRB Instruction = "rrb"

	// InstructionRRR executes both InstructionRRA and InstructionRRB
	InstructionRRR Instruction = "rrr"
)
