package day19

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPossible(t *testing.T) {
	towels := []string{"r", "rw", "ubrr", "b", "w", "gru"}

	assert.True(t, isPossible("rrwbgruww", towels))
	assert.True(t, isPossible("rrrrrrrrrrr", towels))
	assert.True(t, isPossible("ubrrubrrubrr", towels))
	assert.True(t, isPossible("rrwubrrbwgru", towels))

	assert.False(t, isPossible("rrwgrb", towels))
	assert.False(t, isPossible("uubrr", towels))
}

func TestLongPattern(t *testing.T) {
	towels := []string{
		"wrg",
		"rgrrg",
		"uurb",
		"urgg",
		"rrrrrwwww",
		"bgwu",
		"bwuww",
		"rg",
		"ggguguuu",
		"ru",
		"gbbgw",
		"rwbggbwuw",
		"wrwwg",
		"ggbg",
		"wuburg",
		"guuw",
		"ubububu",
		"grwbuu",
		"uubgrr",
		"bbgr",
		"rwwggbg",
		"gg",
		"uwgbwg",
		"rw",
		"grrg",
		"uubw",
		"uuu",
		"ggbg",
		"wrww",
		"wgg",
		"rggb",
		"ubu",
		"wuwww",
		"ubgb",
		"urg",
		"rwbg",
		"guruu",
		"bgwuw",
		"rrgb",
	}
	pattern := "wrgrrgbbgwuwguruurwbgurgubgbwuwwwuburggbwggwrwwggbg"

	assert.True(t, isPossible(pattern, towels))
}

func TestReduceTowels(t *testing.T) {
	towels := []string{
		"rrrrrrrr",
		"rr",
		"bwurg",
		"rg",
		"b",
		"wur",
		"g",
		"rwur",
		"rrbgg",
		"wrr",
	}
	expectedReduced := []string{
		"b",
		"g",
		"rr",
		"rg",
		"wur",
		"wrr",
		"rwur",
	}

	reduced := reduceTowels(towels)

	assert.Equal(t, expectedReduced, reduced)
}
