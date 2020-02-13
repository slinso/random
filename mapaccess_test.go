package main_test

import (
	"testing"
)

func generateMap(size int) map[int]int {
	slice := make(map[int]int, size)
	for i := 0; i < size; i++ {
		slice[i] = i
	}

	return slice
}

var result int

func BenchmarkMapAccess(b *testing.B) {
	chosenPizzasIndices := generateMap(5000)
	skippedPizzasIndices := generateMap(200)
	skippedPizzas := generateMap(200)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := len(chosenPizzasIndices) - 1; i >= 0; i-- {
			for j := len(skippedPizzasIndices) - 1; j >= 0; j-- {
				for k := j - 1; k >= 0; k-- {
					result = skippedPizzas[skippedPizzasIndices[j]] + skippedPizzas[skippedPizzasIndices[k]]
					if result == skippedPizzas[k] {
						skippedPizzas[skippedPizzasIndices[j]] = skippedPizzas[skippedPizzasIndices[k]]
					}
				}
			}
		}
	}
}
