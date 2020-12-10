package main

import (
	"fmt"
	"sort"

	"github.com/dds/aoc2020/lib"
	"github.com/dds/aoc2020/lib/inputs"
)

func parse(input [][]int) (r []int) {
	for _, row := range input {
		r = append(r, row[0])
	}
	return
}

var Input = parse(lib.InputInts(inputs.Day10(), lib.NumberParser))

func part1(input []int) (rc int) {
	sort.Ints(input)
	ones, threes := 1, 1
	for i := 0; i < len(input)-1; i++ {
		switch c := input[i+1] - input[i]; c {
		case 3:
			threes++
		case 1:
			ones++
		}
	}
	return threes * ones
}

func part2(input []int) (rc uint64) {
	sort.Ints(input)
	input = append([]int{0}, append(input, input[len(input)-1]+3)...)
	m := map[int]uint64{}
	var next func(int) uint64
	next = func(idx int) uint64 {
		if m[idx] != 0 {
			return m[idx]
		}
		if idx == len(input)-1 {
			return 1
		}
		var sum uint64
		for i := idx + 1; i < len(input); i++ {
			if input[i]-input[idx] > 3 {
				break
			}
			n := next(i)
			sum += n
			m[i] = n
		}
		return sum
	}
	return next(0)
}

func main() {
	fmt.Println(part1(Input))
	fmt.Println(part2(Input))
}
