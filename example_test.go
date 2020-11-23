package main

import (
	"fmt"

	"github.com/dds/aoc2020/util"
)

func Example() {
	input := util.InputInts(util.Inputs[0], func(s string) []string { return util.TrimSpace(util.CSVParser(s)) })
	for i, l := range input {
		fmt.Println(i, l)
	}
	// Output:
	// error parsing line 1 field 3 as float "-": strconv.ParseFloat: parsing "-": invalid syntax
	// 0 [1 2 3]
	// 1 [4 5 6 0]
	// 2 [7 8 9 10]
}
