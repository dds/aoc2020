package main

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/dds/aoc2020/lib"
	"github.com/dds/aoc2020/lib/inputs"
)

var Input = lib.ParseInput(inputs.Day5(), func(s string) []string { return strings.Split(s, "") })

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

func part1(input [][]string) (rc int) {
	for _, line := range input {
		rows := line[:7]
		var row int8 = 0
		for i := 6; i >= 0; i-- {
			switch rows[i] {
			case "B":
				row |= int8(1 << (6 - i))
			}
		}
		seats := line[len(rows):]
		var seat int8 = 0
		for i := 2; i >= 0; i-- {
			switch seats[i] {
			case "R":
				seat |= int8(1 << (2 - i))
			}
		}
		id := int(row)*8 + int(seat)
		rc = lib.Max(id, rc)
	}
	return
}

func part2(input [][]string) (rc int) {
	ids := []int{}
	for _, line := range input {
		rows := line[:7]
		var row int8 = 0
		for i := 6; i >= 0; i-- {
			switch rows[i] {
			case "B":
				row |= int8(1 << (6 - i))
			}
		}
		seats := line[len(rows):]
		var seat int8 = 0
		for i := 2; i >= 0; i-- {
			switch seats[i] {
			case "R":
				seat |= int8(1 << (2 - i))
			}
		}
		id := int(row)*8 + int(seat)
		ids = append(ids, id)
	}
	sort.Ints(ids)
	for i := 0; i < len(ids); i++ {
		if ids[i+1] != ids[i]+1 {
			return ids[i] + 1
		}
	}
	return
}
