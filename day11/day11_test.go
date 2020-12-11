package main

import (
	"fmt"
	"strings"
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
			input: `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`,
			expect: 37,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			i := lib.ParseInput(test.input, func(s string) []string { return strings.Split(s, "") })
			require.Equal(t, test.expect, part1(i))
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
			input: `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`,
			expect: 26,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			i := lib.ParseInput(test.input, func(s string) []string { return strings.Split(s, "") })
			require.Equal(t, test.expect, part2(i))
		})
	}
}
