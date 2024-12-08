package day7

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/revzik/aoc_2024/common/files"
	"github.com/revzik/aoc_2024/common/parsers"
)

func RunTask() {
	path := "day7/input"
	lines := files.ReadLines(path)

	fmt.Printf("Possible equations: %d\n", countPossibleEquations(lines, false))
	fmt.Printf("Possible equations with concatenation: %d\n", countPossibleEquations(lines, true))
}

func countPossibleEquations(lines []string, concatenate bool) int {
	total := 0

	for _, line := range lines {
		target, ingredients := parseLine(line)
		if canBeCalculated(target, ingredients, concatenate) {
			total += target
		}
	}

	return total
}

// Input processing
func parseLine(line string) (int, []string) {
	splitLine := strings.Split(line, ":")

	target := parsers.ParseInt(splitLine[0])
	ingredients := strings.Split(strings.Trim(splitLine[1], " "), " ")

	return target, ingredients
}

// Operations processing
func createOperations(ingredients []string) []string {
	operations := make([]string, len(ingredients)-1)
	for i := range operations {
		operations[i] = "+"
	}
	return operations
}

func nextOperationCombination(operations []string, concatenate bool) {
	stopProcessing := false

	for i, operation := range operations {
		if stopProcessing {
			break
		}

		if concatenate {
			if operation == "+" {
				operations[i] = "*"
				stopProcessing = true
			} else if operation == "*" {
				operations[i] = "||"
				stopProcessing = true
			} else if operation == "||" {
				operations[i] = "+"
			}
		} else {
			if operation == "+" {
				operations[i] = "*"
				stopProcessing = true
			} else if operation == "*" {
				operations[i] = "+"
			}
		}
	}
}

func isLastOperationCombination(operations []string, concatenate bool) bool {
	for _, operation := range operations {
		if operation == "+" || (concatenate && operation == "*") {
			return false
		}
	}
	return true
}

// Calculation
func canBeCalculated(target int, ingredients []string, concatenate bool) bool {
	operations := createOperations(ingredients)
	if target == calculate(ingredients, operations) {
		return true
	}

	for !isLastOperationCombination(operations, concatenate) {
		nextOperationCombination(operations, concatenate)
		if target == calculate(ingredients, operations) {
			return true
		}
	}

	return false
}

func calculate(ingredients []string, operations []string) int {
	total := parsers.ParseInt(ingredients[0])

	for i, operation := range operations {
		if operation == "+" {
			total += parsers.ParseInt(ingredients[i+1])
		} else if operation == "*" {
			total *= parsers.ParseInt(ingredients[i+1])
		} else if operation == "||" {
			totalStr := strconv.Itoa(total)
			totalStr += ingredients[i+1]
			total = parsers.ParseInt(totalStr)
		}
	}

	return total
}
