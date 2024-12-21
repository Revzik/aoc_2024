package structures

type Vector struct {
	X int
	Y int
}

func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{v1.X + v2.X, v1.Y + v2.Y}
}

type Board struct {
	Plane [][]rune
}

func (b *Board) MinX() int {
	return 0
}

func (b *Board) MinY() int {
	return 0
}

func (b *Board) MaxX() int {
	return len(b.Plane[0])
}

func (b *Board) MaxY() int {
	return len(b.Plane)
}

func (b *Board) Get(v Vector) rune {
	return b.Plane[v.Y][v.X]
}

func (b *Board) Set(v Vector, r rune) {
	b.Plane[v.Y][v.X] = r
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
	return Board{Plane: CreateMatrix(lines)}
}