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
			input: `939
7,13,x,x,59,x,31,19
`,
			expect: 295,
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
		input  string
		expect uint64
	}

	tests := []test{
		// 23,x,x,x,x,x,x,x,x,x,x,x,x,41,x,x,x,37,x,x,x,x,x,421,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,17,x,19,x,x,x,x,x,x,x,x,x,29,x,487,x,x,x,x,x,x,x,x,x,x,x,x,13
		test{
			input: `939
23,x,x,x,x,x,x,x,x,x,x,x,x,41,x,x,x,37,x,x,x,x,x,421,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,17,x,19,x,x,x,x,x,x,x,x,x,29
`,
			expect: 111652513668,
			//	10527928,
		},
		test{
			input: `939
2,x,7
`,
			expect: 12,
		},
		test{
			input: `939
3,5,x,7
`,
			expect: 39,
		},
		test{
			input: `939
17,13
`,
			expect: 51,
		},
		test{
			input: `939
13,19
`,
			expect: 208,
		},
		test{
			input: `939
17,13,x,19
`,
			expect: 1156,
		},
		test{
			input: `939
17,13,19
`,
			expect: 2924,
		},
		test{
			input: `939
13,17,19
`,
			expect: 169,
		},
		test{
			input: `939
7,13,x,x,59,x,31,19
`,
			expect: 1068781,
		},
		test{
			input: `939
17,x,13,19
`,
			expect: 3417,
		},
		test{
			input: `939
1789,37,47,1889
`,
			expect: 1202161486,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			r := part2(test.input)
			require.Equal(t, test.expect, r, fmt.Sprintf("%d %d", test.expect, r))
			fmt.Println(test.expect)
		})
	}
}
