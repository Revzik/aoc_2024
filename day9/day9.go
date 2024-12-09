package day9

import (
	"errors"
	"fmt"

	"github.com/revzik/aoc_2024/common/files"
	"github.com/revzik/aoc_2024/common/parsers"
)

func RunTask() {
	path := "day9/input"
	line := files.ReadLines(path)[0]

	freeSpace, sectors := markSectors(line)
	fragmentatedSectors := rearrangeByFragmentation(sectors, freeSpace)
	shiftedSectors := rearrangeByShifting(sectors)

	fmt.Printf("Fragmentated checksum: %v\n", calculateChecksum(fragmentatedSectors))
	fmt.Printf("Shifted checksum: %v\n", calculateChecksum(shiftedSectors))
}

func markSectors(diskMap string) (int, []int) {
	sectors := make([]int, 0)

	id := -1
	value := id

	freeSpaces := 0

	// will not work with bigger unicode characters
	for i, char := range diskMap {
		newSector := make([]int, parsers.RuneToInt(char))

		if i%2 == 0 {
			id++
			value = id
		} else {
			value = -1
			freeSpaces += len(newSector)
		}

		for j := 0; j < len(newSector); j++ {
			newSector[j] = value
		}

		sectors = append(sectors, newSector...)
	}

	return freeSpaces, sectors
}

func rearrangeByFragmentation(sectors []int, freeSpaces int) []int {
	newSectors := make([]int, len(sectors)-freeSpaces)

	tailIdx := len(sectors) - 1
	for i := range newSectors {
		if sectors[i] == -1 {
			for sectors[tailIdx] == -1 {
				tailIdx--
			}
			newSectors[i] = sectors[tailIdx]
			tailIdx--
		} else {
			newSectors[i] = sectors[i]
		}
	}

	return newSectors
}

func rearrangeByShifting(sectors []int) []int {
	newSectors := make([]int, len(sectors))
	copy(newSectors, sectors)

	blockSearchIndex := len(newSectors) - 1
	var block []int

	for blockSearchIndex > 0 {
		blockSearchIndex, block = nextBlock(sectors, blockSearchIndex)
		spaceIndex, err := nextFreeSpace(newSectors, block, blockSearchIndex)
		if err == nil {
			for i := range block {
				newSectors[i+spaceIndex] = block[0]
			}
			for i := range block {
				newSectors[i+blockSearchIndex+1] = -1
			}
		}
	}

	return newSectors
}

func nextBlock(sectors []int, searchIndex int) (int, []int) {
	for searchIndex >= 0 {
		if sectors[searchIndex] >= 0 {
			break
		}
		searchIndex--
	}

	nextIndex := searchIndex
	for nextIndex >= 0 && sectors[nextIndex] == sectors[searchIndex] {
		nextIndex--
	}

	return nextIndex, sectors[nextIndex+1 : searchIndex+1]
}

func nextFreeSpace(sectors []int, block []int, maxIndex int) (int, error) {
	boundary := len(sectors) - len(block) - 1
	if boundary < maxIndex {
		maxIndex = boundary
	}

	for i := 0; i < boundary && i < maxIndex; i++ {
		blockFound := true
		for j := range block {
			if sectors[i+j] >= 0 {
				blockFound = false
				break
			}
		}

		if blockFound {
			return i, nil
		}
	}

	return -1, errors.New("no free space")
}

func calculateChecksum(sectors []int) int {
	total := 0

	for i, val := range sectors {
		if val >= 0 {
			total += i * val
		}
	}

	return total
}
