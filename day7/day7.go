package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dds/aoc2020/lib"
	"github.com/dds/aoc2020/lib/inputs"
)

func parse(s string) []string {
	parts := lib.TrimSpace(strings.Split(s, "bags contain"))
	if len(parts) < 2 {
		return []string{}
	}
	bags := strings.Split(parts[1], ", ")
	r := []string{parts[0]}
	// Drop the last word from each group.
	for _, b := range bags {
		t := strings.Fields(b)
		r = append(r, strings.Join(t[:len(t)-1], " "))
	}
	return r
}

var Input = lib.ParseInput(inputs.Day7(), parse)

func part1(input [][]string) (rc int) {
	m := map[string][]string{}
	for _, row := range input {
		for _, s := range row[1:] {
			bag := s[2:]
			m[bag] = append(m[bag], row[0])
		}
	}
	bagQueue := []string{"shiny gold"}
	seenBags := map[string]int{}
	for len(bagQueue) > 0 {
		bag := bagQueue[0]
		bagQueue = bagQueue[1:]
		for _, u := range m[bag] {
			if seenBags[u] == 1 {
				continue
			}
			seenBags[u] = 1
			bagQueue = append(bagQueue, u)
			rc++
		}
	}
	return
}

func part2(input [][]string) (rc int) {
	type bagCount struct {
		count int
		bag   string
	}
	m := map[string][]bagCount{}
	for _, row := range input {
		for _, s := range row[1:] {
			count, err := strconv.Atoi(s[:1])
			if err != nil {
				count = 0
			}
			bag := s[2:]
			m[row[0]] = append(m[row[0]], bagCount{count: count, bag: bag})
		}
	}
	var count func(map[string][]bagCount, string) int
	count = func(m map[string][]bagCount, bag string) (sum int) {
		for _, u := range m[bag] {
			sum += u.count * (1 + count(m, u.bag))
		}
		return sum
	}
	return count(m, "shiny gold")
}

func main() {
	fmt.Println(part1(Input))
	fmt.Println(part2(Input))
}
