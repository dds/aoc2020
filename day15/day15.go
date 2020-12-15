package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dds/aoc2020/lib/inputs"
)

func parse(s string) (r []int) {
	for _, i := range strings.Split(s, ",") {
		n, _ := strconv.Atoi(i)
		r = append(r, n)
	}
	return
}

var Input = parse(inputs.Day15())

func solve(input []int, limit int) (rc int) {
	m := map[int][]int{}
	for i, x := range input {
		m[x] = append(m[x], i)
	}
	i := len(input) - 1
	t := input[len(input)-1]
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

func part1(input []int) (rc int) {
	return solve(input, 2019)
}
func part2(input []int) (rc int) {
	return solve(input, 29999999)
}

func main() {
	fmt.Println(part1(Input))
	fmt.Println(part2(Input))
}
