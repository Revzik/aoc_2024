package day14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePositions(t *testing.T) {
	input := []string{
		"p=0,4 v=3,-3",
		"p=6,3 v=-1,-3",
		"p=10,3 v=-1,2",
		"p=2,0 v=2,-1",
		"p=0,0 v=1,3",
		"p=3,0 v=-2,-2",
		"p=7,6 v=-1,-3",
		"p=3,0 v=-1,-2",
		"p=9,3 v=2,3",
		"p=7,3 v=-1,2",
		"p=2,4 v=2,-3",
		"p=9,5 v=-3,-3",
	}
	hall := createHall(input, 11, 7)

	hall = calculatePositionsAfter(hall, 100)
	safetyFactor := calculateSafetyFactor(hall)

	assert.Equal(t, 12, safetyFactor)
}
