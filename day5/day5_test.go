package main

import "testing"

func BenchmarkPart1AddMulApproach(b *testing.B) {

	for i := 0; i < b.N; i++ {
		part1(Input)
	}
}

func BenchmarkPart1ShiftOrApproach(b *testing.B) {

	for i := 0; i < b.N; i++ {
		part1_ShiftApproach(Input)
	}
}
func BenchmarkPart2SortApproach(b *testing.B) {

	for i := 0; i < b.N; i++ {
		part2(Input)
	}
}

func BenchmarkPart2XorApproach(b *testing.B) {

	for i := 0; i < b.N; i++ {
		part2_XorApproach(Input)
	}
}
