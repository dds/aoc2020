package util

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// All puzzle inputs stored as an array of UTF-8 strings.
var Inputs []string

// Parses each line of the input with the given parser function.
func ParseInput(input string, parser func(string) []string) [][]string {
	lines := strings.Split(input, "\n")
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
func InputFloats(input string, parser func(string) []string) [][]float64 {
	lines := ParseInput(input, parser)

	r := make([][]float64, len(lines))
	var err error
	for lineNo, fields := range lines {
		nums := make([]float64, len(fields))
		for i, f := range fields {
			nums[i], err = strconv.ParseFloat(f, 64)
			if err != nil {
				fmt.Printf("error parsing line %v field %v as float %q: %v\n", lineNo, i, f, err)
				nums[i] = math.NaN()
			}
		}
		r[lineNo] = nums
	}

	return r
}

// Returns the input as a two-dimensional array of int.
func InputInts(input string, parser func(string) []string) [][]int {
	lines := InputFloats(input, parser)

	r := make([][]int, len(lines))
	for lineNo, fields := range lines {
		nums := make([]int, len(fields))
		for i, f := range fields {
			if math.IsNaN(f) || math.IsInf(f, 0) {
				nums[i] = 0
				continue
			}
			nums[i] = int(f)
		}
		r[lineNo] = nums
	}

	return r
}

// CSVParser ...
func CSVParser(input string) []string {
	return strings.FieldsFunc(input, func(c rune) bool { return c == ',' })
}

// DashParser ...
func DashParser(input string) []string {
	return strings.FieldsFunc(input, func(c rune) bool { return c == '-' })
}

// Trimmer ...
func TrimSpace(input []string) (r []string) {
	for _, i := range input {
		r = append(r, strings.TrimSpace(i))
	}
	return
}

func init() {
	Inputs = make([]string, 25)
	Inputs[0] = testInput()
	// As the inputs are released, add them to the Inputs slice.

}
