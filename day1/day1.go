package main

import (
	"fmt"

	"github.com/dds/aoc2020/lib"
	"github.com/dds/aoc2020/lib/inputs"
)

var Input = lib.InputInts(inputs.Day1(), lib.NumberParser)

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

func part2_DictionaryApproach(input [][]int) (rc int) {
	m := map[int][]int{}

	for _, row := range input {
		i := row[0]
		if i > 2020 {
			continue
		}

		for k, v := range m {
			if !(k+i <= 2020 && len(v) < 3) {
				continue
			}
			if len(v) == 2 && v[0]+v[1]+i == 2020 {
				return v[0] * v[1] * i
			}
			m[k+i] = append(v, i)
		}
		if _, ok := m[i]; !ok {
			m[i] = []int{i}
		}
	}
	return
}
