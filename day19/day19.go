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
	re := regexp.MustCompile(`(?m)^` + p.rules.regex("0") + `$`)
	return len(re.FindAllString(p.msgs, -1))
}

func part2(in string) (rc int) {
	p := parse(in)
	p.rules["8"] = `"` + p.rules.regex("42") + `+"`
	p.rules["11"] = ""
	for i := 1; i < 11; i++ {
		p.rules["11"] += fmt.Sprintf("|%s{%d}%s{%d}", p.rules.regex("42"), i, p.rules.regex("31"), i)
	}
	p.rules["11"] = `"(?:` + p.rules["11"][1:] + `)"`
	fmt.Println(p.rules)
	re := regexp.MustCompile(`(?m)^` + p.rules.regex("0") + `$`)
	return len(re.FindAllString(p.msgs, -1))
}

func (q rules) regex(rule string) string {
	if q[rule][0] == '"' {
		return q[rule][1 : len(q[rule])-1]
	}
	re := ""
	for _, s := range strings.Split(q[rule], " | ") {
		re += "|"
		for _, s := range strings.Fields(s) {
			re += q.regex(s)
		}
	}
	return `(?:` + re[1:] + `)`
}

func parse(in string) (r input) {
	parts := strings.Split(in, "\n\n")
	r.rules = parseRules(parts[0])
	r.msgs = parts[1]
	return
}

type input struct {
	rules rules
	msgs  string
}

var ruleRE = regexp.MustCompile(`^(\d+): (.*)$`)

func parseRules(in string) (r rules) {
	lines := strings.Split(in, "\n")
	r = make(rules)
	for _, l := range lines {
		matches := ruleRE.FindStringSubmatch(l)
		if len(matches) == 0 {
			continue
		}
		r[matches[1]] = matches[2]
	}
	return
}

type rules map[string]string

func (s rules) String() (r string) {
	r += "{\n"
	c := make([]string, len(s))
	for i, l := range s {
		n, _ := strconv.Atoi(i)
		c[n] = fmt.Sprintf("%v: %q", i, l)
	}
	r += strings.Join(c, "\n") + "\n}"
	return
}

func main() {
	fmt.Println(part1(Input))
	fmt.Println(part2(Input))
}
