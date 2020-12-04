package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/dds/aoc2020/lib/inputs"
)

func parse(input string) [][]string {
	lines := strings.Split(input, "\n\n")
	r := make([][]string, 0)
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		r = append(r, fields)
	}

	return r
}

var Input = parse(inputs.Day4())

func Test(t *testing.T) {
	// type test struct {
	// 	input  int
	// 	expect int
	// }

	// tests := []test{
	// 	test{
	// 		// ...
	// 	},
	// }

	// for i, test := range tests {
	// 	t.Run(fmt.Sprint(i), func(t *testing.T) {
	// 		require.Equal(t, test.expect, test.input)
	// 	})
	// }
}

func main() {
	fmt.Println(part1(Input))
	fmt.Println(part2(Input))
}

func part1(input [][]string) (rc int) {
loop:
	for _, row := range input {
		fields := map[string]string{}
		required := []string{
			"byr",
			"iyr",
			"eyr",
			"hgt",
			"hcl",
			"ecl",
			"pid",
		}
		for _, field := range row {
			s := strings.Split(field, ":")
			fields[s[0]] = s[1]
		}
		for _, k := range required {
			if fields[k] == "" {
				continue loop
			}
		}
		rc++
	}
	return
}

func part2(input [][]string) (rc int) {
	type rule struct {
		min, max int
		re       *regexp.Regexp
	}
loop:
	for _, row := range input {
		fields := map[string]string{}
		required := map[string]rule{
			"byr": rule{
				min: 1920, max: 2002,
				re: regexp.MustCompile(`^(\d{4})$`)},
			"iyr": rule{
				min: 2010, max: 2020,
				re: regexp.MustCompile(`^(\d{4})$`)},
			"eyr": rule{
				min: 2020, max: 2030,
				re: regexp.MustCompile(`^(\d{4})$`)},
			"hgt": rule{
				min: 150, max: 193,
				re: regexp.MustCompile(`^(\d+)(cm|in)$`)},
			"hcl": rule{
				re: regexp.MustCompile(`^#[0-9a-f]{6}$`)},
			"ecl": rule{re: regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)},
			"pid": rule{re: regexp.MustCompile(`^(\d{9})$`)},
			"cid": rule{re: regexp.MustCompile(`.*`)},
		}
		for _, field := range row {
			s := strings.Split(field, ":")
			rule := required[s[0]]
			if rule.re == nil {
				continue loop
			}
			if !rule.re.MatchString(s[1]) {
				continue loop
			}
			switch s[0] {
			case "byr":
				fallthrough
			case "iyr":
				fallthrough
			case "eyr":
				u, err := strconv.Atoi(rule.re.FindStringSubmatch(s[1])[1])
				if err != nil {
					panic(err)
				}
				if u < rule.min || u > rule.max {
					continue loop
				}
			case "hgt":
				u, err := strconv.Atoi(rule.re.FindStringSubmatch(s[1])[1])
				if err != nil {
					panic(err)
				}
				switch rule.re.FindStringSubmatch(s[1])[2] {
				case "in":
					if u < 59 || u > 76 {
						continue loop
					}
				default:
					if u < rule.min || u > rule.max {
						continue loop
					}
				}
			}
			fields[s[0]] = s[1]
		}
		delete(fields, "cid")
		fields["cid"] = "ignored"
		for k := range required {
			if fields[k] == "" {
				continue loop
			}
		}
		rc++
	}
	return
}
