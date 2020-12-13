package main

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/dds/aoc2020/lib/inputs"
)

func part1(input string) (rc int) {
	s := parse(input)
	mins := make([]int, len(s.buses()))
	for idx, n := range s.busIDs() {
		for mins[idx] < s.id {
			mins[idx] += n
		}
	}
	min, minIdx := 1<<31, 0
	for i, n := range mins {
		if n < min {
			min = n
			minIdx = i
		}
	}
	return s.busIDs()[minIdx] * (min - s.id)
}

func part2(input string) (rc uint64) {
	s := parse(input)
	buses := s.buses()
	i := 0
	eqs := make([]string, len(buses))
	maxBus, maxBusIdx := 0, 0
	for j := 0; j < len(buses); j++ {
		if buses[j].id > maxBus {
			maxBus = buses[j].id
			maxBusIdx = buses[j].pos
		}
		eqs[j] = fmt.Sprintf("(t+%d)%%%d=0", buses[j].pos, buses[j].id)
	}

	fmt.Println("Visit www.wolframalpha.com/input?i=" + url.QueryEscape(strings.Join(eqs, ", ")))
	fmt.Println("max bus: ", maxBus, maxBusIdx)
loop:
	for {
		i++
		p := uint64(i)*uint64(maxBus) - uint64(maxBusIdx)
		for _, bus := range buses {
			t := uint64(p) + uint64(bus.pos)
			if t%uint64(bus.id) != 0 {
				continue loop
			}
		}
		return p
	}
}

type bus struct {
	id  int
	pos int
}

func (s sched) buses() (r []bus) {
	for i, id := range s.ids {
		n, err := strconv.Atoi(id)
		if err != nil {
			continue
		}
		r = append(r, bus{n, i})
	}
	return
}
func parse(i string) (r sched) {
	parts := strings.Split(i, "\n")
	n, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	r.id = n
	r.ids = strings.Split(parts[1], ",")
	return
}

type sched struct {
	id  int
	ids []string
}

func (s sched) busIDs() (r []int) {
	for _, i := range s.ids {
		n, err := strconv.Atoi(i)
		if err != nil {
			continue
		}
		r = append(r, n)
	}
	return
}

func main() {
	fmt.Println(part1(inputs.Day13()))
	fmt.Println(part2(inputs.Day13()))
}
