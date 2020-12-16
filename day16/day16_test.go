package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	type test struct {
		input  string
		expect int
	}

	tests := []test{
		test{
			input: `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12
`,
			expect: 71,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			in := parse(test.input)
			require.Equal(t, test.expect, part1(in))
		})
	}
}

func TestPart2(t *testing.T) {
	type input struct {
		prefix string
		data   string
	}
	type test struct {
		input  input
		expect int
	}

	tests := []test{
		test{
			input: input{
				prefix: `row`,
				data: `class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9`,
			},
			expect: 11,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			in := parse(test.input.data)
			require.Equal(t, test.expect, part2(in, test.input.prefix))
		})
	}
}
