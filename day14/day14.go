package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dds/aoc2020/lib/inputs"
)

var memre = regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)

var Input = inputs.Day14()

func parse(s string) (r []*prog) {
	lines := strings.Split(s, "\n")
	maskprefix := "mask = "
	i := 0
maskline:
	for i < len(lines) {
		if !strings.HasPrefix(lines[i], maskprefix) {
			i++
			continue maskline
		}
		p := &prog{mask: lines[i][len(maskprefix):]}
		r = append(r, p)
		i++
		for {
			matches := memre.FindStringSubmatch(lines[i])
			if len(matches) == 0 {
				continue maskline
			}
			i++
			addr, _ := strconv.Atoi(matches[1])
			val, _ := strconv.Atoi(matches[2])
			p.mem = append(p.mem, mem{addr, val})
		}
	}
	return
}

type prog struct {
	mask string
	mem  []mem
}

func (p prog) and(i int) (r int) {
	r = i
	for i, c := range p.mask {
		pos := len(p.mask) - 1 - i
		if c == 'X' {
			continue
		}
		if c == '0' {
			r &^= 1 << pos
		} else {
			r |= 1 << pos
		}
	}
	return
}

type mem struct {
	addr, val int
}

func part1(input []*prog) (rc int) {
	m := map[int]int{}
	for _, i := range input {
		for _, mem := range i.mem {
			m[mem.addr] = i.and(mem.val)
		}
	}
	for _, v := range m {
		rc += v
	}
	return
}

func (p prog) decode(i int) (r []int) {
	z := i
	floats := []int{}
	for i, c := range p.mask {
		pos := len(p.mask) - 1 - i
		switch c {
		case 'X':
			floats = append(floats, pos)
		case '1':
			z |= (1 << pos)
		}
	}
	var expfloats func(int, []int) []int
	expfloats = func(z int, f []int) (r []int) {
		if len(f) == 0 {
			return []int{z}
		}
		if len(f) == 1 {
			i := []int{
				z &^ (1 << f[0]),
				z | (1 << f[0]),
			}
			return i
		}
		for _, x := range expfloats(z, f[1:]) {
			i := []int{
				x &^ (1 << f[0]),
				x | (1 << f[0]),
			}
			r = append(r, i...)
		}
		return
	}
	return expfloats(z, floats)
}

func part2(input []*prog) (rc int) {
	m := map[int]int{}
	for _, i := range input {
		for _, mem := range i.mem {
			for _, addr := range i.decode(mem.addr) {
				m[addr] = mem.val
			}
		}
	}
	for _, v := range m {
		rc += v
	}
	return
}

func main() {
	fmt.Println(part1(parse(Input)))
	fmt.Println(part2(parse(Input)))
}
