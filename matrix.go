package condensedmatrix

import (
	"fmt"
)

// Matrix is an implementation of CondensedMatrix
// that computes everything on the fly.
type Matrix struct {
	size int
}

// Position finds the indexes given and offset.
func (m Matrix) Position(offset int) (i, j int) {
	return ComputeCoordinates(m.size, offset)
}

// Offset finds the offset of for the given indexes.
func (m Matrix) Offset(i, j int) int {
	p := createPosition(m.size, i, j)
	return p.getOffset(m.size)
}

// ForEachPosition runs the passed callback for each value of the matrix.
func (m Matrix) ForEachPosition(callback func(i, j int)) {
	for i := 0; i < m.size-1; i++ {
		for j := i + 1; j < m.size; j++ {
			callback(i, j)
		}
	}
}

// ForEachPositionMultiThreaded works like ForEachPosition but with routines.
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

// Size returns the size of the square matrix.
func (m Matrix) Size() int {
	return m.size
}

// RealSize return the size of the condensed matrix.
func (m Matrix) RealSize() int {
	return GetRealSize(m.size)
}

// CreateMatrix is the factory for Matrix values.
func CreateMatrix(size int) CondensedMatrix {
	if size < 2 {
		errorMsg := fmt.Sprintf("cached matrix size can not be less than 2 requested : %d", size)
		panic(errorMsg)
	}
	return Matrix{
		size: size,
	}
}
