package day11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/revzik/aoc_2024/common/files"
	"github.com/revzik/aoc_2024/common/parsers"
)

func RunTask() {
	path := "day11/input"
	lines := files.ReadLines(path)
	stones := parseLine(lines[0])

	fmt.Printf("After 25 blinks: %d\n", countStones(stones, 25))
	fmt.Printf("After 75 blinks: %d\n", countStones(stones, 75))
}

func parseLine(line string) []int {
	splitLine := strings.Split(line, " ")
	numbers := make([]int, len(splitLine))

	for i, stringNumber := range splitLine {
		numbers[i] = parsers.StringToInt(stringNumber)
	}

	return numbers
}

// more performant attempt
// order doesn't matter so on each step we can count repeating numbers, process them once and multiply the result
func countStones(stones []int, blink int) int {
	total := 0

	stonesMap := make(map[int]int)
	for _, stone := range stones {
		putOrAdd(stonesMap, stone, 1)
	}

	for i := 0; i < blink; i++ {
		stonesMap = processBlink(stonesMap)
	}

	for _, amount := range stonesMap {
		total += amount
	}

	return total
}

func processBlink(stones map[int]int) map[int]int {
	newStones := make(map[int]int, 0)

	for stone, amount := range stones {
		stoneString := strconv.Itoa(stone)

		if len(stoneString)%2 == 0 {
			str1 := stoneString[:len(stoneString)/2]
			str2 := stoneString[len(stoneString)/2:]
			putOrAdd(newStones, parsers.StringToInt(str1), amount)
			putOrAdd(newStones, parsers.StringToInt(str2), amount)

		} else if stone == 0 {
			putOrAdd(newStones, 1, amount)

		} else {
			putOrAdd(newStones, stone*2024, amount)
		}
	}

	return newStones
}

func putOrAdd(m map[int]int, k int, v int) {
	if val, ok := m[k]; ok {
		m[k] = val + v
	} else {
		m[k] = v
	}
}
