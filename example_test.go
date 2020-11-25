package main

import (
	"fmt"

	"github.com/dds/aoc2020/lib"
)

func Example() {
	input := lib.InputInts(lib.Inputs[0], lib.NumberParser)
	for i, l := range input {
		fmt.Println(i, l)
	}
	// Output:
	// error parsing line 1 field 3: strconv.ParseFloat: parsing "-": invalid syntax
	// 0 [1 2 3]
	// 1 [4 5 6]
	// 2 [7 8 9 10]
}
