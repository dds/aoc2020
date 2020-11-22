package util

import (
	"fmt"
	"strconv"
	"strings"
)

// All puzzle inputs stored as an array of UTF-8 strings.
var inputs []string

// Returns the input for the given day, each line of the input parsed by the
// parser function into strings. Panics on error.
func Input(day int, parser func(string) []string) [][]string {
	if day > len(inputs) {
		panic(fmt.Errorf("no input for day %v", day))
	}
	lines := strings.Split(inputs[day], "\n")
	r := make([][]string, 0)
	for _, line := range lines {
		fields := parser(line)
		if len(fields) == 0 {
			continue
		}
		r = append(r, fields)
	}

	return r
}

// Returns the input as a two-dimensional array of float64.
func InputNums(day int, parser func(string) []string) [][]float64 {
	lines := Input(day, parser)

	r := make([][]float64, len(lines))
	var err error
	for lineNo, fields := range lines {
		nums := make([]float64, len(fields))
		for i, f := range fields {
			nums[i], err = strconv.ParseFloat(f, 64)
			if err != nil {
				panic(err)
			}
		}
		r[lineNo] = nums
	}

	return r
}

// CSVParser ...
func CSVParser(input string) []string {
	return strings.FieldsFunc(input, func(c rune) bool { return c == ',' })
}

func init() {
	inputs = make([]string, 25)
	inputs[0] = `1,2,3
4,5,6
7,8,9,10
`
	// As the inputs are released, store them right here inline. Simple.
}
