package main

import "testing"

func TestCountSafeReports(t *testing.T) {
	reportList := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}

	safeCount := CountSafeReports(reportList, false)

	if safeCount != 2 {
		t.Fatalf("Expected 2 reports to be safe, got %d", safeCount)
	}
}

func TestCountSafeReportsWithDampener(t *testing.T) {
	reportList := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}

	safeCount := CountSafeReports(reportList, true)

	if safeCount != 4 {
		t.Fatalf("Expected 4 reports to be safe, got %d", safeCount)
	}
}

func TestIsReportSafeAscending(t *testing.T) {
	report := [5]int{1, 2, 4, 7, 8}

	if !isReportSafe(report[:]) {
		t.Fatalf("Expected report %v to be safe", report)
	}
}

func TestIsReportSafeDescending(t *testing.T) {
	report := [5]int{11, 10, 7, 5, 3}

	if !isReportSafe(report[:]) {
		t.Fatalf("Expected report %v to be safe", report)
	}
}

func TestIsReportUnsafeAscending(t *testing.T) {
	report := [5]int{1, 2, 6, 7, 8}

	if isReportSafe(report[:]) {
		t.Fatalf("Expected report %v to be unsafe", report)
	}
}

func TestIsReportUnsafeDescending(t *testing.T) {
	report := [5]int{11, 10, 6, 5, 3}

	if isReportSafe(report[:]) {
		t.Fatalf("Expected report %v to be unsafe", report)
	}
}

func TestIsReportUnsafeSameValues(t *testing.T) {
	report := [5]int{11, 10, 7, 7, 5}

	if isReportSafe(report[:]) {
		t.Fatalf("Expected report %v to be unsafe", report)
	}
}

func TestIsReportUnsafeChangeDirection(t *testing.T) {
	report := [5]int{11, 10, 7, 8, 9}

	if isReportSafe(report[:]) {
		t.Fatalf("Expected report %v to be unsafe", report)
	}
}

func TestIsDampenedReportSafeOneLevel(t *testing.T) {
	report := [5]int{11, 10, 7, 8, 4}

	if !isReportSafeWithDampener(report[:]) {
		t.Fatalf("Expected report %v to be safe with dampener", report)
	}
}

func TestIsDampenedReportSafeOneLevelBeginning(t *testing.T) {
	report := [5]int{9, 10, 7, 6, 4}

	if !isReportSafeWithDampener(report[:]) {
		t.Fatalf("Expected report %v to be safe with dampener", report)
	}
}

func TestIsDampenedReportSafeOneLevelEnd(t *testing.T) {
	report := [5]int{11, 10, 7, 6, 1}

	if !isReportSafeWithDampener(report[:]) {
		t.Fatalf("Expected report %v to be safe with dampener", report)
	}
}

func TestIsDampenedReportUnsafeTwoLevels(t *testing.T) {
	report := [6]int{5, 1, 2, 3, 4, 1}

	if isReportSafeWithDampener(report[:]) {
		t.Fatalf("Expected report %v to be unsafe with dampener", report)
	}
}
