package day9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSectorLenght(t *testing.T) {
	sector := Sector{Id: 0, Start: 3, End: 6}

	assert.Equal(t, 3, sector.Length())
}
func TestMarkOccupiedSpace(t *testing.T) {
	diskMap := "12345"
	// 0..111....22222
	expectedOutput := []Sector{
		{Id: 0, Start: 0, End: 1},
		{Id: -1, Start: 1, End: 3},
		{Id: 1, Start: 3, End: 6},
		{Id: -1, Start: 6, End: 10},
		{Id: 2, Start: 10, End: 15},
	}

	occupiedSpace := markSectors(diskMap)

	assert.Equal(t, expectedOutput, occupiedSpace)
}

func TestRearrangeSectors(t *testing.T) {
	// 0..111....22222
	sectors := []Sector{
		{Id: 0, Start: 0, End: 1},
		{Id: -1, Start: 1, End: 3},
		{Id: 1, Start: 3, End: 6},
		{Id: -1, Start: 6, End: 10},
		{Id: 2, Start: 10, End: 15},
	}
	// 022111222
	expectedOutput := []Sector{
		{Id: 0, Start: 0, End: 1},
		{Id: 2, Start: 1, End: 3},
		{Id: 1, Start: 3, End: 6},
		{Id: 2, Start: 6, End: 9},
	}

	newDiskMap := rearrangeSectors(sectors)

	assert.Equal(t, expectedOutput, newDiskMap)
}

func TestLastSector(t *testing.T) {
	sectors := []Sector{
		{Id: 1, Start: 0, End: 1},
		{Id: -1, Start: 1, End: 3},
		{Id: 2, Start: 3, End: 6},
		{Id: -1, Start: 6, End: 10},
		{Id: 3, Start: 10, End: 15},
	}

	lastSectorIndex, lastSector := getLastOccupiedSector(sectors, len(sectors))
	assert.Equal(t, 4, lastSectorIndex)
	assert.Equal(t, Sector{Id: 3, Start: 10, End: 15}, lastSector)

	lastSectorIndex, lastSector = getLastOccupiedSector(sectors, len(sectors)-1)
	assert.Equal(t, 2, lastSectorIndex)
	assert.Equal(t, Sector{Id: 2, Start: 3, End: 6}, lastSector)

	lastSectorIndex, lastSector = getLastOccupiedSector(sectors, len(sectors)-3)
	assert.Equal(t, 0, lastSectorIndex)
	assert.Equal(t, Sector{Id: 1, Start: 0, End: 1}, lastSector)
}

func TestNextSector(t *testing.T) {
	sectors := []Sector{
		{Id: 1, Start: 0, End: 1},
		{Id: -1, Start: 1, End: 3},
		{Id: 2, Start: 3, End: 6},
		{Id: -1, Start: 6, End: 10},
		{Id: 3, Start: 10, End: 15},
	}

	sectorIndex, sector := getNextSector(sectors, -1)
	assert.Equal(t, 0, sectorIndex)
	assert.Equal(t, Sector{Id: 1, Start: 0, End: 1}, sector)

	sectorIndex, sector = getNextSector(sectors, 0)
	assert.Equal(t, 1, sectorIndex)
	assert.Equal(t, Sector{Id: -1, Start: 1, End: 3}, sector)

	sectorIndex, sector = getNextSector(sectors, 3)
	assert.Equal(t, 4, sectorIndex)
	assert.Equal(t, Sector{Id: 3, Start: 10, End: 15}, sector)

	sectorIndex, sector = getNextSector(sectors, 4)
	assert.Equal(t, -1, sectorIndex)
	assert.Equal(t, Sector{}, sector)
}

func TestCalculateChecksum(t *testing.T) {
	sectors := []Sector{
		{Id: 0, Start: 0, End: 1},
		{Id: 2, Start: 1, End: 3},
		{Id: 1, Start: 3, End: 6},
		{Id: 2, Start: 6, End: 9},
	}

	checksum := calculateChecksum(sectors)

	assert.Equal(t, 60, checksum)
}

func TestPartOne(t *testing.T) {
	diskMap := "2333133121414131402"

	sectors := markSectors(diskMap)
	sectors = rearrangeSectors(sectors)
	checksum := calculateChecksum(sectors)

	assert.Equal(t, 1928, checksum)
}
