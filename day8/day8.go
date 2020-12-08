package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/dds/aoc2020/lib"
	"github.com/dds/aoc2020/lib/inputs"
)

var inputRE = regexp.MustCompile(`^(acc|jmp|nop)\s([-+]\d*)$`)

func parse(s string) []string {
	// 1-4 s: lssss
	matches := inputRE.FindStringSubmatch(s)
	if len(matches) == 0 {
		return matches
	}
	return []string{matches[1], matches[2]}
}

type op struct {
	inst string
	arg  int
}

type state struct {
	stack, accumulator, offset int
}

func part1(input [][]string) (rc int) {
	instructions := map[int]op{}
	for lineNo, row := range input {
		arg, _ := strconv.Atoi(row[1])
		instructions[lineNo] = op{inst: row[0], arg: arg}
	}
	st := state{}
	m := map[int]int{}

	for {
		if m[st.stack] == 1 {
			return st.accumulator
		}
		m[st.stack] = 1
		i := instructions[st.stack]
		switch i.inst {
		case "nop":
			st.stack++
		case "acc":
			st.accumulator += i.arg
			st.stack++
		case "jmp":
			st.stack += i.arg
		}
	}
}

func run(instructions map[int]op) (bool, int) {
	st := state{}

	seenJumps := map[int]int{}
	for {
		if st.stack == len(instructions) {
			return true, st.accumulator
		}
		i := instructions[st.stack]
		switch i.inst {
		case "nop":
			st.stack++
		case "acc":
			st.accumulator += i.arg
			st.stack++
		case "jmp":
			if seenJumps[st.stack] > 0 {
				return false, 0
			}
			seenJumps[st.stack] = 1
			st.stack += i.arg
		}
	}
}

func part2(input [][]string) (rc int) {
	instructions := map[int]op{}
	for lineNo, row := range input {
		arg, _ := strconv.Atoi(row[1])
		instructions[lineNo] = op{inst: row[0], arg: arg}
	}

	for k, v := range instructions {
		orig := instructions[k]
		switch v.inst {
		case "nop":
			instructions[k] = op{inst: "jmp", arg: v.arg}
		case "jmp":
			instructions[k] = op{inst: "nop", arg: v.arg}
		default:
			continue
		}
		res, acc := run(instructions)
		if res {
			return acc
		}
		instructions[k] = orig
	}
	return
}

var Input = lib.ParseInput(inputs.Day8(), parse)

func main() {
	fmt.Println(part1(Input))
	fmt.Println(part2(Input))
}
