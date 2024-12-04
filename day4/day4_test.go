package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinesToMatrix(t *testing.T) {
	stringLines := []string{
		"abc",
		"def",
		"ghi",
		"jkl",
	}

	result := linesToMatrix(stringLines)

	assert.Equal(t, 'a', result[0][0])
	assert.Equal(t, 'e', result[1][1])
	assert.Equal(t, 'l', result[3][2])
}

func TestSearchXMAS(t *testing.T) {
	matrix := [][]rune{
		[]rune(".X.X."),
		[]rune("SMMS."),
		[]rune(".A.A."),
		[]rune("SSMMS"),
		[]rune("SAMX."),
	}

	assert.Equal(t, 2, searchXMAS(matrix, 'X', 1, 0, 0))
	assert.Equal(t, 1, searchXMAS(matrix, 'X', 3, 0, 0))
	assert.Equal(t, 3, searchXMAS(matrix, 'X', 3, 4, 0))
}

func TestSearchXShapedMAS(t *testing.T) {
	matrix := [][]rune{
		[]rune(".MMM..."),
		[]rune(".AAMSSM"),
		[]rune("SSSSAA."),
		[]rune(".AAMSSM"),
		[]rune("M.MMAS."),
	}

	assert.True(t, searchXShapedMAS(matrix, 2, 1))
	assert.True(t, searchXShapedMAS(matrix, 1, 3))
	assert.True(t, searchXShapedMAS(matrix, 4, 2))
	assert.True(t, searchXShapedMAS(matrix, 5, 2))

	assert.False(t, searchXShapedMAS(matrix, 1, 1))
	assert.False(t, searchXShapedMAS(matrix, 2, 3))
	assert.False(t, searchXShapedMAS(matrix, 3, 3))
}
