package main

import (
	"fmt"

	"github.com/dds/aoc2020/util"
)

func Example() {
	input := util.InputNums(0, util.CSVParser)
	for i, l := range input {
		fmt.Println(i, l)
	}
	// Output:
	// 0 [1 2 3]
	// 1 [4 5 6]
	// 2 [7 8 9 10]
}
