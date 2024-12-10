package day10

import (
	"testing"

	"github.com/revzik/aoc_2024/common/parsers"
	"github.com/revzik/aoc_2024/common/types"
	"github.com/stretchr/testify/assert"
)

func TestCountTrailsFor(t *testing.T) {
	stringMap := []string{
		"89010123",
		"78121874",
		"87430965",
		"96549874",
		"45678903",
		"32019012",
		"01329801",
		"10456732",
	}
	testMap := Map{Plane: parsers.RuneToIntMatrix(types.CreateMatrix(stringMap))}

	trails := countTrails(testMap, false)
	rating := countTrails(testMap, true)

	assert.Equal(t, 36, trails)
	assert.Equal(t, 81, rating)
}
