package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input := []string{
		"Register A: 729",
		"Register B: 0",
		"Register C: 0",
		"",
		"Program: 0,1,5,4,3,0",
	}
	reg, inst := parseInput(input)

	output := execute(inst, reg)

	assert.Equal(t, []int{4, 6, 3, 5, 6, 3, 5, 2, 1, 0}, output)
}
