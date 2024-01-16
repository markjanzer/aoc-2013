package main

import (
	"advent-of-code-2023/lib"
	"testing"
)

func BenchmarkSolvePart2(b *testing.B) {
	dataString := lib.GetDataString(DataFile)
	for i := 0; i < b.N; i++ {
		solvePart2(dataString)
	}
}
