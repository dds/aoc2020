package main

import (
	"fmt"
	"strings"

	"github.com/dds/aoc2020/lib"
	"github.com/dds/aoc2020/lib/inputs"
)

type point struct {
	X, Y, Z, W int
}

func (p point) add(q point) point {
	return point{p.X + q.X, p.Y + q.Y, p.Z + q.Z, p.W + q.W}
}

func pt(x, y, z, w int) point {
	return point{x, y, z, w}
}

func (p point) neighbors() (r []point) {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				for w := -1; w <= 1; w++ {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						continue
					}
					r = append(r, p.add(pt(x, y, z, w)))
				}
			}
		}
	}
	return
}

type board map[point]rune

type dims struct {
	minX, minY, minZ, maxX, maxY, maxZ, minW, maxW int
}

func (b board) dims() (r dims) {
	// Dims are the min and max X, Y, and Z coord in each dimension.
	m := map[point]bool{}
	for p := range b {
		for _, q := range p.neighbors() {
			if m[q] {
				continue
			}
			r.minX = lib.Min(r.minX, q.X)
			r.minY = lib.Min(r.minY, q.Y)
			r.minZ = lib.Min(r.minZ, q.Z)
			r.minW = lib.Min(r.minW, q.W)
			r.maxX = lib.Max(r.maxX, q.X)
			r.maxY = lib.Max(r.maxY, q.Y)
			r.maxZ = lib.Max(r.maxZ, q.Z)
			r.maxW = lib.Max(r.maxW, q.W)
			m[q] = true
		}
	}
	return r
}

func parse(s string) (r board) {
	r = board{}
	for y, line := range strings.Split(s, "\n") {
		for x, c := range line {
			p := pt(x, y, 0, 0)
			if c == '#' {
				r[p] = c
			}
		}
	}
	return
}

func (b *board) cycle() {
	old := *b
	new := board{}
	dims := old.dims()
	for w := dims.minW; w <= dims.maxW; w++ {
		for z := dims.minZ; z <= dims.maxZ; z++ {
			for y := dims.minY; y <= dims.maxY; y++ {
				for x := dims.minX; x <= dims.maxX; x++ {
					p := pt(x, y, z, w)
					active := 0
					for _, q := range p.neighbors() {
						if old[q] == '#' {
							active++
						}
					}
					if old[p] == '#' && (active == 2 || active == 3) {
						new[p] = '#'
					} else if old[p] != '#' && (active == 3) {
						new[p] = '#'
					}
				}
			}
		}
	}
	*b = new
}

func part2(input board) (rc int) {
	for i := 0; i < 6; i++ {
		input.cycle()
	}
	return len(input)
}

func main() {
	fmt.Println(part2(parse(inputs.Day17())))
}
