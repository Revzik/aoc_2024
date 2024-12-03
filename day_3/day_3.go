package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const DoEnabler = "do()"
const DontEnabler = "don't()"

func main() {
	path := "input.txt"

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error occurred while opening file %v, %v", path, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Printf("All multiplications: %d\n", ComputeAllMultiplications(lines))
	fmt.Printf("All multiplications with enablers: %d\n", ComputeAllMultiplicationsWithEnablers(lines))
}

func ComputeAllMultiplications(instructions []string) int {
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

func ComputeAllMultiplicationsWithEnablers(instructions []string) int {
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
	return parseInt(numbers[0]), parseInt(numbers[1])
}

// TODO: Extract common functions into importable package
func parseInt(numString string) int {
	num, err := strconv.Atoi(numString)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
