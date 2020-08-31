package condensedmatrix

import (
	"fmt"
	"testing"

	"github.com/potatomasterrace/catch"
	"github.com/stretchr/testify/assert"
)

func TestMinMax(t *testing.T) {
	t.Run("return values", func(t *testing.T) {
		min, max := findMinMax(0, 42)
		assert.Equal(t, int(min), 0)
		assert.Equal(t, int(max), 42)
		min, max = findMinMax(42, 0)
		assert.Equal(t, int(min), 0)
		assert.Equal(t, int(max), 42)
	})
	t.Run("panic", func(t *testing.T) {
		panicked, err := catch.Panic(func() {
			findMinMax(1, 1)
		})
		assert.Equal(t, fmt.Sprint(err), "can not have positon at (1,1)")
		assert.True(t, panicked)

	})
}

func TestGetRealSize(t *testing.T) {

	t.Run("return values", func(t *testing.T) {
		size := GetRealSize(50)
		assert.Equal(t, int(size), 1225)
		size = GetRealSize(15)
		assert.Equal(t, int(size), 105)
	})

	t.Run("panic", func(t *testing.T) {
		panicked, err := catch.Panic(func() {
			GetRealSize(0)
		})
		assert.Equal(t, fmt.Sprint(err), "can not have a condensed array for 0 element")
		assert.True(t, panicked)
		panicked, err = catch.Panic(func() {
			GetRealSize(1)
		})
		assert.Equal(t, fmt.Sprint(err), "can not have a condensed array for 1 element")
		assert.True(t, panicked)
		size := GetRealSize(2)
		assert.Equal(t, int(size), 1)
	})
}

func TestComputeOffset(t *testing.T) {
	for size := 2; size < 100; size++ {
		realSize := int(GetRealSize(NumberType(size)))
		t.Run("offsets", func(t *testing.T) {
			for o := 0; o < realSize; o++ {
				i, j := ComputeCoordinates(NumberType(size), NumberType(o))
				computed_offset := ComputeOffset(NumberType(size), i, j)
				assert.Equal(t, int(computed_offset), o)
			}
		})
	}
}
