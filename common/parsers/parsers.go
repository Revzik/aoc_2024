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
