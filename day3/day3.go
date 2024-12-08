package day3

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/revzik/aoc_2024/common/files"
	"github.com/revzik/aoc_2024/common/parsers"
)

const DoEnabler = "do()"
const DontEnabler = "don't()"

func RunTask() {
	lines := files.ReadLines("day3/input")

	fmt.Printf("All multiplications: %d\n", computeAllMultiplications(lines))
	fmt.Printf("All multiplications with enablers: %d\n", computeAllMultiplicationsWithEnablers(lines))
}

func computeAllMultiplications(instructions []string) int {
	total := 0
	for _, instruction := range instructions {
		total += computeMultiplications(instruction)
	}
	return total
}

func computeMultiplications(instruction string) int {
	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	matches := re.FindAllString(instruction, -1)

	total := 0
	for _, match := range matches {
		a, b := extractNumbers(match)
		total += a * b
	}

	return total
}

func computeAllMultiplicationsWithEnablers(instructions []string) int {
	enabler := DoEnabler
	var result int
	total := 0

	for _, instruction := range instructions {
		result, enabler = computeWithEnablers(instruction, enabler)
		total += result
	}

	return total
}

func computeWithEnablers(instruction string, enabler string) (int, string) {
	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`)
	matches := re.FindAllString(instruction, -1)

	total := 0
	for _, match := range matches {
		if match == DoEnabler || match == DontEnabler {
			enabler = match
		} else if enabler == DoEnabler {
			a, b := extractNumbers(match)
			total += a * b
		}
	}

	return total, enabler
}

func extractNumbers(multiplication string) (int, int) {
	strippedMatch := multiplication[4 : len(multiplication)-1]
	numbers := strings.Split(strippedMatch, ",")
	return parsers.ParseInt(numbers[0]), parsers.ParseInt(numbers[1])
}
