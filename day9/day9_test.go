package day9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarkSectors(t *testing.T) {
	diskMap := "12345"
	expectedSectors := []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2}

	freeSpaces, sectors := markSectors(diskMap)

	assert.Equal(t, expectedSectors, sectors)
	assert.Equal(t, 6, freeSpaces)
}

func TestRearrangeByFragmentation(t *testing.T) {
	sectors := []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2}
	expectedSectors := []int{0, 2, 2, 1, 1, 1, 2, 2, 2}

	newSectors := rearrangeByFragmentation(sectors, 6)

	assert.Equal(t, expectedSectors, newSectors)
}

func TestCalculateChecksum(t *testing.T) {
	sectors := []int{0, 2, 2, 1, 1, 1, 2, 2, 2}

	checksum := calculateChecksum(sectors)

	assert.Equal(t, 60, checksum)
}

func TestPartOne(t *testing.T) {
	diskMap := "2333133121414131402"

	freeSpace, sectors := markSectors(diskMap)
	sectors = rearrangeByFragmentation(sectors, freeSpace)
	checksum := calculateChecksum(sectors)

	assert.Equal(t, 1928, checksum)
}

func TestNextBlock(t *testing.T) {
	diskMap := "2333133121414131402"

	// 00...111...2...333.44.5555.6666.777.888899
	_, sectors := markSectors(diskMap)

	index, block := nextBlock(sectors, len(sectors)-1)
	assert.Equal(t, 39, index)
	assert.Equal(t, []int{9, 9}, block)

	index, block = nextBlock(sectors, index)
	assert.Equal(t, 35, index)
	assert.Equal(t, []int{8, 8, 8, 8}, block)

	index, block = nextBlock(sectors, 11)
	assert.Equal(t, 10, index)
	assert.Equal(t, []int{2}, block)
}

func TestNextFreeSpace(t *testing.T) {
	diskMap := "2233133121414134402"

	// 00..111...2...333.44.5555.6666.777....888899
	_, sectors := markSectors(diskMap)

	index, err := nextFreeSpace(sectors, []int{9, 9}, 41)
	assert.Equal(t, 2, index)
	assert.Nil(t, err)

	index, err = nextFreeSpace(sectors, []int{7, 7, 7}, 30)
	assert.Equal(t, 7, index)
	assert.Nil(t, err)

	index, err = nextFreeSpace(sectors, []int{8, 8, 8, 8}, 38)
	assert.Equal(t, 34, index)
	assert.Nil(t, err)

	_, err = nextFreeSpace(sectors, []int{5, 5, 5, 5}, 21)
	assert.NotNil(t, err)
}

func TestPartTwo(t *testing.T) {
	diskMap := "2333133121414131402"

	// 00...111...2...333.44.5555.6666.777.888899
	_, sectors := markSectors(diskMap)
	sectors = rearrangeByShifting(sectors)
	checksum := calculateChecksum(sectors)

	assert.Equal(t, 2858, checksum)
}
