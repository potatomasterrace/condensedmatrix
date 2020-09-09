package condensedmatrix

import (
	"fmt"
	"sort"
	"sync"
	"testing"

	"github.com/potatomasterrace/catch"
	"github.com/stretchr/testify/assert"
)

func TestCreateMatrix(t *testing.T) {
	panicked, err := catch.Panic(func() {
		CreateMatrix(-42)
	})
	assert.Equal(t, fmt.Sprint(err), "cached matrix size can not be less than 2 requested : -42")
	assert.True(t, panicked)
}

func TestMatrix(t *testing.T) {
	t.Run("TestAssignDefault", func(t *testing.T) {
		matrix := CreateMatrix(4)
		m, ok := matrix.(Matrix)
		if !ok {
			panic("wrong return type")
		}
		assert.Equal(t, m.size, 4)
		assertConsistency(t, m)
	})

	t.Run("TestForEachPosition", func(t *testing.T) {
		m := CreateMatrix(4)
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
	})
}
