package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperations(t *testing.T) {
	numbers := []string{"1", "1", "1", "1"}

	operations := createOperations(numbers)
	assert.Equal(t, []string{"+", "+", "+"}, operations)

	nextOperationCombination(operations, false)
	assert.Equal(t, []string{"*", "+", "+"}, operations)

	nextOperationCombination(operations, false)
	assert.Equal(t, []string{"+", "*", "+"}, operations)

	nextOperationCombination(operations, false)
	assert.Equal(t, []string{"*", "*", "+"}, operations)

	nextOperationCombination(operations, false)
	assert.Equal(t, []string{"+", "+", "*"}, operations)

	nextOperationCombination(operations, false)
	nextOperationCombination(operations, false)
	nextOperationCombination(operations, false)
	assert.Equal(t, []string{"*", "*", "*"}, operations)
}

func TestOperationsWithConcatenation(t *testing.T) {
	numbers := []string{"1", "1", "1"}

	operations := createOperations(numbers)
	assert.Equal(t, []string{"+", "+"}, operations)

	nextOperationCombination(operations, true)
	assert.Equal(t, []string{"*", "+"}, operations)

	nextOperationCombination(operations, true)
	assert.Equal(t, []string{"||", "+"}, operations)

	nextOperationCombination(operations, true)
	assert.Equal(t, []string{"+", "*"}, operations)

	nextOperationCombination(operations, true)
	nextOperationCombination(operations, true)
	nextOperationCombination(operations, true)
	assert.Equal(t, []string{"+", "||"}, operations)
}

func TestPartOne(t *testing.T) {
	input := []string{
		"190: 10 19",
		"3267: 81 40 27",
		"83: 17 5",
		"156: 15 6",
		"7290: 6 8 6 15",
		"161011: 16 10 13",
		"192: 17 8 14",
		"21037: 9 7 18 13",
		"292: 11 6 16 20",
	}

	total := countPossibleEquations(input, false)

	assert.Equal(t, 3749, total)
}

func TestPartTwo(t *testing.T) {
	input := []string{
		"190: 10 19",
		"3267: 81 40 27",
		"83: 17 5",
		"156: 15 6",
		"7290: 6 8 6 15",
		"161011: 16 10 13",
		"192: 17 8 14",
		"21037: 9 7 18 13",
		"292: 11 6 16 20",
	}

	total := countPossibleEquations(input, true)

	assert.Equal(t, 11387, total)
}
