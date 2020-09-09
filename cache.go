package condensedmatrix

import "fmt"

type CachedMatrix struct {
	size      int
	realSize  int
	offsets   []position
	positions [][]int
}

func (c CachedMatrix) Position(offset int) (i, j int) {
	p := c.offsets[offset]
	return p.i, p.j
}

func (c CachedMatrix) Offset(i, j int) int {
	min, max := findMinMax(i, j)
	return c.positions[min][max-min-1]
}

func (c CachedMatrix) ForEachPosition(callback func(i, j int)) {
	for _, p := range c.offsets {
		callback(p.i, p.j)
	}
}

func (c CachedMatrix) ForEachPositionMultiThreaded(callback func(i, j int)) {
	done := make(chan bool)
	for _, p := range c.offsets {
		i, j := p.i, p.j
		go func() {
			callback(i, j)
			done <- true
		}()
	}
	// wait for routines to finish
	for i := 0; i < c.realSize; i++ {
		<-done
	}
}
func (c CachedMatrix) Size() int {
	return c.size
}
func (c CachedMatrix) RealSize() int {
	return c.realSize
}

func CreateCachedMatrix(size int) CondensedMatrix {
	if size < 2 {
		errorMsg := fmt.Sprintf("cached matrix size can not be less than 2 requested : %d", size)
		panic(errorMsg)
	}
	realSize := int(size * (size - 1) / 2)
	offset := make([]position, realSize)
	positions := make([][]int, size-1)
	var cnt int = 0
	for i := 0; i < size-1; i++ {
		col := make([]int, size-i-1)
		for j := i + 1; j < size; j++ {
			col[j-i-1] = cnt
			offset[cnt] = createPosition(size, i, j)
			cnt++
		}
		positions[i] = col
	}

	return CachedMatrix{
		size:      size,
		realSize:  realSize,
		offsets:   offset,
		positions: positions,
	}
}
