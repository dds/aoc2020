package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1(Input, 25)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2(Input, 25)
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
				35,
				20,
				15,
				25,
				47,
				40,
				62,
				55,
				65,
				95,
				102,
				117,
				150,
				182,
				127,
				219,
				299,
				277,
				309,
				576,
			},
			expect: 127,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, test.expect, part1(test.input, 5))
		})
	}
}

func TestPart2(t *testing.T) {
	type test struct {
		input  []int
		expect int
	}

	tests := []test{
		test{
			input: []int{
				35,
				20,
				15,
				25,
				47,
				40,
				62,
				55,
				65,
				95,
				102,
				117,
				150,
				182,
				127,
				219,
				299,
				277,
				309,
				576,
			},
			expect: 62,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, test.expect, part2(test.input, 5))
		})
	}
}
