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
			input:  `2 * 3 + (4 * 5)`,
			expect: 26,
		},
		test{
			input:  `5 + (8 * 3 + 9 + 3 * 4 * 3)`,
			expect: 437,
		},
		test{
			input:  `5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))`,
			expect: 12240,
		},
		test{
			input:  `((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`,
			expect: 13632,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, test.expect, part1(test.input))
		})
	}
}
