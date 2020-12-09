package main

import (
	"fmt"

	"github.com/dds/aoc2020/lib"
	"github.com/dds/aoc2020/lib/inputs"
)

func parse(input [][]int) (r []int) {
	for _, row := range input {
		r = append(r, row[0])
	}
	return
}

var Input = parse(lib.InputInts(inputs.Day9(), lib.NumberParser))

func validNumbers(in []int, preambleSize int) (r []int) {
	if len(in) < preambleSize {
		return
	}
	for i := 0; i < preambleSize-1; i++ {
		for j := i + 1; j < preambleSize; j++ {
			r = append(r, in[i]+in[j])
		}
	}
	return
}

func contains(in []int, needle int) bool {
	for _, i := range in {
		if i == needle {
			return true
		}
	}
	return false
}

func part1(input []int, preambleSize int) (rc int) {
	for i := preambleSize; i < len(input); i++ {
		candidates := validNumbers(input[i-preambleSize:], preambleSize)
		if !contains(candidates, input[i]) {
			return input[i]
		}
	}
	return
}

func part2(input []int, preambleSize int) (rc int) {
	target := part1(input, preambleSize)
loop:
	for i := 0; i < len(input)-1; i++ {
		max, min, sum := 0, 1<<31, 0
		for j := i; j < len(input); j++ {
			t := input[j]
			sum += t
			if sum > target {
				continue loop
			}
			min = lib.Min(min, t)
			max = lib.Max(max, t)
			if sum == target {
				return min + max
			}
		}
	}
	return
}
func main() {
	fmt.Println(part1(Input, 25))
	fmt.Println(part2(Input, 25))
}
