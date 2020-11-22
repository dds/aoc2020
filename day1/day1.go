package main

import (
	"fmt"

	"github.com/dds/aoc2020/util"
)

func main() {
	lines, err := util.InputInts(1)
	if err != nil {
		panic(err)
	}
	for i, l := range lines {
		fmt.Println(i, l)
	}
}
