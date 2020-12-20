package main

import (
	"fmt"
	"image"
	"regexp"
	"strconv"

	"github.com/dds/aoc2020/lib"
	"github.com/dds/aoc2020/lib/inputs"
)

func parse(s string) []string {
	// R180
	re := regexp.MustCompile(`(\w)(\d+)`)
	matches := re.FindStringSubmatch(s)
	if len(matches) == 0 {
		return matches
	}
	return []string{matches[1], matches[2]}
}

var Input = lib.ParseInput(inputs.Day12(), parse)

// directions sorted E S W N
var dirs = []image.Point{
	image.Point{1, 0},
	image.Point{0, -1},
	image.Point{-1, 0},
	image.Point{0, 1},
}

func part1(input [][]string) (rc int) {
	p := image.Point{}
	dir := 0
	for _, row := range input {
		d, ns := row[0], row[1]
		n, _ := strconv.Atoi(ns)
		switch d {
		case "N":
			p = p.Add(lib.Directions[0].Mul(n))
		case "S":
			p = p.Add(lib.Directions[1].Mul(n))
		case "W":
			p = p.Add(lib.Directions[2].Mul(n))
		case "E":
			p = p.Add(lib.Directions[3].Mul(n))
		case "L":
			dir -= n / 90
			dir %= 4
			if dir < 0 {
				dir += 4
			}
		case "R":
			dir += n / 90
			dir %= 4
		case "F":
			p = p.Add(dirs[dir].Mul(n))
		}
	}
	return int(lib.Taxi(p, image.Point{}))
}

// 1 for 90, 2, for 180, 3 for 270 CW
// -1 .. -2 for CCW
func rotate(p image.Point, dir int) image.Point {
	dir %= 4
	switch dir {
	case -1:
		dir = 3
	case -2:
		dir = 2
	case -3:
		dir = 1
	}
	switch dir {
	case 1:
		return image.Point{p.Y, -p.X}
	case 2:
		return image.Point{-p.X, -p.Y}
	case 3:
		return image.Point{-p.Y, p.X}
	default:
		return p
	}
}

func part2(input [][]string) (rc int) {
	p := image.Point{}
	wpt := image.Point{10, 1}
	for _, row := range input {
		d, ns := row[0], row[1]
		n, _ := strconv.Atoi(ns)
		dir := 0
		switch d {
		case "N":
			wpt = wpt.Add(lib.Directions[0].Mul(n))
		case "S":
			wpt = wpt.Add(lib.Directions[1].Mul(n))
		case "W":
			wpt = wpt.Add(lib.Directions[2].Mul(n))
		case "E":
			wpt = wpt.Add(lib.Directions[3].Mul(n))
		case "L":
			dir = -n / 90
			wpt = rotate(wpt, dir)
		case "R":
			dir = n / 90
			wpt = rotate(wpt, dir)
		case "F":
			p = p.Add(wpt.Mul(n))
		}
	}
	return int(lib.Taxi(p, image.Point{}))
}

func main() {
	fmt.Println(part1(Input))
	fmt.Println(part2(Input))
}
