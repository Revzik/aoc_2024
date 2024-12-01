package main

import "testing"

func TestDistancesMinimalExample(t *testing.T) {
	idListsLines := []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}
	expectedTotal := 11

	total := TotalDistance(idListsLines)

	if total != expectedTotal {
		t.Errorf("Wrong distance, expected %d, got %d", expectedTotal, total)
	}
}

func TestSimilarityScoreMinimalExample(t *testing.T) {
	idListsLines := []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}
	expectedTotal := 31

	total := SimilarityScore(idListsLines)

	if total != expectedTotal {
		t.Errorf("Wrong score, expected %d, got %d", expectedTotal, total)
	}
}
