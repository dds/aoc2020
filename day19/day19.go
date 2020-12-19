package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dds/aoc2020/lib/inputs"
)

var Input = inputs.Day19()

func part1(in string) (rc int) {
	p := parse(in)
	fmt.Println(p.rules)
	return
}

func parse(in string) (r input) {
	parts := strings.Split(in, "\n\n")
	rules, _ := parts[0], parts[1]
	r.rules = parseRules(rules)
	// r.msgs = parseMsgs(messages)
	return
}

type input struct {
	rules rules
	msgs  []msg
}

var ruleRE = regexp.MustCompile(`^(\d+): (.*)$`)

func parseRules(in string) (r rules) {
	lines := strings.Split(in, "\n")
	r = make(rules, len(lines))
	for _, l := range lines {
		matches := ruleRE.FindStringSubmatch(l)
		if len(matches) == 0 {
			continue
		}
		n, _ := strconv.Atoi(matches[1])
		r[n] = rule(matches[2])
	}
	return
}

type rules []rule

func (s rules) String() (r string) {
	r += "["
	c := []string{}
	for _, l := range s {
		c = append(c, fmt.Sprintf("%q", l))
	}
	r += strings.Join(c, ", ") + "]"
	return
}

type rule string
type msg struct {
}

func main() {
	fmt.Println(part1(Input))
}
