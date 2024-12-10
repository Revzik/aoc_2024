package parsers

import (
	"log"
	"strconv"
)

func StringToInt(numString string) int {
	num, err := strconv.Atoi(numString)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func RuneToInt(numRune rune) int {
	return int(numRune - '0')
}

func RuneToIntMatrix(matrix [][]rune) [][]int {
	intMatrix := make([][]int, len(matrix))
	for i, row := range matrix {
		intMatrix[i] = make([]int, len(row))
		for j, char := range row {
			intMatrix[i][j] = RuneToInt(char)
		}
	}
	return intMatrix
}
