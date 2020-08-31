package condensedmatrix

import (
	"fmt"
	"math"
	"testing"
)

func BenchmarkStateless(b *testing.B) {
	for i := 1; i < 3; i++ {
		size := int(math.Pow(10, float64(i*2)))
		realSize := int(GetRealSize(NumberType(size)))
		m := CreateCachedMatrix(size)

		b.Run(fmt.Sprintf("Computed for (size = %d) ", size), func(b *testing.B) {
			b.Run("offset to postion", func(b *testing.B) {
				for o := 0; o < realSize; o++ {
					ComputeCoordinates(NumberType(o), NumberType(size))
				}
			})
			b.Run("postion to offset", func(b *testing.B) {
				for i := 0; i < size; i++ {
					for j := i + 1; j < size; j++ {
						ComputeOffset(NumberType(size), NumberType(i), NumberType(j))
					}
				}
			})
		})
		b.Run(fmt.Sprintf("Cached init for (size = %d) ", size), func(b *testing.B) {
			CreateCachedMatrix(size)
		})
		b.Run(fmt.Sprintf("Cached for (size = %d)", size), func(b *testing.B) {
			b.Run("offset to postion", func(b *testing.B) {
				for o := 0; o < realSize; o++ {
					m.Position(NumberType(o))
				}
			})
			b.Run("postion to offset", func(b *testing.B) {
				for i := 0; i < size; i++ {
					for j := i + 1; j < size; j++ {
						m.Offset(NumberType(i), NumberType(j))
					}
				}
			})
		})
	}
}
