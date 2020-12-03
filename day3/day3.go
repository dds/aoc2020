package main

import (
	"fmt"
	"image"
	"strings"
	"testing"

	"github.com/dds/aoc2020/lib"
	"github.com/dds/aoc2020/lib/inputs"
)

var Input = lib.ParseInput(inputs.Day3(), func(s string) []string { return strings.Split(s, "") })

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
	x := 0
	for _, row := range input {
		n := len(row)
		if row[x%n] == "#" {
			rc++
		}
		x += 3
	}
	return
}

func part2(input [][]string) (rc int) {
	counters := map[image.Point]int{
		image.Point{1, 1}: 0,
		image.Point{3, 1}: 0,
		image.Point{5, 1}: 0,
		image.Point{7, 1}: 0,
		image.Point{1, 2}: 0,
	}
	for y, row := range input {
		n := len(row)
		for p := range counters {
			if p.Y == 2 && y%2 == 1 {
				continue
			}
			if row[y/p.Y*p.X%n] == "#" {
				counters[p]++
			}
		}
	}
	rc = 1
	for _, c := range counters {
		rc *= c
	}
	return
}
