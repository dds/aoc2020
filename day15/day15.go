package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dds/aoc2020/lib/inputs"
)

func parse(s string) (r []int) {
	for _, i := range strings.Split(s, ",") {
		n, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		r = append(r, n)
	}
	return
}

var Input = parse(inputs.Day15())

func solveSliceApproach(input []int, limit int) (rc int) {
	m := map[int][]int{}
	for i, x := range input {
		m[x] = append(m[x], i)
	}
	i := len(input) - 1
	t := input[i]
	for i < limit {
		i++
		if seen := m[t]; len(seen) < 2 {
			t = 0
			m[0] = append(m[0], i)
			continue
		}
		n := len(m[t]) // number of times we have seen this item
		m[t] = []int{m[t][n-2], m[t][n-1]}
		t = m[t][1] - m[t][0] // how many turns apart
		m[t] = append(m[t], i)
	}
	return t
}

func solveIntApproach(input []int, limit int) (rc int) {
	m := map[int]int{}
	for i, x := range input {
		m[x] = i + 1
	}
	t := input[len(input)-1]
	i := len(input)
	fmt.Println(input)
	for i < limit {
		i++
		var a, b int
		if m[t]>>32 == 0 {
			t = 0
		} else {
			a = m[t] >> 32 & 0xFFFFFFFF
			b = m[t] & 0xFFFFFFFF
			t = b - a
		}
		m[t] = m[t]<<32 | i
	}
	return t
}

func part1(input []int) (rc int) {
	return solveIntApproach(input, 2020)
}
func part2(input []int) (rc int) {
	return solveIntApproach(input, 30000000)
}

func main() {
	fmt.Println(part1(Input))
	fmt.Println(part2(Input))
}
