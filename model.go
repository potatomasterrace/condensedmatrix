package condensedmatrix

import "fmt"

// NumberType is data type used for indexes and offsets
type NumberType uint64

// Offset is the type for finding a position offset in the array
type Offset NumberType

type CachedMatrix struct {
	Size      NumberType
	RealSize  NumberType
	offsets   []position
	positions [][]NumberType
}

func (c CachedMatrix) Position(offset NumberType) (i, j NumberType) {
	p := c.offsets[offset]
	return p.i, p.j
}

func (c CachedMatrix) Offset(i, j NumberType) NumberType {
	return c.positions[i][j-i]
}

func (c CachedMatrix) ForEachPosition(callback func(i, j NumberType), sync bool) {
	for _, p := range c.offsets {
		if sync {
			callback(p.i, p.j)
		} else {
			go callback(p.i, p.j)
		}
	}
}

func CreateCachedMatrix(size int) CachedMatrix {
	if size < 2 {
		errorMsg := fmt.Sprintf("cached matrix size can not be less than 2 requested : %d", size)
		panic(errorMsg)
	}
	realSize := int(size * (size - 1) / 2)
	offset := make([]position, realSize)
	positions := make([][]NumberType, size)
	var cnt NumberType = 0
	for i := 0; i < size; i++ {
		col := make([]NumberType, size-i)
		for j := i + 1; j < size; j++ {
			col[j-i] = cnt
			offset[cnt] = createPosition(size, i, j)
			cnt++
		}
		positions[i] = col
	}

	_size, _realSize := NumberType(size), NumberType(realSize)
	return CachedMatrix{
		Size:      _size,
		RealSize:  _realSize,
		offsets:   offset,
		positions: positions,
	}
}
