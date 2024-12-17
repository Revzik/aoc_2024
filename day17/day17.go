package day17

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	f "github.com/revzik/aoc_2024/common/files"
	p "github.com/revzik/aoc_2024/common/parsers"
)

func RunTask() {
	path := "day17/input"
	lines := f.ReadLines(path)

	registers, program := parseInput(lines)
	output := execute(program, registers)

	fmt.Printf("Output: %v\n", formatOutput(output))
	fmt.Printf("Minimal A value: %d\n", findRegisterValue(program))
}

type Registers struct {
	A int
	B int
	C int
}

func parseInput(lines []string) (*Registers, []int) {
	reg := Registers{
		A: p.StringToInt(strings.TrimPrefix(lines[0], "Register A: ")),
		B: p.StringToInt(strings.TrimPrefix(lines[1], "Register B: ")),
		C: p.StringToInt(strings.TrimPrefix(lines[2], "Register C: ")),
	}

	splitInstructions := strings.Split(strings.TrimPrefix(lines[4], "Program: "), ",")
	program := make([]int, len(splitInstructions))
	for i, instruction := range splitInstructions {
		program[i] = p.StringToInt(instruction)
	}

	return &reg, program
}

// operations
func combo(operand int, reg *Registers) int {
	if operand < 4 {
		return operand
	}
	if operand == 4 {
		return reg.A
	}
	if operand == 5 {
		return reg.B
	}
	if operand == 6 {
		return reg.C
	}
	panic("program invalid")
}

func adv(operand int, reg *Registers) {
	reg.A = divide(operand, reg)
}

func bxl(operand int, reg *Registers) {
	reg.B ^= operand
}

func bst(operand int, reg *Registers) {
	reg.B = combo(operand, reg) % 8
}

func jnz(operand int, reg *Registers, ptr *int) {
	if reg.A != 0 {
		*ptr = operand - 2
	}
}

func bxc(_ int, reg *Registers) {
	reg.B ^= reg.C
}

func out(operand int, reg *Registers) int {
	return combo(operand, reg) % 8
}

func bdv(operand int, reg *Registers) {
	reg.B = divide(operand, reg)
}

func cdv(operand int, reg *Registers) {
	reg.C = divide(operand, reg)
}

func divide(operand int, reg *Registers) int {
	return reg.A / int(math.Pow(2, float64(combo(operand, reg))))
}

// part one
func execute(program []int, reg *Registers) []int {
	ptr := 0
	output := make([]int, 0)

	for {
		if ptr >= len(program) {
			return output
		}

		opcode := program[ptr]
		operand := program[ptr+1]

		switch opcode {
		case 0:
			adv(operand, reg)
		case 1:
			bxl(operand, reg)
		case 2:
			bst(operand, reg)
		case 3:
			jnz(operand, reg, &ptr)
		case 4:
			bxc(operand, reg)
		case 5:
			output = append(output, out(operand, reg))
		case 6:
			bdv(operand, reg)
		case 7:
			cdv(operand, reg)
		}

		ptr += 2
	}
}

func formatOutput(input []int) string {
	stringInput := make([]string, len(input))

	for i, item := range input {
		stringInput[i] = strconv.Itoa(item)
	}

	return strings.Join(stringInput, ",")
}

// part two
// print out which register values print out the matching beginning of the program,
// hopefully find a relation between the two
func findRegisterValue(program []int) int {
	targetMatches := len(program)
	matches := 0
	biggestMatches := 0
	counter := 0

	for targetMatches != matches {
		registers := Registers{A: counter, B: 0, C: 0}

		result := execute(program, &registers)
		matches = comparePrograms(program, result)

		if matches >= biggestMatches {
			biggestMatches = matches
			fmt.Printf("A: %d, expected: %v, actual: %v\n", counter, program, result)
		}

		counter++
	}

	return counter
}

func comparePrograms(expected []int, actual []int) int {

	iterations := min(len(expected), len(actual))
	i := 0

	for i < iterations {
		if expected[i] != actual[i] {
			break
		}
		i++
	}

	return i
}
