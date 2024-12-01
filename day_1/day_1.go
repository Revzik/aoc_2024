package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

	fmt.Printf("Part 1: %d\n", TotalDistance(lines))
	fmt.Printf("Part 2: %d\n", SimilarityScore(lines))
}

func TotalDistance(idListsLines []string) int {
	list1, list2 := linesToLists(idListsLines)

	sort.Ints(list1)
	sort.Ints(list2)

	totalDistance := 0
	for i := 0; i < len(idListsLines); i++ {
		distance := list1[i] - list2[i]
		if distance < 0 {
			distance = -distance
		}
		totalDistance += distance
	}

	return totalDistance
}

func SimilarityScore(idListsLines []string) int {
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
		list1[i] = parseInt(splitLine[0])
		list2[i] = parseInt(splitLine[1])
	}
	return list1, list2
}

func parseInt(numString string) int {
	num, err := strconv.Atoi(numString)
	if err != nil {
		log.Fatal(err)
	}
	return num
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
