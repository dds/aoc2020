package main

import (
	"fmt"
	"image"
	"testing"

	"github.com/dds/aoc2020/lib"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	type test struct {
		input  string
		expect int
	}

	tests := []test{
		test{
			input: `F10
N3
F7
R90
F11`,
			expect: 25,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, test.expect, part1(lib.ParseInput(test.input, parse)))
		})
	}
}

func TestPart2(t *testing.T) {
	type test struct {
		input  string
		expect int
	}

	tests := []test{
		test{
			input: `F10
N3
F7
R90
F11`,
			expect: 286,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, test.expect, part2(lib.ParseInput(test.input, parse)))
		})
	}
}

func TestRotation(t *testing.T) {
	type input struct {
		p image.Point
		d int
	}
	type test struct {
		input  input
		expect image.Point
	}

	tests := []test{
		test{
			input:  input{p: image.Point{3, 4}, d: 2},
			expect: image.Point{-3, -4},
		},
		test{
			input:  input{p: image.Point{3, 4}, d: 6},
			expect: image.Point{-3, -4},
		},
		test{
			input:  input{p: image.Point{3, 4}, d: -6},
			expect: image.Point{-3, -4},
		},
		test{
			input:  input{p: image.Point{3, 4}, d: -1},
			expect: image.Point{-4, 3},
		},
		test{
			input:  input{p: image.Point{3, 4}, d: 1},
			expect: image.Point{4, -3},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, test.expect, rotate(test.input.p, test.input.d))
		})
	}
}
