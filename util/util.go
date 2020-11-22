package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Input the input for the day given, each line of the input parsed by the
// parser function into fields of any type. A slice of fields per line is
// returned. Returns input read and error on invalid input.
func Input(day int, parser func(string) ([]string, error)) ([][]string, error) {
	var (
		dir *os.File
		err error
	)
	if dir, err = os.Open("input"); err != nil {
		return nil, err
	}

	var filenames []string
	if filenames, err = dir.Readdirnames(30); err != nil {
		return nil, err
	}

	if day > len(filenames) {
		return nil, fmt.Errorf("no input for day %v", day)
	}

	var b []byte
	if b, err = ioutil.ReadFile("input/" + filenames[day-1]); err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
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

// Give the input for the given day as ints seperated by whitespace.
func InputInts(day int) ([][]int, error) {
	parser := func(input string) ([]string, error) {
		// Split on whitespace into strings.
		r := strings.Fields(input)
		return r, nil
	}

	lines, err := Input(day, parser)
	if err != nil {
		return nil, err
	}

	r := make([][]int, len(lines))
	for lineNo, fields := range lines {
		ints := make([]int, len(fields))
		for i, f := range fields {
			ints[i], err = strconv.Atoi(f)
			if err != nil {
				return r, err
			}
		}
		r[lineNo] = ints
	}

	return r, nil
}
