package condensedmatrix

import (
	"testing"
)

func getPositions(size int) [][]int {
	pos := [][]int{}
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			pos = append(pos, []int{i, j})
		}
	}
	return pos
}
func doNothing(v int) {
}
func BenchmarkMatrices(b *testing.B) {
	size := 10000
	realSize := GetRealSize(size)

	b.Run("init", func(b *testing.B) {
		b.Run("baseline_no_op", func(b *testing.B) {
			doNothing(size)
		})
		b.Run("cached", func(b *testing.B) {
			CreateCachedMatrix(size)
		})
		b.Run("computed", func(b *testing.B) {
			CreateMatrix(size)
		})
	})
	b.Run("offset to position", func(b *testing.B) {
		computed := CreateMatrix(size)
		cached := CreateCachedMatrix(size)
		b.Run("baseline_no_op", func(b *testing.B) {
			for o := 0; o < realSize; o++ {
				doNothing(o)
			}
		})
		b.Run("cached", func(b *testing.B) {
			for o := 0; o < realSize; o++ {
				cached.Position(o)
			}
		})
		b.Run("computed", func(b *testing.B) {
			for o := 0; o < realSize; o++ {
				computed.Position(o)
			}
		})

	})
	b.Run("position to offset", func(b *testing.B) {
		computed := CreateMatrix(size)
		cached := CreateCachedMatrix(size)
		args := getPositions(size)
		b.Run("baseline_no_op", func(b *testing.B) {
			for _, v := range args {
				doNothing(v[0])
			}
		})
		b.Run("cached", func(b *testing.B) {
			for _, v := range args {
				cached.Offset(v[0], v[1])
			}
		})
		b.Run("computed", func(b *testing.B) {
			for _, v := range args {
				computed.Offset(v[0], v[1])
			}
		})

	})

}
