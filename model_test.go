package condensedmatrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertConsistency(t *testing.T, m CachedMatrix) {
	for offset, position := range m.offsets {
		i, j := m.Position(NumberType(offset))
		assert.Equal(t, i, position.i)
		assert.Equal(t, j, position.j)
		retOffset := m.Offset(i, j)
		assert.Equal(t, offset, int(retOffset))
	}
}

func TestAssignDefault(t *testing.T) {
	m := CreateCachedMatrix(4)
	assert.Equal(t, int(m.Size), 4)
	assert.Equal(t, int(m.RealSize), 6)
	assertConsistency(t, m)
}
