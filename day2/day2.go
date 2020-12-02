package main

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/dds/aoc2020/lib"
	"github.com/dds/aoc2020/lib/inputs"
)

func parse(s string) []string {
	// 1-4 s: lssss
	re := regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)
	matches := re.FindStringSubmatch(s)
	if len(matches) == 0 {
		return matches
	}
	return matches[1:]
}

var Input = lib.ParseInput(inputs.Day2(), parse)

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
	for _, data := range input {
		a, b := data[0], data[1]
		min, err := strconv.Atoi(a)
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(b)
		if err != nil {
			panic(err)
		}
		c := data[2]
		n := 0
		for _, s := range data[3] {
			if string(s) == c {
				n++
			}
		}
		if min <= n && n <= max {
			rc++
		}
	}
	return
}

func part2(input [][]string) (rc int) {
	for _, data := range input {
		a, b := data[0], data[1]
		idx1, err := strconv.Atoi(a)
		if err != nil {
			panic(err)
		}
		idx1--
		idx2, err := strconv.Atoi(b)
		if err != nil {
			panic(err)
		}
		idx2--
		c := byte(data[2][0])
		if (data[3][idx1] == c && data[3][idx2] != c) || (data[3][idx1] != c && data[3][idx2] == c) {
			rc++
		}
	}
	return
}
