# Condensed matrix
Golang library for managing condensed matrices

# Usage
## Matrix:
Matrix is the type for managing condensed matrix by computing offsets/values on the fly. 
### Code
```Golang
package main

import (
	"fmt"

	"github.com/potatomasterrace/condensedmatrix"
)

func main() {

	// initiate the condensed matrix
	m := condensedmatrix.CreateMatrix(5)
	// Get the size needed for storing a condensedmatrix of 5 values
	realSize := m.RealSize()
	// getting indexes for each offset.
	fmt.Println("Offsets to positions :")
	for o := 0; o < realSize; o++ {
		i, j := m.Position(o)
		fmt.Printf("\toffset %d should contain values (%d,%d)\r\n", o, i, j)
	}
	// getting offset for each index.
	fmt.Println("Positions to offset  :")
	for i := 0; i < m.Size(); i++ {
		for j := 0; j < m.Size(); j++ {
			// calling offset with values such as (x,x) will cause a panic.
			if i != j {
				offset := m.Offset(i, j)
				fmt.Printf("\tPosition (%d,%d) is stored at offset %d\r\n", i, j, offset)

			}
		}
	}

	// Simpler and faster way to iterate over a condensed matrix
	fmt.Println("Iterating over values :")
	printCurrentPosition := func(i, j int) {
		offset := m.Offset(i, j)
		fmt.Printf("\tPosition (%d,%d) is stored at offset %d\r\n", i, j, offset)
	}
	m.ForEachPositionMultiThreaded(printCurrentPosition)
}

```

### Complexity

| Operation              | Function   | Time | Memory |
|------------------------|------------|------|--------|
| Initialization         |CreateMatrix| O(1) |  O(1)  |
| Offset to position     |m.Position  | O(1) |  O(1)  |
| Coordinates to offset  |m.Offset    | O(1) |  O(1)  |

## CachedMatrix
Matrix is the type for managing condensed matrix by computing offsets/values on the fly. 
### Code
```Golang
package main

import (
	"fmt"

	"github.com/potatomasterrace/condensedmatrix"
)

func main() {

	// initiate the condensed matrix
	c := condensedmatrix.CreateCachedMatrix(5)
	// Get the size needed for storing a condensedmatrix of 5 values
	realSize := c.RealSize()
	// getting indexes for each offset.
	fmt.Println("Offsets to positions :")
	for o := 0; o < realSize; o++ {
		i, j := c.Position(o)
		fmt.Printf("\toffset %d should contain values (%d,%d)\r\n", o, i, j)
	}
	// getting offset for each index.
	fmt.Println("Positions to offset  :")
	for i := 0; i < c.Size(); i++ {
		for j := 0; j < c.Size(); j++ {
			// calling offset with values such as (x,x) will cause a panic.
			if i != j {
				offset := c.Offset(i, j)
				fmt.Printf("\tPosition (%d,%d) is stored at offset %d\r\n", i, j, offset)

			}
		}
	}

	// Simpler and faster way to iterate over a condensed matrix
	fmt.Println("Iterating over values :")
	printCurrentPosition := func(i, j int) {
		offset := c.Offset(i, j)
		fmt.Printf("\tPosition (%d,%d) is stored at offset %d\r\n", i, j, offset)
	}
	c.ForEachPositionMultiThreaded(printCurrentPosition)
}

```
### Complexity

| Operation              | Function   | Time  | Memory  |
|------------------------|------------|-------|---------|
| Initialization         |CreateMatrix| O(n²) |  O(n²)  |
| Offset to position     |m.Position  | O(1)  |  O(1)   |
| Coordinates to offset  |m.Offset    | O(1)  |  O(1)   |


# Benchmark
Both implementations have the same complexity but their run time is different.
## Initialization
```
BenchmarkMatrices/init/baseline_no_op-16                 1000000                 0.000211 ns/op
BenchmarkMatrices/init/cached-16                         1000000               379 ns/op
BenchmarkMatrices/init/computed-16                       1000000                 0.000280 ns/op
```
CachedMatrix's initialization is way slower than Matrix's.

## Offset to position
```
BenchmarkMatrices/offset_to_position/baseline_no_op-16           1000000                12.5 ns/op
BenchmarkMatrices/offset_to_position/cached-16                   1000000               220 ns/op
BenchmarkMatrices/offset_to_position/computed-16                 1000000               573 ns/op
```
CachedMatrix is around 2.5 times faster at finding offsets from positions. 

## Position to offset
```
BenchmarkMatrices/position_to_offset/baseline_no_op-16           1000000                52.4 ns/op
BenchmarkMatrices/position_to_offset/cached-16                   1000000               396 ns/op
BenchmarkMatrices/position_to_offset/computed-16                 1000000               550 ns/op
```
CachedMatrix is around 50% faster at finding positions from offsets.  