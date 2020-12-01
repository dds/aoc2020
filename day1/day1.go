package main

import (
	"fmt"
	"testing"

	"github.com/dds/aoc2020/lib"
	"github.com/dds/aoc2020/lib/inputs"
)

var Input = lib.InputInts(inputs.Day1(), lib.NumberParser)

func Test(t *testing.T) {
	// type test struct {
	// 	input  int
	// 	expect int
	// }

	// tests := []test{
	// 	test{
	// 		// ...
	// 	},
	// }

	// for i, test := range tests {
	// 	t.Run(fmt.Sprint(i), func(t *testing.T) {
	// 		require.Equal(t, test.expect, test.input)
	// 	})
	// }
}

func main() {
	fmt.Println(part1(Input))
	fmt.Println(part2(Input))
}

func part1(input [][]int) (rc int) {
	for i := 0; i < len(input)-1; i++ {
		for j := i; j < len(input); j++ {
			if input[i][0]+input[j][0] == 2020 {
				return input[i][0] * input[j][0]
			}
		}
	}
	return
}

func part2(input [][]int) (rc int) {
	for i := 0; i < len(input)-2; i++ {
		for j := i; j < len(input)-1; j++ {
			for k := j; k < len(input); k++ {
				if input[i][0]+input[j][0]+input[k][0] == 2020 {
					return input[i][0] * input[j][0] * input[k][0]
				}
			}
		}
	}
	return
}
