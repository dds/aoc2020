package main

import (
	"fmt"
	"testing"

	"github.com/dds/aoc2020/lib/inputs"
	"github.com/stretchr/testify/require"
)

func BenchmarkSliceApproach(b *testing.B) {
	input := parse(inputs.Day15())
	for i := 0; i < b.N; i++ {
		solveSliceApproach(input, 30000000)
	}
}

func BenchmarkIntApproach(b *testing.B) {
	input := parse(inputs.Day15())
	for i := 0; i < b.N; i++ {
		solveIntApproach(input, 30000000)
	}
}

func TestPart1(t *testing.T) {
	type test struct {
		input  []int
		expect int
	}

	tests := []test{
		test{
			input:  []int{1, 3, 2},
			expect: 1,
		},
		test{
			input:  []int{2, 1, 3},
			expect: 10,
		},
		test{
			input:  []int{1, 2, 3},
			expect: 27,
		},
		test{
			input:  []int{2, 3, 1},
			expect: 78,
		},
		test{
			input:  []int{3, 2, 1},
			expect: 438,
		},
		test{
			input:  []int{3, 1, 2},
			expect: 1836,
		},
		test{
			input:  []int{0, 3, 6},
			expect: 436,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, test.expect, part1(test.input))
		})
	}
}
