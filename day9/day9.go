package day9

import (
	"fmt"

	"github.com/revzik/aoc_2024/common/files"
	"github.com/revzik/aoc_2024/common/parsers"
)

func RunTask() {
	path := "day9/input"
	line := files.ReadLines(path)[0]

	sectors := markSectors(line)
	rearrangeSectors(sectors)
	checksum := calculateChecksum(sectors)

	fmt.Printf("New checksum: %d\n", checksum)
}

type Sector struct {
	Id    int
	Start int
	End   int
}

func (s Sector) Length() int {
	return s.End - s.Start
}

func (s Sector) IsOccupied() bool {
	return s.Id >= 0
}

func markSectors(diskMap string) []Sector {
	sectors := make([]Sector, 0)
	id := 0
	sectorStart := 0
	sectorEnd := 0

	// will not work with bigger unicode characters
	for i, char := range diskMap {
		sectorLength := parsers.RuneToInt(char)
		sectorStart = sectorEnd
		sectorEnd = sectorStart + sectorLength

		if i%2 == 0 {
			sectors = append(sectors, Sector{Id: id, Start: sectorStart, End: sectorEnd})
			id++
		} else {
			sectors = append(sectors, Sector{Id: -1, Start: sectorStart, End: sectorEnd})
		}
	}

	return sectors
}

func rearrangeSectors(sectors []Sector) []Sector {
	newSectors := make([]Sector, 0)

	sourceSectorIndex, sourceSector := getLastOccupiedSector(sectors, len(sectors))
	nextSectorIndex, nextSector := getNextSector(sectors, -1)

	for withinBounds(sourceSectorIndex, nextSectorIndex) {
		if nextSector.IsOccupied() {
			newSectors = append(newSectors, nextSector)
			nextSectorIndex, nextSector = getNextSector(sectors, nextSectorIndex)

		} else {
			if nextSector.Length() == sourceSector.Length() {

				newSector := Sector{Id: sourceSector.Id, Start: nextSector.Start, End: nextSector.End}
				newSectors = append(newSectors, newSector)

				sourceSectorIndex, sourceSector = getLastOccupiedSector(sectors, sourceSectorIndex)
				nextSectorIndex, nextSector = getNextSector(sectors, nextSectorIndex)

			} else if nextSector.Length() > sourceSector.Length() {

				newSectorEnd := nextSector.Start + sourceSector.Length()
				newSector := Sector{Id: sourceSector.Id, Start: nextSector.Start, End: newSectorEnd}
				newSectors = append(newSectors, newSector)

				nextSector.Start += sourceSector.Length()
				sourceSectorIndex, sourceSector = getLastOccupiedSector(sectors, sourceSectorIndex)

			} else {

				newSector := Sector{Id: sourceSector.Id, Start: nextSector.Start, End: nextSector.End}
				newSectors = append(newSectors, newSector)

				sourceSector.End -= nextSector.Length()
				nextSectorIndex, nextSector = getNextSector(sectors, nextSectorIndex)
			}
		}
	}

	if sourceSector.Length() > 0 {
		newSectors = append(newSectors, sourceSector)
	}

	return newSectors
}

func withinBounds(sourceSectorIndex, nextSectorIndex int) bool {
	return nextSectorIndex < sourceSectorIndex && nextSectorIndex > -1 && sourceSectorIndex > -1
}

func getLastOccupiedSector(sectors []Sector, searchIndex int) (int, Sector) {
	for i := searchIndex - 1; i >= 0; i-- {
		sector := sectors[i]
		if sector.Id >= 0 {
			return i, sector
		}
	}

	return -1, Sector{}
}

func getNextSector(sectors []Sector, searchIndex int) (int, Sector) {
	searchIndex++
	if searchIndex < len(sectors) {
		return searchIndex, sectors[searchIndex]
	}
	return -1, Sector{}
}

func calculateChecksum(sectors []Sector) int {
	total := 0
	position := 0

	for _, sector := range sectors {
		for i := sector.Start; i < sector.End; i++ {
			product := position * sector.Id
			total += product
			position++
		}
	}

	return total
}
