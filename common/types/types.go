package types

type Point struct {
	X int
	Y int
}

func (p1 *Point) Add(p2 Point) Point {
	return Point{p1.X + p2.X, p1.Y + p2.Y}
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

func CopyMatrix(matrix [][]rune) [][]rune {
	dst := make([][]rune, len(matrix))
	for i := range matrix {
		dst[i] = make([]rune, len(matrix[i]))
		copy(dst[i], matrix[i])
	}
	return dst
}

func CreateBoard(lines []string) Board {
	matrix := CreateMatrix(lines)
	return Board{matrix, 0, len(matrix[0]) - 1, 0, len(matrix) - 1}
}
