package main

import "testing"

func BenchmarkPart2SortApproach(b *testing.B) {

	for i := 0; i < b.N; i++ {
		part2(Input)
	}
}

func BenchmarkPart2MapAndXorApproach(b *testing.B) {

	for i := 0; i < b.N; i++ {
		part2_shiftapproach(Input)
	}
}
