package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPlot(t *testing.T) {
	field := createField([]string{
		"AAABC",
		"DAAAC",
		"DDACC",
		"DEEEE",
	})

	expectedPlotA := map[Region]map[rune]bool{
		{X: 0, Y: 0}: {'R': false, 'D': true, 'L': true, 'U': true},
		{X: 1, Y: 0}: {'R': false, 'D': false, 'L': false, 'U': true},
		{X: 2, Y: 0}: {'R': true, 'D': false, 'L': false, 'U': true},
		{X: 1, Y: 1}: {'R': false, 'D': true, 'L': true, 'U': false},
		{X: 2, Y: 1}: {'R': false, 'D': false, 'L': false, 'U': false},
		{X: 3, Y: 1}: {'R': true, 'D': true, 'L': false, 'U': true},
		{X: 2, Y: 2}: {'R': true, 'D': true, 'L': true, 'U': false},
	}
	plotA := findPlot(Region{X: 3, Y: 1}, field)
	expectedPlotB := map[Region]map[rune]bool{
		{X: 3, Y: 0}: {'R': true, 'D': true, 'L': true, 'U': true},
	}
	plotB := findPlot(Region{X: 3, Y: 0}, field)

	assert.Equal(t, expectedPlotA, plotA)
	assert.Equal(t, expectedPlotB, plotB)
}

func TestCalculatePrice(t *testing.T) {
	field := createField([]string{
		"AAACC",
		"BAAAC",
		"DDACC",
		"DEEEE",
	})

	regionA := Region{X: 3, Y: 1}
	regionB := Region{X: 0, Y: 1}
	regionC := Region{X: 3, Y: 2}

	regularA, discountedA := calculatePlotPrices(field, regionA)
	regularB, discountedB := calculatePlotPrices(field, regionB)
	regularC, discountedC := calculatePlotPrices(field, regionC)

	assert.Equal(t, 14*7, regularA)
	assert.Equal(t, 4, regularB)
	assert.Equal(t, 12*5, regularC)
	assert.Equal(t, 12*7, discountedA)
	assert.Equal(t, 4, discountedB)
	assert.Equal(t, 8*5, discountedC)
}

func TestGroupRegions(t *testing.T) {
	// shape is:
	// CC
	//  C
	// CC
	regions1 := []Region{
		{X: 4, Y: 0},
		{X: 4, Y: 1},
		{X: 4, Y: 2},
	}
	expectedGrouped1 := map[int][]int{
		4: {0, 1, 2},
	}

	regions2 := []Region{
		{X: 1, Y: 0},
		{X: 0, Y: 0},
		{X: 0, Y: 2},
	}
	expectedGrouped2 := map[int][]int{
		0: {0, 1},
		2: {0},
	}

	grouped1 := getWallsByPosition(regions1, 'R')
	grouped2 := getWallsByPosition(regions2, 'U')

	assert.Equal(t, expectedGrouped1, grouped1)
	assert.Equal(t, expectedGrouped2, grouped2)
}
