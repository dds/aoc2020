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

func (s rules) filter(tickets []ticket) (r []ticket) {
t:
	for _, t := range tickets {
		for _, pos := range t {
			if !s.valid(pos) {
				continue t
			}
		}
		r = append(r, t)
	}
	return
}

func part1(in input) (rc int) {
	for _, t := range in.nearby {
		for _, field := range t {
			if !in.rules.valid(field) {
				rc += field
			}
		}
	}
	return
}

func part2(in input, prefix string) (rc int) {
	in.nearby = in.rules.filter(in.nearby)

	fields := []string{}
	for k := range in.rules {
		fields = append(fields, k)
	}

	// Construct a map for each position in your ticket to a list of possible fields.
	m := map[int]map[string]struct{}{}
	for i := range in.yours {
		m[i] = map[string]struct{}{}
		for _, k := range fields {
			m[i][k] = struct{}{}
		}
	}
	for _, t := range in.nearby {
		for x, v := range t {
			for k, r := range in.rules {
				if !r.valid(v) {
					delete(m[x], k)
				}
			}
		}
	}
	// Desired result: map of position to field.
	fieldOrder := map[int]string{}

	// Loop: if we have found all fields, exit loop.
	for len(fieldOrder) != len(fields) {
		// Find a position in some ticket that can only be represented by a single field.
		var field string
		for pos, candidates := range m {
			if len(candidates) != 1 {
				continue
			}
			for k := range candidates {
				field = k
				break
			}
			fieldOrder[pos] = field
			break
		}
		// Remove that field from every ticket's possible field list.
		for _, candidates := range m {
			delete(candidates, field)
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
