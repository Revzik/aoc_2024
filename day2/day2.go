package day2

import (
	"fmt"
	"strings"

	"github.com/revzik/aoc_2024/common/files"
	"github.com/revzik/aoc_2024/common/parsers"
)

func RunTask() {
	lines := files.ReadLines("day_2/input.txt")

	fmt.Printf("Safe reports: %d\n", countSafeReports(lines, false))
	// Ashamed of this but I'll fix it
	fmt.Printf("Safe reports with dampener (tried to be smart): %d\n", countSafeReports(lines, true))
	fmt.Printf("Safe reports with dampener (brute force): %d\n", countDampenedReportsBruteForce(lines))
}

func countSafeReports(reportList []string, useDampener bool) int {
	safeCount := 0
	for _, report := range reportList {
		parsedReport := parseReport(report)
		if useDampener && isReportSafeWithDampener(parsedReport) {
			safeCount++
		} else if isReportSafe(parsedReport) {
			safeCount++
		}
	}
	return safeCount
}

func countDampenedReportsBruteForce(reportList []string) int {
	safeCount := 0
	for _, report := range reportList {
		parsedReport := parseReport(report)
		if isReportSafe(parsedReport) || bruteForceDampener(parsedReport) {
			safeCount++
		}
	}
	return safeCount
}

func bruteForceDampener(report []int) bool {
	for i := 0; i < len(report); i++ {
		if checkReportWithDampener(report, i) {
			return true
		}
	}
	return false
}

func parseReport(report string) []int {
	splitReport := strings.Split(report, " ")

	parsedReport := make([]int, len(splitReport))
	for i, val := range splitReport {
		parsedReport[i] = parsers.ParseInt(val)
	}

	return parsedReport
}

func isReportSafeWithDampener(report []int) bool {
	// TODO: Think of a better way and fix it...
	//        - one edge case is missing
	//        - analyze only fragments with offenders

	// edge case for offender at the beginning
	if checkReportWithDampener(report, 0) {
		return true
	}

	reportCorrect := true
	previousDifference := 0

	for i := 1; i < len(report); i++ {
		difference := report[i] - report[i-1]

		if difference == 0 || difference > 3 || difference < -3 || difference*previousDifference < 0 {
			if checkReportWithDampener(report, i) {
				return true
			}
			reportCorrect = false
		}

		previousDifference = difference
	}

	return reportCorrect
}

func checkReportWithDampener(report []int, index int) bool {
	amendedReport := make([]int, len(report))
	copy(amendedReport, report)
	amendedReport = append(amendedReport[:index], amendedReport[index+1:]...)
	return isReportSafe(amendedReport)
}

func isReportSafe(report []int) bool {
	previousDifference := 0

	for i := 1; i < len(report); i++ {
		difference := report[i] - report[i-1]

		if difference == 0 || difference > 3 || difference < -3 || difference*previousDifference < 0 {
			return false
		}

		previousDifference = difference
	}

	return true
}
