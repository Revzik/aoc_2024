package parsers

import (
	"log"
	"strconv"
)

func ParseInt(numString string) int {
	num, err := strconv.Atoi(numString)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
