package day19

import (
	"fmt"
	"slices"
	"strings"

	f "github.com/revzik/aoc_2024/common/files"
)

func RunTask() {
	path := "day19/input"
	lines := f.ReadLines(path)

	towels, patterns := parseInput(lines)
	reducedTowels := reduceTowels(towels)

	amountPossible, possiblePatterns := countPossiblePatterns(patterns, reducedTowels)
	allCombinations := countCombinationsForAll(possiblePatterns, towels)

	fmt.Printf("Possible patterns: %d\n", amountPossible)
	fmt.Printf("All possible combinations: %d\n", allCombinations)
}

func parseInput(lines []string) ([]string, []string) {
	towels := strings.Split(lines[0], ", ")
	patterns := lines[2:]
	return towels, patterns
}

func reduceTowels(towels []string) []string {
	slices.SortFunc(towels, func(a, b string) int {
		return len(a) - len(b)
	})

	for i := len(towels) - 1; i >= 0; i-- {
		reduced := make([]string, len(towels))
		copy(reduced, towels)
		reduced = append(reduced[:i], reduced[i+1:]...)
		if isPossible(towels[i], reduced) {
			towels = reduced
		}
	}

	return towels
}

func countPossiblePatterns(patterns []string, towels []string) (int, []string) {
	total := 0
	possible := make([]string, 0)

	for _, pattern := range patterns {
		if isPossible(pattern, towels) {
			possible = append(possible, pattern)
			total++
		}
	}
	return total, possible
}

func isPossible(pattern string, towels []string) bool {
	for _, towel := range towels {
		if pattern == towel {
			return true
		}
		if len(pattern) >= len(towel) && pattern[:len(towel)] == towel {
			if isPossible(pattern[len(towel):], towels) {
				return true
			}
		}
	}
	return false
}

func countCombinationsForAll(patterns []string, towels []string) int {
	knownCombinations := make(map[string]int)
	total := 0

	for _, pattern := range patterns {
		total += countCombinations(pattern, towels, knownCombinations, 0)
	}

	return total
}

func countCombinations(pattern string, towels []string, knownCombinations map[string]int, amount int) int {
	if val, ok := knownCombinations[pattern]; ok {
		return val
	}

	for _, towel := range towels {
		if pattern == towel {
			amount += 1
		} else if len(pattern) >= len(towel) && pattern[:len(towel)] == towel {
			amount += countCombinations(pattern[len(towel):], towels, knownCombinations, amount)
		}
	}

	knownCombinations[pattern] = amount
	return amount
}
