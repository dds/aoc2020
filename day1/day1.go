package main

import (
	"fmt"

	"github.com/dds/aoc2020/util"
)

func main() {
	input, err := util.InputInts(1, util.CSVParser)
	if err != nil {
		panic(err)
	}
	for i, l := range input {
		fmt.Println(i, l)
	}
}
