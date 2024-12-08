package day5

import (
	"fmt"
	"strings"

	"github.com/revzik/aoc_2024/common/files"
	"github.com/revzik/aoc_2024/common/parsers"
)

func RunTask() {
	path := "day5/input"
	lines := files.ReadLines(path)

	ruleMap, pagesSlice := parseInput(lines)
	totalCorrect, totalIncorrect := sumPages(pagesSlice, ruleMap)
	fmt.Printf("Correct order sum: %d\n", totalCorrect)
	fmt.Printf("Correct order sum: %d\n", totalIncorrect)
}

// processing input
func parseInput(lines []string) (map[int][]int, [][]int) {
	orderRules := make([][]int, 0)
	pagesSlice := make([][]int, 0)
	finishedOrderRules := false
	for _, line := range lines {
		if line == "" {
			finishedOrderRules = true
			continue
		}

		if finishedOrderRules {
			pagesSlice = append(pagesSlice, parsePages(line))
		} else {
			orderRules = append(orderRules, parseOrderRule(line))
		}
	}

	return getRuleMap(orderRules), pagesSlice
}

func parseOrderRule(line string) []int {
	order := strings.Split(line, "|")
	return []int{parsers.ParseInt(order[0]), parsers.ParseInt(order[1])}
}

func parsePages(line string) []int {
	pagesString := strings.Split(line, ",")
	pages := make([]int, len(pagesString))
	for i, page := range pagesString {
		pages[i] = parsers.ParseInt(page)
	}
	return pages
}

func getRuleMap(orderRules [][]int) map[int][]int {
	ruleMap := make(map[int][]int)
	for _, rule := range orderRules {
		if val, ok := ruleMap[rule[0]]; ok {
			val = append(val, rule[1])
			ruleMap[rule[0]] = val
		} else {
			ruleMap[rule[0]] = []int{rule[1]}
		}
	}
	return ruleMap
}

// Processing for part one
func sumPages(pagesSlice [][]int, ruleMap map[int][]int) (int, int) {
	totalCorrect := 0
	totalIncorrect := 0

	for _, pages := range pagesSlice {
		if isCorrectlyOrdered(pages, ruleMap) {
			totalCorrect += getMiddlePage(pages)
		} else {
			pages = sortPages(pages, ruleMap)
			totalIncorrect += getMiddlePage(pages)
		}
	}

	return totalCorrect, totalIncorrect
}

func isCorrectlyOrdered(pages []int, ruleMap map[int][]int) bool {
	// for each number in the slice, check if none of the numbers before it should be after it
	for i, page := range pages {
		for j := 0; j < i; j++ {
			if !isInCorrectPlace(pages[j], ruleMap[page]) {
				return false
			}
		}
	}

	return true
}

func isInCorrectPlace(page int, incorrectPages []int) bool {
	for _, incorrectPage := range incorrectPages {
		if page == incorrectPage {
			return false
		}
	}
	return true
}

func sortPages(pages []int, ruleMap map[int][]int) []int {
	for i := 0; i < len(pages); i++ {
		for j := 0; j < i; j++ {
			if !isInCorrectPlace(pages[j], ruleMap[pages[i]]) {
				tmp := pages[i]
				pages[i] = pages[j]
				pages[j] = tmp
			}
		}
	}

	return pages
}

func getMiddlePage(pages []int) int {
	return pages[len(pages)/2]
}
