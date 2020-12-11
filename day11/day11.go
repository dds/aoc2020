package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/dds/aoc2020/lib"
	"github.com/dds/aoc2020/lib/inputs"
)

var Input = lib.ParseInput(inputs.Day11(), func(s string) []string { return strings.Split(s, "") })

type grid struct {
	w, h int
	g    map[image.Point]string
}

func (g grid) occupiedAdj(p image.Point) (rc int) {
	for _, q := range lib.Neighbors8(p) {
		if g.g[q] != "#" {
			continue
		}
		rc++
	}
	return
}

func (g grid) String() (r string) {
	for y := 0; y <= g.h; y++ {
		for x := 0; x <= g.w; x++ {
			r += g.g[image.Point{x, y}]
		}
		r += "\n"
	}
	r += "\n"
	return
}

func (g *grid) round() (rc int) {
	gp := map[image.Point]string{}
	for k, v := range g.g {
		if v == "L" && g.occupiedAdj(k) == 0 {
			gp[k] = "#"
			rc++
		} else if v == "#" && g.occupiedAdj(k) > 3 {
			gp[k] = "L"
			rc++
		} else {
			gp[k] = g.g[k]
		}
	}
	g.g = gp
	return
}

func (g grid) occupied() (r int) {
	for _, v := range g.g {
		if v == "#" {
			r++
		}
	}
	return
}

func part1(input [][]string) (rc int) {
	g := grid{g: map[image.Point]string{}}
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			g.w = lib.Max(g.w, x)
			g.h = lib.Max(g.h, y)
			g.g[image.Point{x, y}] = input[y][x]
		}
	}
	for k := g.round(); k != 0; k = g.round() {

	}
	return g.occupied()
}

var dirs []image.Point = append(lib.Directions, lib.Diagnals...)

func (g grid) occupiedCanSee(p image.Point) (rc int) {
	// For each direction, extend to we hit something or are outside the map.
loop:
	for _, d := range dirs {
		i := 1
		for {
			q := p.Add(d.Mul(i))
			if g.g[q] == "L" {
				continue loop
			}
			if g.g[q] == "#" {
				rc++
				continue loop
			}
			if q.X > g.w || q.Y > g.h || q.X < 0 || q.Y < 0 {
				continue loop
			}
			i++
		}
	}
	return
}

func (g *grid) round2() (rc int) {
	gp := map[image.Point]string{}
	for k, v := range g.g {
		if v == "L" && g.occupiedCanSee(k) == 0 {
			gp[k] = "#"
			rc++
		} else if v == "#" && g.occupiedCanSee(k) > 4 {
			gp[k] = "L"
			rc++
		} else {
			gp[k] = g.g[k]
		}
	}
	g.g = gp
	return
}

func part2(input [][]string) (rc int) {
	g := grid{g: map[image.Point]string{}}
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			g.w = lib.Max(g.w, x)
			g.h = lib.Max(g.h, y)
			g.g[image.Point{x, y}] = input[y][x]
		}
	}
	for k := g.round2(); k != 0; k = g.round2() {
	}
	return g.occupied()
}

func main() {
	fmt.Println(part1(Input))
	fmt.Println(part2(Input))
}
