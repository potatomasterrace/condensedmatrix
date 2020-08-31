package condensedmatrix

import (
	"fmt"
	"math"
)

type position struct {
	size NumberType
	i    NumberType
	j    NumberType
}

func (p position) format() position {
	p.j, p.i = findMinMax(p.i, p.j)

	if p.i > p.size {
		errorMsg := fmt.Sprintf("can not have positon at (%d,%d), exceeding size %d", p.i, p.j, p.size)
		panic(errorMsg)
	}
	if p.j > p.size-1 {
		errorMsg := fmt.Sprintf("can not have positon at (%d,%d), exceeding size %d", p.i, p.j, p.size)
		panic(errorMsg)
	}
	return p
}

func (p position) getOffset() NumberType {
	return (p.j * p.size) + p.i - (p.j * (p.j + 1) / 2) - p.j - 1
}

func createPosition(s, i, j int) position {
	_s, _i, _j := NumberType(s), NumberType(i), NumberType(j)
	return position{
		size: _s,
		i:    _i,
		j:    _j,
	}
}

func findMinMax(i, j NumberType) (max, min NumberType) {
	if i == j {
		errorMsg := fmt.Sprintf("can not have positon at (%d,%d)", i, j)
		panic(errorMsg)
	}
	if i > j {
		return j, i
	}
	return i, j
}

func GetRealSize(s NumberType) NumberType {
	if s < 2 {
		errorMsg := fmt.Sprintf("can not have a condensed array for %d element", s)
		panic(errorMsg)
	}
	return (s * (s - 1)) / 2
}

func ComputeOffset(s, i, j NumberType) NumberType {
	p := position{
		size: s,
		i:    i,
		j:    j,
	}.format()
	return p.getOffset()
}

func colInRow(s, j NumberType) NumberType {
	return j*(s-1-j) + (j*(j+1))/2
}

func ComputeCoordinates(s, offset NumberType) (i, j NumberType) {
	_s, _o := float64(s), float64(offset)
	j = NumberType(math.Ceil((1/2.)*(-math.Sqrt(-8*_o+4*math.Pow(_s, 2)-4*_s-7)+2*_s-1) - 1))
	i = s - colInRow(s, j+1) + offset
	return i, j
}
