package condensedmatrix

import (
	"fmt"
	"sort"
	"sync"
	"testing"

	"github.com/potatomasterrace/catch"
	"github.com/stretchr/testify/assert"
)

func assertConsistency(t *testing.T, m CondensedMatrix) {
	realSize := m.RealSize()
	for offset := 0; offset < realSize; offset++ {
		i, j := m.Position(offset)
		computedI, computedJ := ComputeCoordinates(m.Size(), offset)
		assert.Equal(t, computedI, i)
		assert.Equal(t, computedJ, j)
		posI, posJ := m.Position(offset)
		assert.Equal(t, computedI, posI)
		assert.Equal(t, computedJ, posJ)
		retOffset := m.Offset(i, j)
		assert.Equal(t, offset, retOffset)
		assert.Equal(t, offset, ComputeOffset(m.Size(), i, j))
	}
}
func TestAssignDefault(t *testing.T) {
	matrix := CreateCachedMatrix(4)
	m, ok := matrix.(CachedMatrix)
	if !ok {
		panic("wrong return type")
	}
	assert.Equal(t, m.size, 4)
	assert.Equal(t, m.realSize, 6)
	assertConsistency(t, m)
}

func TestForBadSizes(t *testing.T) {
	t.Run("createPosition size<1", func(t *testing.T) {
		panicked, err := catch.Panic(func() {
			createPosition(-42, 0, 1)
		})
		assert.Equal(t, fmt.Sprint(err), "can not have a size of -42")
		assert.True(t, panicked)
	})
	t.Run("CreateCachedMatrix size<1", func(t *testing.T) {
		panicked, err := catch.Panic(func() {
			CreateCachedMatrix(-42)
		})
		assert.Equal(t, fmt.Sprint(err), "cached matrix size can not be less than 2 requested : -42")
		assert.True(t, panicked)
	})
	t.Run("p.i > size", func(t *testing.T) {
		panicked, err := catch.Panic(func() {
			createPosition(5, 41, 42)
		})
		assert.Equal(t, fmt.Sprint(err), "can not have positon at (41,42), exceeding size 5")
		assert.True(t, panicked)
	})
	t.Run("p.i < size-1", func(t *testing.T) {
		panicked, err := catch.Panic(func() {
			createPosition(5, 0, 5)
		})
		assert.True(t, panicked)
		assert.Equal(t, fmt.Sprint(err), "can not have positon at (0,5), exceeding size 5")
	})
}
func TestForEachPosition(t *testing.T) {

	m := CreateCachedMatrix(4)
	t.Run("sync", func(t *testing.T) {
		passedArgs := make([][]int, 0)
		mock := func(i, j int) {
			args := []int{i, j}
			passedArgs = append(passedArgs, args)
		}
		m.ForEachPosition(mock)
		assert.Equal(t, passedArgs, [][]int{[]int{0, 1}, []int{0, 2}, []int{0, 3}, []int{1, 2}, []int{1, 3}, []int{2, 3}})
	})

	t.Run("async", func(t *testing.T) {
		var mutex = &sync.Mutex{}
		passedArgs := []string{}
		mockWithLock := func(i, j int) {
			mutex.Lock()
			defer mutex.Unlock()
			if i == j {
				panic("can have i == j")
			}
			min, max := findMinMax(i, j)
			args := fmt.Sprintf("%d-%d", min, max)
			passedArgs = append(passedArgs, args)
		}

		m.ForEachPositionMultiThreaded(mockWithLock)
		sort.Strings(passedArgs)
		assert.Equal(t, passedArgs, []string([]string{"0-1", "0-2", "0-3", "1-2", "1-3", "2-3"}))
	})
}
