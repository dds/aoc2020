package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1(Input)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2(Input)
	}
}

func TestPart1(t *testing.T) {
	type test struct {
		input  []int
		expect int
	}

	tests := []test{
		test{
			input: []int{
				28,
				33,
				18,
				42,
				31,
				14,
				46,
				20,
				48,
				47,
				24,
				23,
				49,
				45,
				19,
				38,
				39,
				11,
				1,
				32,
				25,
				35,
				8,
				17,
				7,
				9,
				4,
				2,
				34,
				10,
				3,
			},
			expect: 220,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, test.expect, part1(test.input))
		})
	}
}

func TestPart2(t *testing.T) {
	type test struct {
		input  []int
		expect uint64
	}

	tests := []test{
		test{
			input: []int{16,
				1,
				4,
				5,
				6,
				7,
				10,
				11,
				12,
				15,
				19,
			},
			expect: uint64(8),
		},
		test{
			input: []int{
				28,
				33,
				18,
				42,
				31,
				14,
				46,
				20,
				48,
				47,
				24,
				23,
				49,
				45,
				19,
				38,
				39,
				11,
				1,
				32,
				25,
				35,
				8,
				17,
				7,
				9,
				4,
				2,
				34,
				10,
				3,
			},
			expect: uint64(19208),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, test.expect, part2(test.input))
		})
	}
}
