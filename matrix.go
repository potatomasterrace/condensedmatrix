package condensedmatrix

import (
	"fmt"
)

type Matrix struct {
	size int
}

func (m Matrix) Position(offset int) (i, j int) {
	return ComputeCoordinates(m.size, offset)
}

func (m Matrix) Offset(i, j int) int {
	p := createPosition(m.size, i, j)
	return p.getOffset(m.size)
}

func (m Matrix) ForEachPosition(callback func(i, j int)) {
	for i := 0; i < m.size-1; i++ {
		for j := i + 1; j < m.size; j++ {
			callback(i, j)
		}
	}
}

func (m Matrix) ForEachPositionMultiThreaded(callback func(i, j int)) {
	done := make(chan bool)
	for i := 0; i < m.size-1; i++ {
		for j := i + 1; j < m.size; j++ {
			argI, argJ := i, j
			go func() {
				callback(argI, argJ)
				done <- true
			}()
		}
	}
	realSize := m.RealSize()
	// wait for routines to finish
	for i := 0; i < realSize; i++ {
		<-done
	}
}

func (m Matrix) Size() int {
	return m.size
}

func (m Matrix) RealSize() int {
	return GetRealSize(m.size)
}

func CreateMatrix(size int) CondensedMatrix {
	if size < 2 {
		errorMsg := fmt.Sprintf("cached matrix size can not be less than 2 requested : %d", size)
		panic(errorMsg)
	}
	return Matrix{
		size: size,
	}
}
