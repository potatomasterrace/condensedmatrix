package condensedmatrix

import (
	"fmt"
	"math"
)

func findMinMax(i, j int) (max, min int) {
	if i == j {
		errorMsg := fmt.Sprintf("can not have positon at (%d,%d)", i, j)
		panic(errorMsg)
	}
	if i > j {
		return j, i
	}
	return i, j
}

type position struct {
	i int
	j int
}

func (p position) format(size int) position {
	p.i, p.j = findMinMax(p.i, p.j)
	if size < 1 {
		errorMsg := fmt.Sprintf("can not have a size of %d", size)
		panic(errorMsg)
	}
	if p.j > size-1 {
		errorMsg := fmt.Sprintf("can not have positon at (%d,%d), exceeding size %d", p.i, p.j, size)
		panic(errorMsg)
	}
	return p
}

func (p position) getOffset(size int) int {
	return (p.i * size) + p.j - (p.i * (p.i + 1) / 2) - p.i - 1
}

func createPosition(s, i, j int) position {
	return position{
		i: i,
		j: j,
	}.format(s)
}

// GetRealSize computes the number of elements inside a condensed matrix of size s.
func GetRealSize(s int) int {
	if s < 2 {
		errorMsg := fmt.Sprintf("can not have a condensed array for %d element", s)
		panic(errorMsg)
	}
	return (s * (s - 1)) / 2
}

// ComputeOffset computes the offset the coordinates (i,j) inside a condensed matrix of size s.
func ComputeOffset(s, i, j int) int {
	p := createPosition(s, i, j)
	return p.getOffset(s)
}

func colInRow(s, j int) int {
	return j*(s-1-j) + (j*(j+1))/2
}

// ComputeCoordinates computes the coordinates at the offset inside a condensed matrix of size s.
func ComputeCoordinates(s, offset int) (i, j int) {
	_s, _o := float64(s), float64(offset)
	i = int(math.Ceil((1/2.)*(-math.Sqrt(-8*_o+4*math.Pow(_s, 2)-4*_s-7)+2*_s-1) - 1))
	j = s - colInRow(s, i+1) + offset
	return i, j
}
