package day1

import (
	"fmt"
	"sort"
	"strings"

	"github.com/revzik/aoc_2024/common/files"
	"github.com/revzik/aoc_2024/common/intmath"
	"github.com/revzik/aoc_2024/common/parsers"
)

func RunTask() {
	lines := files.ReadLines("day1/input.txt")

	fmt.Printf("Part 1: %d\n", totalDistance(lines))
	fmt.Printf("Part 2: %d\n", similarityScore(lines))
}

func totalDistance(idListsLines []string) int {
	list1, list2 := linesToLists(idListsLines)

	sort.Ints(list1)
	sort.Ints(list2)

	totalDistance := 0
	for i := 0; i < len(idListsLines); i++ {
		distance := intmath.Abs(list1[i] - list2[i])
		totalDistance += distance
	}

	return totalDistance
}

func similarityScore(idListsLines []string) int {
	list1, list2 := linesToLists(idListsLines)

	computedValues := make(map[int]int)
	totalScore := 0

	for i := 0; i < len(idListsLines); i++ {
		number := list1[i]
		if val, ok := computedValues[number]; ok {
			totalScore += val
		} else {
			val = number * countOccurrences(number, list2)
			computedValues[number] = val
			totalScore += val
		}
	}

	return totalScore
}

func linesToLists(idListsLines []string) ([]int, []int) {
	list1 := make([]int, len(idListsLines))
	list2 := make([]int, len(idListsLines))

	for i, line := range idListsLines {
		splitLine := strings.Split(line, "   ")
		list1[i] = parsers.ParseInt(splitLine[0])
		list2[i] = parsers.ParseInt(splitLine[1])
	}
	return list1, list2
}

func countOccurrences(number int, list []int) int {
	occurrences := 0
	for _, listNumber := range list {
		if listNumber == number {
			occurrences++
		}
	}
	return occurrences
}
