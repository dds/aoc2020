package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/dds/aoc2020/lib/inputs"
)

func parse(input string) [][]string {
	lines := strings.Split(input, "\n\n")
	r := make([][]string, 0)
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		r = append(r, fields)
	}

	return r
}

var Input = parse(inputs.Day6())

func part1(input [][]string) (rc int) {
	for _, r := range input {
		m := map[rune]int{}
		for _, c := range strings.Join(r, "") {
			m[c]++
		}
		rc += len(m)
	}
	return
}

func part2(input [][]string) (rc int) {
	for _, r := range input {
		m := map[rune]int{}
		for _, c := range strings.Join(r, "") {
			m[c]++
			if m[c] == len(r) {
				rc++
			}
		}
	}
	return
}

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
