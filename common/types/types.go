package types

type Point struct {
	X int
	Y int
}

type Board struct {
	Plane [][]rune
	MinX  int
	MaxX  int
	MinY  int
	MaxY  int
}

func CreateMatrix(lines []string) [][]rune {
	x := len(lines)

	matrix := make([][]rune, x)
	for i, line := range lines {
		matrix[i] = []rune(line)
	}

	return matrix
}

func CreateBoard(lines []string) Board {
	matrix := CreateMatrix(lines)
	return Board{matrix, 0, len(matrix[0]) - 1, 0, len(matrix) - 1}
}
