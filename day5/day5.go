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
	fmt.Println(part2_shiftapproach(Input))
}

func part1(input [][]string) (rc int) {
	for _, line := range input {
		rows := line[:7]
		row := 0
		for i := 0; i < 7; i++ {
			switch rows[i] {
			case "B":
				row |= 1 << (6 - i)
			}
		}
		seats := line[len(rows):]
		seat := 0
		for i := 0; i < 3; i++ {
			switch seats[i] {
			case "R":
				seat |= 1 << (2 - i)
			}
		}
		id := row*8 + seat
		rc = lib.Max(id, rc)
	}
	return
}

func part2(input [][]string) (rc int) {
	var ids []int
	for _, line := range input {
		rows := line[:7]
		row := 0
		for i := 0; i < 7; i++ {
			switch rows[i] {
			case "B":
				row |= 1 << (6 - i)
			}
		}
		seats := line[len(rows):]
		seat := 0
		for i := 0; i < 3; i++ {
			switch seats[i] {
			case "R":
				seat |= 1 << (2 - i)
			}
		}
		id := row*8 + seat
		ids = append(ids, id)
	}
	sort.Ints(ids)
	for i := 0; i < len(ids)-1; i++ {
		if ids[i+1] != ids[i]+1 {
			return ids[i] + 1
		}
	}
	return
}

func part2_shiftapproach(input [][]string) (rc int) {
	ids := map[int]struct{}{}
	for _, line := range input {
		rows := line[:7]
		row := 0
		for i := 0; i < 7; i++ {
			switch rows[i] {
			case "B":
				row |= 1 << (6 - i)
			}
		}
		seats := line[len(rows):]
		seat := 0
		for i := 0; i < 3; i++ {
			switch seats[i] {
			case "R":
				seat |= 1 << (2 - i)
			}
		}
		id := (row << 3) | seat
		ids[id] = struct{}{}
	}
	var r int
	for k, _ := range ids {
		r ^= k
	}
	return r + 1
}
