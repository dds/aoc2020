package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dds/aoc2020/lib"
	"github.com/dds/aoc2020/lib/inputs"
)

var ruleRE = regexp.MustCompile(`([^:]+): (\d+)-(\d+) or (\d+)-(\d+)`)

func parse(in string) (r input) {
	// class: 1-3 or 5-7
	// ...
	// \n\n
	parts := strings.Split(in, "\n\n")
	if len(parts) != 3 {
		panic(fmt.Errorf("input doesn't have two double newlines"))
	}
	r.rules = rules{}
	for _, line := range strings.Split(parts[0], "\n") {
		matches := ruleRE.FindStringSubmatch(line)
		if len(matches) == 0 {
			panic(fmt.Errorf("unexpected rule line: %v", line))
		}
		term, min1s, max1s, min2s, max2s := matches[1], matches[2], matches[3], matches[4], matches[5]
		min1, _ := strconv.Atoi(min1s)
		min2, _ := strconv.Atoi(min2s)
		max1, _ := strconv.Atoi(max1s)
		max2, _ := strconv.Atoi(max2s)
		r.rules[term] = rule{min1, max1, min2, max2}
	}
	yourticket := strings.Split(parts[1], "\n")
	if yourticket[0] != "your ticket:" {
		panic(fmt.Errorf("unexpected your ticket line: %v", yourticket[0]))
	}
	for _, s := range lib.CSVParser(yourticket[1]) {
		i, _ := strconv.Atoi(s)
		r.yours = append(r.yours, i)
	}
	nearby := strings.Split(parts[2], "\n")
	if nearby[0] != "nearby tickets:" {
		panic(fmt.Errorf("unexpected nearby tickets line: %v", nearby[0]))
	}
	for _, l := range nearby[1:] {
		parts := lib.CSVParser(l)
		if len(parts) == 0 {
			continue
		}
		t := ticket{}
		for _, s := range parts {
			i, _ := strconv.Atoi(s)
			t = append(t, i)
		}
		r.nearby = append(r.nearby, t)
	}
	return
}

type input struct {
	rules  rules
	yours  ticket
	nearby []ticket
}

type rules map[string]rule

type rule struct {
	min1, max1, min2, max2 int
}

func (r rule) valid(i int) bool {
	if (r.min1 <= i && i <= r.max1) || (r.min2 <= i && i <= r.max2) {
		return true
	}
	return false
}

func (s rules) valid(i int) bool {
	for _, v := range s {
		if v.valid(i) {
			return true
		}
	}
	return false
}

type ticket []int

func part1(in input) (rc int) {
	for _, t := range in.nearby {
		for _, field := range t {
			if !in.rules.valid(field) {
				fmt.Printf("invalid %v\n", field)
				rc += field
			}
		}
	}
	return
}

// Generate all permutations.
// For each permutation, find the first line that is impossible and continue
// Return the ordering that succeeded

func (s rules) perms() (r [][]string) {
	fields := make([]string, len(s))
	i := 0
	for k, _ := range s {
		fields[i] = k
		i++
	}

	var perm func([]string, int)
	perm = func(f []string, i int) {
		if i == len(f) {
			r = append(r, append([]string{}, f...))
			return
		}
		for j := i; j < len(f); j++ {
			f[i], f[j] = f[j], f[i]
			perm(f, i+1)
			f[i], f[j] = f[j], f[i]
		}
		return
	}
	perm(fields, 0)
	return
}

func (s rules) possible(t ticket, ordering []string) bool {
	for i, field := range ordering {
		if !s[field].valid(t[i]) {
			return false
		}
	}
	return true
}

func part2(in input, prefix string) (rc int) {
	filtered := []ticket{}
ticket:
	for _, t := range in.nearby {
		for _, field := range t {
			if !in.rules.valid(field) {
				continue ticket
			}
		}
		filtered = append(filtered, t)
	}
	in.nearby = filtered

	possibleOrderings := in.rules.perms()
	i := 0
	fieldOrder := possibleOrderings[i]

	for _, t := range in.nearby {
		for {
			if !in.rules.possible(t, fieldOrder) {
				i++
				fieldOrder = possibleOrderings[i]
				continue
			}
			break
		}
	}

	rc = 1
	for i, field := range fieldOrder {
		if !strings.HasPrefix(field, prefix) {
			continue
		}
		rc *= in.yours[i]
	}
	return
}

var Input = parse(inputs.Day16())

func main() {
	fmt.Println(part1(Input))
	fmt.Println(part2(Input, "departure"))
}
