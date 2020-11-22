package util

import (
	"fmt"
	"strconv"
	"strings"
)

// All puzzle inputs stored as an array of UTF-8 strings.
var inputs []string

// Returns the input for the given day, each line of the input parsed by the
// parser function into strings. On error, returns nil or as much of the input
// read so far and the error.
func Input(day int, parser func(string) ([]string, error)) ([][]string, error) {
	day = day - 1 // array index 0 == day 1
	if day > len(inputs) {
		return nil, fmt.Errorf("no input for day %v", day)
	}
	lines := strings.Split(inputs[day], "\n")
	r := make([][]string, len(lines))
	for lineNo, line := range lines {
		fields, err := parser(line)
		if err != nil {
			return r, err
		}
		r[lineNo] = fields
	}

	return r, nil
}

// Returns the input as a two-dimensional array of float64.
func InputNums(day int, parser func(string) ([]string, error)) ([][]float64, error) {
	lines, err := Input(day, parser)
	if err != nil {
		return nil, err
	}

	r := make([][]float64, len(lines))
	for lineNo, fields := range lines {
		nums := make([]float64, len(fields))
		for i, f := range fields {
			nums[i], err = strconv.ParseFloat(f, 64)
			if err != nil {
				return r, err
			}
		}
		r[lineNo] = nums
	}

	return r, nil
}

// CSVParser ...
func CSVParser(input string) ([]string, error) {
	r := strings.FieldsFunc(input, func(c rune) bool { return c == ',' })
	return r, nil
}
