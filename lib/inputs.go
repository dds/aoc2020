package lib

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/dds/aoc2020/lib/inputs"
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

	r := make([][]float64, 0)
	for lineNo, fields := range lines {
		nums := make([]float64, 0)
		for i, f := range fields {
			if f == "" {
				continue
			}
			f, err := strconv.ParseFloat(f, 64)
			if err != nil {
				fmt.Printf("error parsing line %v field %v: %v\n", lineNo, i, err)
				continue
			}
			nums = append(nums, f)
		}
		if len(nums) == 0 {
			continue
		}
		r = append(r, nums)
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

var NumberRE = regexp.MustCompile(`[-+]?\d*\.?\d*`)

// NumberParser ...
func NumberParser(input string) []string {
	return NumberRE.FindAllString(input, -1)
}

func init() {
	Inputs = make([]string, 26)
	Inputs[0] = inputs.TestInput1()
	Inputs[1] = inputs.Day1()
	Inputs[2] = inputs.Day2()
	Inputs[3] = inputs.Day3()
	Inputs[4] = inputs.Day4()
	Inputs[5] = inputs.Day5()
	Inputs[6] = inputs.Day6()
	Inputs[7] = inputs.Day7()
	Inputs[8] = inputs.Day8()
	Inputs[9] = inputs.Day9()
	Inputs[10] = inputs.Day10()
	Inputs[11] = inputs.Day11()
	Inputs[12] = inputs.Day12()
	Inputs[13] = inputs.Day13()
	Inputs[14] = inputs.Day14()
	Inputs[15] = inputs.Day15()
	Inputs[16] = inputs.Day16()
	Inputs[17] = inputs.Day17()
	Inputs[18] = inputs.Day18()
	Inputs[19] = inputs.Day19()
	Inputs[20] = inputs.Day20()
	Inputs[21] = inputs.Day21()
	Inputs[22] = inputs.Day22()
	Inputs[23] = inputs.Day23()
	Inputs[24] = inputs.Day24()
	Inputs[25] = inputs.Day25()
}
